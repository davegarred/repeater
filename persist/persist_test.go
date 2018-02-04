package persist

import (
	"os"
	"testing"
)

const (
	aTestKey    = "a_key"
	aTestKey2   = "another_key"
	aTestValue  = "a test string to store"
	aTestValue2 = "a second test string to store"
	applicationJson = "application/json"
)

type storer interface {
	Store(string, string, string) error
	Retrieve(string) (string, error)
	Delete(string) error
}

var test_directory string

func init() {
	test_directory = os.Getenv("HOME") + "/.test"
}

func TestInterface(t *testing.T) {
	var _ storer = NewMemStore()
	var _ storer = NewLocalStore(test_directory)
}

func TestMemStore(t *testing.T) {
	s := NewMemStore()
	storeAndRetrieve(t, s)
	storeAndRetrieveAgain(t, s)
}

func TestLocalStore(t *testing.T) {
	s := NewLocalStore(test_directory)
	s.deleteAll()
	storeAndRetrieve(t, s)
	storeAndRetrieveAgain(t, s)
}

func TestMemStoreConflict(t *testing.T) {
	s := NewMemStore()
	storeConflict(t, s)
}

func TestLocalStoreConflict(t *testing.T) {
	s := NewLocalStore(test_directory)
	s.deleteAll()
	storeConflict(t, s)
}

func storeAndRetrieve(t *testing.T, s storer) {
	if e := s.Store(applicationJson, aTestKey, aTestValue); e != nil {
		t.Errorf("error storing value: %v", e)
	}
	result, err := s.Retrieve(aTestKey)
	if result != aTestValue || err != nil {
		t.Errorf("incorrect value found")
	}
}

func storeAndRetrieveAgain(t *testing.T, s storer) {
	if e := s.Store(applicationJson, aTestKey2, aTestValue2); e != nil {
		t.Errorf("error storing value: %v", e)
	}
	result, err := s.Retrieve(aTestKey2)
	if result != aTestValue2 || err != nil {
		t.Errorf("incorrect value found")
	}
}

func storeConflict(t *testing.T, s storer) {
	if err := s.Store(applicationJson, aTestKey, aTestValue); err != nil {
		t.Errorf("error saving value")
	}
	if err := s.Store(applicationJson, aTestKey, aTestValue2); err == nil {
		t.Errorf("identical key saved with no error, error expected")
	}
}
