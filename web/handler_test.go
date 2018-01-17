package web

import (
	"net/http"
	"net/url"
	"testing"

	"github.com/davegarred/repeater/persist"
)

const KEY = "a_key"
const VAL = "{\"name\":\"simple json object\"}"
const QUERY = "name=simple%20json%20object"

func TestStoreHandler(t *testing.T) {
	w := responseWriter()
	r := request("/store?" + QUERY)
	store := persist.NewMemStore()
	storeHandler(w, r, store)

	uuidLength := 36
	assertEquals(t, uuidLength, len(w.writtenOut))
	assertEquals(t, uuidLength, len(w.header["X-Document-Id"][0]))
	storedVal, _ := store.Retrieve(w.writtenOut)
	assertEquals(t, VAL, storedVal)
}

func TestStoreHandler_userDefinedName(t *testing.T) {
	w := responseWriter()
	r := request("/store/someName?" + QUERY)
	store := persist.NewMemStore()
	storeHandler(w, r, store)

	someName := "someName"
	assertEquals(t, someName, w.writtenOut)
	assertEquals(t, someName, w.header["X-Document-Id"][0])
	storedVal, _ := store.Retrieve(w.writtenOut)
	assertEquals(t, VAL, storedVal)
}

func TestStoreHandler_nameUsedTwice(t *testing.T) {
	w := responseWriter()
	r := request("/store/someName?" + QUERY)
	store := persist.NewMemStore()
	storeHandler(w, r, store)

	w = responseWriter()
	storeHandler(w, r, store)

	assertEquals(t, "Document already exists with this key", w.writtenOut)
	assertEquals(t, 400, w.headerInt)
}

func TestRetrieveHandler(t *testing.T) {
	w := responseWriter()
	r := request("/retrieve/" + KEY)
	store := persist.NewMemStore()
	store.Store(KEY, VAL)

	retrieveHandler(w, r, store)
	assertEquals(t, VAL, w.writtenOut)
	assertEquals(t, "application/json", w.header["Content-Type"][0])
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

func TestDeleteHandler(t *testing.T) {
	w := responseWriter()
	r := request("/delete/" + KEY)
	store := persist.NewMemStore()

	deleteHandler(w, r, store)

	assertEquals(t, "", w.writtenOut)
	assertEquals(t, 0, len(w.header))
}

func assertEquals(t *testing.T, expected interface{}, actual interface{}) {
	if actual != expected {
		t.Errorf("Values not equal\nExpected: (%T) - %v\nReceived: (%T) - %v\n", expected, expected, actual, actual)
	}

}

func responseWriter() *ResponseWriterStub {
	return &ResponseWriterStub{
		header: make(map[string][]string),
	}
}

type ResponseWriterStub struct {
	header     map[string][]string
	headerInt  int
	writtenOut string
}

func (w *ResponseWriterStub) Header() http.Header {
	return w.header
}
func (w *ResponseWriterStub) Write(b []byte) (int, error) {
	w.writtenOut = string(b)
	return 0, nil
}
func (w *ResponseWriterStub) WriteHeader(i int) {
	w.headerInt = i
}

func request(path string) *http.Request {
	if rawUrl, err := url.Parse(path); err != nil {
		panic(err)
	} else {
		return &http.Request{
			URL: rawUrl,
		}
	}
}
