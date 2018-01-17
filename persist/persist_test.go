package persist

import (
	"testing"
)

const (
	TEST_FILE_HOME = "/home/ubuntu/.test"
	A_KEY          = "a_key"
	A_KEY_2        = "another_key"
	A_VALUE        = "a test string to store"
	A_VALUE_2      = "a second test string to store"
)

func TestInterface(t *testing.T) {
	var _ Store = NewMemStore()
	var _ Store = NewLocalStore(TEST_FILE_HOME)
}

func TestMemStore(t *testing.T) {
	s := NewMemStore()
	storeAndRetrieve(t, s)
}

func TestLocalStore(t *testing.T) {
	s := NewLocalStore("/home/ubuntu/.test")
	s.deleteAll()
	storeAndRetrieve(t, s)
}

func TestMemStoreConflict(t *testing.T) {
	s := NewMemStore()
	storeConflict(t, s)
}

func TestLocalStoreConflict(t *testing.T) {
	s := NewMemStore()
	storeConflict(t, s)
}

func storeAndRetrieve(t *testing.T, s Store) {
	if e := s.Store(A_KEY, A_VALUE); e != nil {
		t.Errorf("error storing value: %v", e)
	}
	result, err := s.Retrieve(A_KEY)
	if result != A_VALUE || err != nil {
		t.Errorf("incorrect value found")
	}
}

func storeConflict(t *testing.T, s Store) {
	if err := s.Store(A_KEY, A_VALUE); err != nil {
		t.Errorf("error saving value")
	}
	if err := s.Store(A_KEY, A_VALUE_2); err == nil {
		t.Errorf("identical key saved with no error, error expected")
	}
}
