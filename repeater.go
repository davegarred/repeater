package main

import (
	"fmt"
	"net/http"

	"github.com/davegarred/repeater/persist"
)

var store *persist.Store

func storeHandler(w http.ResponseWriter, r *http.Request) {
	configuredStoreHandler(w, r, store)
}
func retrieveHandler(w http.ResponseWriter, r *http.Request) {
	configuredRetrieveHandler(w, r, store)
}

func main() {
	store = persist.NewStore()
	http.HandleFunc("/store", storeHandler)
	http.HandleFunc("/retrieve/", retrieveHandler)
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		fmt.Printf("%v\n", err)
	}

}
