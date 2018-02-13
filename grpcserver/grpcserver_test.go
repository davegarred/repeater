package server

import (
	"testing"
	"fmt"
	"github.com/davegarred/repeater/persist"
)

func TestServer_Pushfile(t *testing.T) {
	store := persist.NewMemStore()
	fmt.Printf("%+v\n", store)
	//StartGRPCServer(store, defaultPort)
}