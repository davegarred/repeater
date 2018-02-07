package web

import (
	"testing"

	"github.com/davegarred/repeater/persist"
)


func TestRetrieveHandler(t *testing.T) {
	w := responseWriter()
	r := request("/retrieve/" + aTestKey)
	store := persist.NewMemStore()
	store.Store("application/json", aTestKey, aTestValue)

	retrieveHandler(w, r, store)
	assertEquals(t, aTestValue, w.writtenOut)
	assertEquals(t, "application/json", w.header["Content-Type"][0])
}

func TestRetrieveHandler_incorrectSignature(t *testing.T) {
	w := responseWriter()
	r := request("/retrieve")
	store := persist.NewMemStore()
	store.Store("application/json", aTestKey, aTestValue)

	retrieveHandler(w, r, store)
	assertEquals(t, "", w.writtenOut)
	assertEquals(t, 404, w.headerInt)
}

func TestRetrieveHandler_notFound(t *testing.T) {
	w := responseWriter()
	r := request("/retrieve/not_found")
	store := persist.NewMemStore()

	retrieveHandler(w, r, store)

	expected := "404 page not found\n"
	assertEquals(t, expected, w.writtenOut)
	assertEquals(t, 404, w.headerInt)
}

