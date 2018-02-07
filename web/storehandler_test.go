package web

import (
	"testing"

	"github.com/davegarred/repeater/persist"
)

func TestStoreHandler(t *testing.T) {
	w := responseWriter()
	r := request("/store?" + aTestQuery)
	store := persist.NewMemStore()
	storeHandler(w, r, store)

	uuidLength := 36
	assertEquals(t, uuidLength, len(w.writtenOut))
	assertEquals(t, uuidLength, len(w.header["X-Document-Id"][0]))
	storedVal, _ := store.Retrieve(w.writtenOut)
	assertEquals(t, aTestValue, storedVal.Object)
}

func TestStoreHandler_userDefinedName(t *testing.T) {
	w := responseWriter()
	r := request("/store/someName?" + aTestQuery)
	store := persist.NewMemStore()
	storeHandler(w, r, store)

	someName := "someName"
	assertEquals(t, someName, w.writtenOut)
	assertEquals(t, someName, w.header["X-Document-Id"][0])
	storedVal, _ := store.Retrieve(w.writtenOut)
	assertEquals(t, aTestValue, storedVal.Object)
}

func TestStoreHandler_nameUsedTwice(t *testing.T) {
	w := responseWriter()
	r := request("/store/someName?" + aTestQuery)
	store := persist.NewMemStore()
	storeHandler(w, r, store)

	w = responseWriter()
	storeHandler(w, r, store)

	assertEquals(t, "Document already exists with this key", w.writtenOut)
	assertEquals(t, 400, w.headerInt)
}