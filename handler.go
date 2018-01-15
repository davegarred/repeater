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

func configuredStoreHandler(w http.ResponseWriter, r *http.Request, store persist.Store) {
	params := r.URL.Query()
	for k, v := range params {
		fmt.Printf("%v=%v\n", k, v)
	}
	serialized, err := json.Marshal(params)
	if err != nil {
		t := reflect.TypeOf(params).Name()
		fmt.Printf("Could not serialize type %v: %v\n", t, err)
		return
	}
	key := store.Store(string(serialized))
	fmt.Fprintf(w, "%v", key.URN())
}

func configuredRetrieveHandler(w http.ResponseWriter, r *http.Request, store persist.Store) {
	splitUrl := strings.Split(r.URL.Path, "/")
	if len(splitUrl) < 3 {
		return
	}
	//	id := r.URL.Query()["id"]
	id := splitUrl[2]
	key, err := uuid.Parse(id)
	if err != nil {
		fmt.Printf("Error parsing UUID - %v, %v \n", id, err)
		return
	}
	fmt.Printf("key=%v\n", key)
	val := store.Retrieve(key)
	if val == "" {
		return
	}
	fmt.Fprintf(w, "%v", val)
}
