package web

import (
	"net/http"
	"net/url"
	"testing"
	"strings"

	"io/ioutil"
)

const aTestKey = "a_key"
const aTestValue = "{\"name\":\"simple json object\"}"
const aTestQuery = "name=simple%20json%20object"


func TestPathResolver(t *testing.T) {
	result := defaultPathResolver()
	assertEquals(t, 5, len(result.handlers))
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
	if rawURL, err := url.Parse(path); err != nil {
		panic(err)
	} else {
		return &http.Request{
			URL: rawURL,
		}
	}
}

func requestWithBody(path string, body string) *http.Request {
	if rawURL, err := url.Parse(path); err != nil {
		panic(err)
	} else {
		return &http.Request{
			URL: rawURL,
			Body: ioutil.NopCloser(strings.NewReader(body)),
		}
	}
}
func requestWithContentTypeAndBody(path string, contentType string, body string) *http.Request {
	headers := make(map[string][]string)
	c := make([]string,1)
	headers["Content-Type"] = c
	c[0] = contentType
	if rawURL, err := url.Parse(path); err != nil {
		panic(err)
	} else {
		return &http.Request{
			URL: rawURL,
			Header: headers,
			Body: ioutil.NopCloser(strings.NewReader(body)),
		}
	}
}
