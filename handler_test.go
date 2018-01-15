package main

import (
	"net/http"
	"net/url"
	"testing"

	"github.com/davegarred/repeater/persist"
)

type ResponseWriterStub struct {
}

func (w *ResponseWriterStub) Header() http.Header {
	return make(map[string][]string)
}
func (w *ResponseWriterStub) Write([]byte) (int, error) {
	return 0, nil
}
func (w *ResponseWriterStub) WriteHeader(int) {
}

func TestStoreHandler(t *testing.T) {
	w := &ResponseWriterStub{}
	//&http.ResponseWriter{}
	r := &http.Request{
		URL: &url.URL{},
	}

	store := persist.NewStore()
	retrieveHandler(w, r, store)
	storeHandler(w, r, store)

}
