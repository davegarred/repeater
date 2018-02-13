package server

import (
	"testing"
	"fmt"
	"github.com/davegarred/repeater/persist"
)

func TestServer_Pushfile(t *testing.T) {
	store := persist.NewMemStore()
	server := NewServer(store)
	fmt.Printf("%+v\n", server)
	//server.Start(defaultPort)
}