package web

import (
	"net/http"
	"strings"

	"github.com/davegarred/repeater/persist"
)

func deleteHandler(w http.ResponseWriter, r *http.Request, store persist.Store) {
	splitUrl := strings.Split(r.URL.Path, "/")
	if len(splitUrl) != 3 {
		return
	}

	if err := store.Delete(splitUrl[2]); err != nil {
		panic(err)
	}
}
