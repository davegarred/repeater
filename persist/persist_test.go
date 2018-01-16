package persist

import (
	"testing"
)

const A_KEY = "a_key"
const A_VALUE = "a test string to store"
const A_VALUE_2 = "a second test string to store"

func TestInterface(t *testing.T) {
	var _ Store = NewStore()
}

func TestMemStore(t *testing.T) {
	s := NewStore()
	s.Store(A_KEY, A_VALUE)
	result, err := s.Retrieve(A_KEY)
	if result != A_VALUE || err != nil {
		t.Errorf("incorrect value found")
	}
}

func TestConflict(t *testing.T) {
	s := NewStore()
	if err := s.Store(A_KEY, A_VALUE); err != nil {
		t.Errorf("error saving value")
	}
	if err := s.Store(A_KEY, A_VALUE_2); err == nil {
		t.Errorf("identical key saved with no error, error expected")
	}
}
