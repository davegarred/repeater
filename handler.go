package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/davegarred/repeater/persist"
)

type handler func(http.ResponseWriter, *http.Request, persist.Store)

func storeHandler(w http.ResponseWriter, r *http.Request, store persist.Store) {
	params := r.URL.Query()
	serialized, err := json.Marshal(params)
	if err != nil {
		fmt.Printf("Could not serialize type %T: %v\n", err, err)
		return
	}
	key := store.Store(string(serialized))
	w.Header().Set("X-Document-Id", key)
	fmt.Fprintf(w, "%v", key)
}

func retrieveHandler(w http.ResponseWriter, r *http.Request, store persist.Store) {
	splitUrl := strings.Split(r.URL.Path, "/")
	if len(splitUrl) != 3 {
		return
	}

	val := store.Retrieve(splitUrl[2])
	if val == "" {
		http.NotFound(w, r)
		return
	}

	w.Header().Set("content-type", "application/json")
	fmt.Fprintf(w, "%v", val)
}
