package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/davegarred/repeater/persist"
	"github.com/google/uuid"
)

type handler func(http.ResponseWriter, *http.Request, persist.Store)

func storeHandler(w http.ResponseWriter, r *http.Request, store persist.Store) {
	params := r.URL.Query()
	serialized, err := json.Marshal(params)
	if err != nil {
		fmt.Printf("Could not serialize type %T: %v\n", err, err)
		return
	}

	key := uuid.New().String()
	if err := store.Store(key, string(serialized)); err != nil {
		panic("no error handling implemented on store yet")
	}
	storeRespond(key, w)
}

func storeRespond(key string, w http.ResponseWriter) {
	w.Header().Set("X-Document-Id", key)
	fmt.Fprintf(w, "%v", key)
}

func retrieveHandler(w http.ResponseWriter, r *http.Request, store persist.Store) {
	splitUrl := strings.Split(r.URL.Path, "/")
	if len(splitUrl) != 3 {
		return
	}

	val, err := store.Retrieve(splitUrl[2])
	if val == "" || err != nil {
		http.NotFound(w, r)
		return
	}

	w.Header().Set("content-type", "application/json")
	fmt.Fprintf(w, "%v", val)
}
