package persist

import (
	"fmt"
	"testing"
)

const A_VALUE = "a test string to store"

func TestStore(t *testing.T) {
	s := NewStore()
	key := s.Store(A_VALUE)
	result := s.Retrieve(key)
	if result != A_VALUE {
		t.Errorf("incorrect value found")
	}
	fmt.Printf("%v\n", key)
}
