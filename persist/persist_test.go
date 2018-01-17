package persist

import (
	"testing"
)

const A_KEY = "a_key"
const A_KEY_2 = "another_key"
const A_VALUE = "a test string to store"
const A_VALUE_2 = "a second test string to store"

func TestInterface(t *testing.T) {
	var _ Store = NewMemStore()
}

func TestMemStore(t *testing.T) {
	s := NewMemStore()
	s.Store(A_KEY, A_VALUE)
	result, err := s.Retrieve(A_KEY)
	if result != A_VALUE || err != nil {
		t.Errorf("incorrect value found")
	}
}

func TestMemStoreConflict(t *testing.T) {
	s := NewMemStore()
	if err := s.Store(A_KEY, A_VALUE); err != nil {
		t.Errorf("error saving value")
	}
	if err := s.Store(A_KEY, A_VALUE_2); err == nil {
		t.Errorf("identical key saved with no error, error expected")
	}
}

func TestLocalStore(t *testing.T) {
	s := NewLocalStore("/home/ubuntu/.test")
	s.deleteAll()
	if e := s.Store(A_KEY, A_VALUE); e != nil {
		t.Errorf("error storing value: %v", e)
	}
	result, err := s.Retrieve(A_KEY)
	if result != A_VALUE || err != nil {
		t.Errorf("incorrect value found")
	}
}
