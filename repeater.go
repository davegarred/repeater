package main

import (
	"github.com/davegarred/repeater/persist"
	"github.com/davegarred/repeater/web"
)

func main() {
	store := persist.NewStore()
	web.Start(store)
}
