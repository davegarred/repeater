package web

import (
	"fmt"
	"net/http"
	"strings"
)

func retrieveHandler(w http.ResponseWriter, r *http.Request, store Storer) {
	splitURL := strings.Split(r.URL.Path, "/")
	if len(splitURL) != 3 {
		http.NotFound(w,r)
		return
	}

	val, err := store.Retrieve(splitURL[2])
	if val == nil || err != nil {
		http.NotFound(w, r)
		return
	}

	w.Header().Set("content-type", val.Mimetype)
	fmt.Fprintf(w, "%v", val.Object)
}
