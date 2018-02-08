package persist

import (
	"fmt"
	"io/ioutil"
	"os"
	"sync"

	"github.com/davegarred/repeater/log"
)

// LocalStore will use the local filesystem to store objects
type LocalStore struct {
	location string
	itemFile string
	mu       sync.Mutex
}

// NewLocalStore returns a new *LocalStore using the given location
func NewLocalStore(l string) *LocalStore {
	dirList, err := ioutil.ReadDir(l)
	if err != nil {
		panic(err)
	}
	log.Log("directory %v (%T)\n\tlisting:\n", l, dirList)
	for _, fileInfo := range dirList {
		log.Log("- %v %v\n", fileInfo.Mode(), fileInfo.Name())
	}

	storedItemFile := l + "/.db"
	return &LocalStore{location: l, itemFile: storedItemFile}
}

// Store will persist the object value and mimetype indexed by the key
func (s *LocalStore) Store(mimetype string, key string, value string) error {
	filename := s.buildFilename(key)
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.exists(key) {
		return KeyConflict
	}
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	fmt.Fprint(file, value)
	s.registerNewItem(key, mimetype)
	return nil
}

// Retrieve takes a key and returns the stored value or an error
func (s *LocalStore) Retrieve(key string) (*StoredObject, error) {
	filename := s.buildFilename(key)
	if _, e := os.Stat(filename); e != nil {
		log.Log("Unable to find object with key: %v", key)
		return nil, nil
	}
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	val, err := ioutil.ReadAll(file)
	if len(val) == 0 || err != nil {
		return nil, nil
	}
	mimetype := s.mimetype(key)
	storedObject := &StoredObject{mimetype, string(val)}
	return storedObject, nil
}

// Delete removes the key-value pair associated with the given key
func (s *LocalStore) Delete(k string) error {
	filename := s.buildFilename(k)
	os.Remove(filename)
	return nil
}

func (s *LocalStore) buildFilename(k string) string {
	return s.location + "/" + k + ".json"
}

func (s *LocalStore) deleteAll() {
	dirList, err := ioutil.ReadDir(s.location)
	if err != nil {
		panic(err)
	}
	for _, fileInfo := range dirList {
		filename := s.location + "/" + fileInfo.Name()
		e := os.Remove(filename)
		if e != nil {
			fmt.Printf("error removing file %v - %v", filename, e)
		}
	}
}
