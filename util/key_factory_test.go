package util

import (
	"testing"

	"github.com/google/uuid"
)

func TestKeyFactoryInterface(t *testing.T) {
	var f KeyFactory = NewKeyFactory()
	key := f.Get()
	if _, err := uuid.Parse(key); err != nil {
		t.Error("invalid UUID found")
	}
}
