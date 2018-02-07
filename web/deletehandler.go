package web

import (
	"net/http"
	"strings"
)

func deleteHandler(w http.ResponseWriter, r *http.Request, store Storer) {
	splitURL := strings.Split(r.URL.Path, "/")
	if len(splitURL) != 3 {
		return
	}

	if err := store.Delete(splitURL[2]); err != nil {
		panic(err)
	}
}
