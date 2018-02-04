package web

import (
	"net/http"
	"strings"
)

func deleteHandler(w http.ResponseWriter, r *http.Request, store Storer) {
	splitUrl := strings.Split(r.URL.Path, "/")
	if len(splitUrl) != 3 {
		return
	}

	if err := store.Delete(splitUrl[2]); err != nil {
		panic(err)
	}
}
