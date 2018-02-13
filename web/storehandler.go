package web

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"io/ioutil"

	"gopkg.in/h2non/filetype.v1"
	"github.com/google/uuid"
	"github.com/davegarred/repeater/persist"
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

func getStoreHandler(w http.ResponseWriter, r *http.Request, store Storer) {
	splitURL := strings.Split(r.URL.Path, "/")
	urlSegments := len(splitURL)

	var key string
	switch {
	case urlSegments < 2:
		return
	case urlSegments == 2:
		key = uuid.New().String()
	case urlSegments > 2:
		key = splitURL[2]
	}

	data, err := parseAndSerialize(r.URL.Query())
	if err != nil {
		panic(err)
	}

	if err := store.Store("application/json", key, data); err != nil {
		if err == persist.KeyConflict {
			w.WriteHeader(400)
			fmt.Fprintf(w, "Document already exists with this key")
			return
		}
		w.WriteHeader(500)
		fmt.Fprintf(w, "Unknown storage error encountered: %v", err)
		return
	}
	w.Header().Set("X-Document-Id", key)
	fmt.Fprintf(w, "%v", key)
}

func postStoreHandler(w http.ResponseWriter, r *http.Request, store Storer) {
	key := uuid.New().String()

	data,err := ioutil.ReadAll(r.Body)
	if err != nil {
		return
	}
	defer r.Body.Close()

	mimetype := r.Header.Get("Content-Type")
	if mimetype == "" {
		kind,unknown := filetype.Match(data)
		if unknown == nil && kind.MIME.Value != "" {
			mimetype = kind.MIME.Value
		} else {
			mimetype = "application/octet-stream"
		}
	}

	if err := store.Store(mimetype, key, string(data)); err != nil {
		if err == persist.KeyConflict {
			w.WriteHeader(400)
			fmt.Fprintf(w, "Document already exists with this key")
			return
		}
		w.WriteHeader(500)
		fmt.Fprintf(w, "Unknown storage error encountered: %v", err)
		return
	}

	w.Header().Set("X-Document-Id", key)
	fmt.Fprintf(w, "%v", key)
}
