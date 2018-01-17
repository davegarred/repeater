package persist

import (
	"fmt"
	"io/ioutil"
	"os"
	"sync"

	"github.com/davegarred/repeater/util"
)

type LocalStore struct {
	location string
	mu       sync.Mutex
}

func NewLocalStore(l string) *LocalStore {
	dirList, err := ioutil.ReadDir(l)
	if err != nil {
		panic(err)
	}
	util.Log("directory %v (%T)\n\tlisting:\n", l, dirList)
	for _, fileInfo := range dirList {
		util.Log("- %v %v\n", fileInfo.Mode(), fileInfo.Name())
	}
	return &LocalStore{location: l}
}

func (s *LocalStore) Store(k string, v string) error {
	filename := s.buildFilename(k)
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, e := os.Stat(filename); e == nil {
		util.Log("Attempted store key conflict with key: %v", k)
		return KEY_CONFLICT
	}
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	fmt.Fprint(file, v)
	return nil
}

func (s *LocalStore) Retrieve(key string) (string, error) {
	filename := s.buildFilename(key)
	if _, e := os.Stat(filename); e != nil {
		util.Log("Unable to find object with key: %v", key)
		return "", nil
	}
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	val, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}

	return string(val), nil
}

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
