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

func (i *fileData) serialize() []byte {
	data := make([]byte, 2+len(i.key)+len(i.mimetype))
	pos := 0
	pos = writeString(data, pos, i.key)
	pos = writeString(data, pos, i.mimetype)
	return data
}

func (s *LocalStore) deserialize(key string) *fileData {
	pos := s.findKey(key)
	if pos < 0 {
		return nil
	}
	file, err := os.Open(s.itemFile)
	if err != nil {
		return nil
	}
	defer file.Close()

	buff := make([]byte, 256)
	file.ReadAt(buff, pos)
	keyLength := int(buff[0])
	valueOffset := 1 + keyLength
	valueLength := int64(buff[valueOffset])
	foundKey := string(buff[1 : keyLength+1])

	foundMimetype := string(buff[valueOffset + 1 : int64(valueOffset + 1) + valueLength])

	return &fileData{foundKey, foundMimetype}
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
	pos := s.findKey(key)
	return pos >= 0
}

func (s *LocalStore) mimetype(key string) string {
	fileData := s.deserialize(key)
	if fileData == nil {
		return ""
	}
	return fileData.mimetype
}

func (s *LocalStore) findKey(key string) int64 {
	file, err := os.Open(s.itemFile)
	if err != nil {
		return -1
	}
	defer file.Close()

	offset := int64(0)
	for i := 0; i < 10000; i++ {
		buff := make([]byte, 256)
		file.ReadAt(buff, offset)
		keyLength := int64(buff[0])
		if keyLength == 0 {
			return -1
		}
		foundKey := string(buff[1 : keyLength+1])
		if foundKey == key {
			return offset
		}
		valueLength := int64(buff[1+keyLength])
		offset += keyLength + valueLength + 2
	}
	panic("took over 10000 iterations")
}