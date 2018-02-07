package web

import (
	"testing"

	"github.com/davegarred/repeater/persist"
)


func TestDeleteHandler(t *testing.T) {
	w := responseWriter()
	r := request("/delete/" + aTestKey)
	store := persist.NewMemStore()

	deleteHandler(w, r, store)

	assertEquals(t, "", w.writtenOut)
	assertEquals(t, 0, len(w.header))
}

func TestDeleteHandler_incorrectSignature(t *testing.T) {
	w := responseWriter()
	r := request("/delete")
	store := persist.NewMemStore()

	deleteHandler(w, r, store)

	assertEquals(t, "", w.writtenOut)
	assertEquals(t, 0, len(w.header))
}