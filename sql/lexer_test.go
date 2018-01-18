package sql

import (
	"fmt"
	"strings"
	"testing"
)

const SIMPLE_SELECT = "seLEct *,id, name from some_tAble Where id=11"
const SIMPLE_SELECT_2 = "select * from some_table where id=11"

func TestSomething(t *testing.T) {
	showParse(t, SIMPLE_SELECT)
	showParse(t, SIMPLE_SELECT_2)
}
func showParse(t *testing.T, v string) {
	fmt.Println(v)
	lexer := NewLexer(strings.NewReader(v))
	for {
		tok, str := lexer.Scan()
		if tok != WS {
			fmt.Printf("%v(%v)   ", string(tok), str)
		}
		if tok == EOF {
			break
		}
	}
	fmt.Println()

}
