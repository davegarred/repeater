package web

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/davegarred/repeater/persist"
	"github.com/google/uuid"
)

func parseAndSerialize(params map[string][]string) (string, error) {
	data := make(map[string]string, len(params))
	for k, v := range params {
		if len(v) > 0 {
			data[k] = v[0]
		}
	}
	serialized, err := json.Marshal(data)
	if err != nil {
		fmt.Printf("Could not serialize type %T: %v\n", err, err)
		return "", err
	}
	return string(serialized), nil
}

func storeHandler(w http.ResponseWriter, r *http.Request, store persist.Store) {
	splitUrl := strings.Split(r.URL.Path, "/")
	urlSegments := len(splitUrl)

	var key string
	switch {
	case urlSegments < 2:
		return
	case urlSegments == 2:
		key = uuid.New().String()
	case urlSegments > 2:
		key = splitUrl[2]
	}

	data, err := parseAndSerialize(r.URL.Query())
	if err != nil {
		panic(err)
	}

	if err := store.Store(key, data); err != nil {
		if err == persist.KEY_CONFLICT {
			w.WriteHeader(400)
			fmt.Fprintf(w, "%v", "Document already exists with this key")
			return
		} else {
			panic("no error handling implemented on store yet")
		}

	}
	w.Header().Set("X-Document-Id", key)
	fmt.Fprintf(w, "%v", key)
}