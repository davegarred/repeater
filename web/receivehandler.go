package web

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/davegarred/repeater/persist"
)

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
