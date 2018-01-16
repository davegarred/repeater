package main

import (
	"testing"

	"github.com/google/uuid"
)

func TestInterface(t *testing.T) {
	var f KeyFactory = NewKeyFactory()
	key := f.Get()
	if _, err := uuid.Parse(key); err != nil {
		t.Error("invalid UUID found")
	}
}
