package web

import (
	"fmt"
	"net/http"
	"path"

	"github.com/davegarred/repeater/util"
)

type Storer interface {
	Store(string, string, string) error
	Retrieve(string) (string, error)
	Delete(string) error
}

var store Storer

type handler func(http.ResponseWriter, *http.Request, Storer)

type pathResolver struct {
	handlers map[string]handler
}

func (p *pathResolver) Add(path string, handler handler) {
	p.handlers[path] = handler
}

func (p *pathResolver) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	signature := r.Method + " " + r.URL.Path
	for pattern, handler := range p.handlers {
		if ok, err := path.Match(pattern, signature); ok && err == nil {
			handler(w, r, store)
			return
		} else if err != nil {
			panic(err)
		}
		//		fmt.Printf("missed pattern: %v\n", pattern)
	}
	//	fmt.Printf("missed grab:    %v\n", signature)

	util.Log("Request path could not be matched: %v", signature)
	http.NotFound(w, r)
}

func Start(s Storer) {
	store = s
	pathResolver := &pathResolver{handlers: make(map[string]handler)}
	pathResolver.Add("GET /store", storeHandler)
	pathResolver.Add("GET /store/*", storeHandler)
	pathResolver.Add("GET /retrieve/*", retrieveHandler)
	pathResolver.Add("GET /delete/*", deleteHandler)
	err := http.ListenAndServe(":8000", pathResolver)
	if err != nil {
		fmt.Printf("%v\n", err)
	}
}
