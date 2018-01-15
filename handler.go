package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"strings"

	"github.com/davegarred/repeater/persist"
	"github.com/google/uuid"
)

type handler func(http.ResponseWriter, *http.Request, persist.Store)

func storeHandler(w http.ResponseWriter, r *http.Request, store persist.Store) {
	params := r.URL.Query()
	serialized, err := json.Marshal(params)
	if err != nil {
		t := reflect.TypeOf(params).Name()
		fmt.Printf("Could not serialize type %v: %v\n", t, err)
		return
	}
	key := store.Store(string(serialized))
	w.Header().Set("X-Document-Id", key.String())
	fmt.Fprintf(w, "%v", key.String())
}

func retrieveHandler(w http.ResponseWriter, r *http.Request, store persist.Store) {
	splitUrl := strings.Split(r.URL.Path, "/")
	if len(splitUrl) != 3 {
		return
	}

	id := splitUrl[2]
	key, err := uuid.Parse(id)
	if err != nil {
		fmt.Printf("Error parsing UUID - %v, %v \n", id, err)
		return
	}

	val := store.Retrieve(key)
	if val == "" {
		http.NotFound(w, r)
		return
	}

	w.Header().Set("content-type", "application/json")
	fmt.Fprintf(w, "%v", val)
}
