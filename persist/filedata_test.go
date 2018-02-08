package persist

import (
	"testing"
)

const genericMimetype = "audio/slayer"

func TestRegister(t *testing.T) {
	s := NewLocalStore(testDirectory)
	s.deleteAll()
	s.registerNewItem(aTestKey, genericMimetype)
	s.registerNewItem(aTestKey2, applicationJSON)
	if !s.exists(aTestKey) {
		t.Error()
	}
	if !s.exists(aTestKey2) {
		t.Error()
	}
	if s.mimetype(aTestKey) != genericMimetype {
		t.Error()
	}
	if s.mimetype(aTestKey2) != applicationJSON {
		t.Error()
	}
}