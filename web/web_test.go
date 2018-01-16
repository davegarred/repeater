package web

import (
	"testing"

	"github.com/davegarred/repeater/persist"
)

func Test(t *testing.T) {
	persist.NewMemStore()

}
