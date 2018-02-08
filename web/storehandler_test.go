package web

import (
	"testing"

	"github.com/davegarred/repeater/persist"
)

func TestGetStoreHandler(t *testing.T) {
	w := responseWriter()
	r := request("/store?" + aTestQuery)
	store := persist.NewMemStore()
	getStoreHandler(w, r, store)

	uuidLength := 36
	assertEquals(t, uuidLength, len(w.writtenOut))
	assertEquals(t, uuidLength, len(w.header["X-Document-Id"][0]))
	storedVal, _ := store.Retrieve(w.writtenOut)
	assertEquals(t, aTestValue, storedVal.Object)
}

func TestGetStoreHandler_userDefinedName(t *testing.T) {
	w := responseWriter()
	r := request("/store/someName?" + aTestQuery)
	store := persist.NewMemStore()
	getStoreHandler(w, r, store)

	someName := "someName"
	assertEquals(t, someName, w.writtenOut)
	assertEquals(t, someName, w.header["X-Document-Id"][0])
	storedVal, _ := store.Retrieve(w.writtenOut)
	assertEquals(t, aTestValue, storedVal.Object)
}

func TestGetStoreHandler_nameUsedTwice(t *testing.T) {
	w := responseWriter()
	r := request("/store/someName?" + aTestQuery)
	store := persist.NewMemStore()
	getStoreHandler(w, r, store)

	w = responseWriter()
	getStoreHandler(w, r, store)

	assertEquals(t, "Document already exists with this key", w.writtenOut)
	assertEquals(t, 400, w.headerInt)
}

func TestPostStoreHandler_noContentType(t *testing.T) {
	w := responseWriter()
	r := requestWithBody("/store", "somedata")
	store := persist.NewMemStore()
	postStoreHandler(w, r, store)

	uuidLength := 36
	assertEquals(t, uuidLength, len(w.writtenOut))
	assertEquals(t, uuidLength, len(w.header["X-Document-Id"][0]))
	storedVal, _ := store.Retrieve(w.writtenOut)
	assertEquals(t, "somedata", storedVal.Object)
	assertEquals(t, "application/octet-stream", storedVal.Mimetype)
}

func TestPostStoreHandler(t *testing.T) {
	w := responseWriter()
	r := requestWithContentTypeAndBody("/store", "image/something", "somedata")
	store := persist.NewMemStore()
	postStoreHandler(w, r, store)

	uuidLength := 36
	assertEquals(t, uuidLength, len(w.writtenOut))
	assertEquals(t, uuidLength, len(w.header["X-Document-Id"][0]))
	storedVal, _ := store.Retrieve(w.writtenOut)
	assertEquals(t, "somedata", storedVal.Object)
	assertEquals(t, "image/something", storedVal.Mimetype)
}
