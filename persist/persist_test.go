package persist

import (
	"os"
	"testing"
)

const (
	aTestKey        = "a_key"
	aTestKey2       = "another_key"
	aTestValue      = "a test string to store"
	aTestValue2     = "a second test string to store"
	applicationJSON = "application/json"
)

type storer interface {
	Store(string, string, string) error
	Retrieve(string) (string, error)
	Delete(string) error
}

var testDirectory string

func init() {
	testDirectory = os.Getenv("HOME") + "/.test"
}

func TestInterface(t *testing.T) {
	var _ storer = NewMemStore()
	var _ storer = NewLocalStore(testDirectory)
}

func TestMemStore(t *testing.T) {
	s := NewMemStore()
	storeAndRetrieve(t, s)
	storeAndRetrieveAgain(t, s)
}

func TestLocalStore(t *testing.T) {
	s := NewLocalStore(testDirectory)
	s.deleteAll()
	storeAndRetrieve(t, s)
	storeAndRetrieveAgain(t, s)
}

func TestMemStoreConflict(t *testing.T) {
	s := NewMemStore()
	storeConflict(t, s)
}

func TestLocalStoreConflict(t *testing.T) {
	s := NewLocalStore(testDirectory)
	s.deleteAll()
	storeConflict(t, s)
}

func storeAndRetrieve(t *testing.T, s storer) {
	if e := s.Store(applicationJSON, aTestKey, aTestValue); e != nil {
		t.Errorf("error storing value: %v", e)
	}
	result, err := s.Retrieve(aTestKey)
	if result != aTestValue || err != nil {
		t.Errorf("incorrect value found")
	}
}

func storeAndRetrieveAgain(t *testing.T, s storer) {
	if e := s.Store(applicationJSON, aTestKey2, aTestValue2); e != nil {
		t.Errorf("error storing value: %v", e)
	}
	result, err := s.Retrieve(aTestKey2)
	if result != aTestValue2 || err != nil {
		t.Errorf("incorrect value found")
	}
}

func storeConflict(t *testing.T, s storer) {
	if err := s.Store(applicationJSON, aTestKey, aTestValue); err != nil {
		t.Errorf("error saving value")
	}
	if err := s.Store(applicationJSON, aTestKey, aTestValue2); err == nil {
		t.Errorf("identical key saved with no error, error expected")
	}
}
