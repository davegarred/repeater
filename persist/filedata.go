package persist

import (
	"os"
)

type fileData struct {
	key      string
	mimetype string
}

func newFileData(key string, mimetype string) *fileData {
	return &fileData{key, mimetype}
}
func (i *fileData) serialize() []byte {
	data := make([]byte, 2+len(i.key)+len(i.mimetype))
	pos := 0
	pos = writeString(data, pos, i.key)
	pos = writeString(data, pos, i.mimetype)
	return data
}

func writeString(data []byte, pos int, value string) int {
	l := len(value)
	data[pos] = byte(l)
	newData := []byte(value)
	for i, v := range newData {
		data[pos+1+i] = v
	}
	return pos + 1 + l
}

func (s *LocalStore) exists(key string) bool {
	file, err := os.Open(s.itemFile)
	if err != nil {
		return false
	}
	defer file.Close()
	offset := int64(0)
	for i := 0; i < 10000; i++ {
		buff := make([]byte, 256)
		file.ReadAt(buff, offset)
		keyLength := int(buff[0])
		if keyLength == 0 {
			return false
		}
		foundKey := string(buff[1 : keyLength+1])
		if foundKey == key {
			return true
		}
		valueLength := int64(buff[2+keyLength])
		offset += valueLength + 2
	}
	panic("took over 10000 iterations")
}

func (s *LocalStore) registerNewItem(key string, mimetype string) {
	storedItem := newFileData(key, mimetype)
	data := storedItem.serialize()

	file, err := os.OpenFile(s.itemFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	if _, err := file.Write(data); err != nil {
		panic(err)
	}
}
