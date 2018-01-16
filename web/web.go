package web

import (
	"fmt"
	"net/http"
	"path"

	"github.com/davegarred/repeater/persist"
)

var store persist.Store

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
	http.NotFound(w, r)
}

func Start(s persist.Store) {
	store = s
	pathResolver := &pathResolver{handlers: make(map[string]handler)}
	pathResolver.Add("GET /store", storeHandler)
	pathResolver.Add("GET /store/*", storeHandler)
	pathResolver.Add("GET /retrieve/*", retrieveHandler)
	err := http.ListenAndServe(":8000", pathResolver)
	if err != nil {
		fmt.Printf("%v\n", err)
	}
}
