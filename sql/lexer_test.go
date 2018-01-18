package sql

import (
	"fmt"
	"strings"
	"testing"
)

const SIMPLE_SELECT = "select * from some_table where id=11"

func TestSomething(t *testing.T) {
	showParse(t, SIMPLE_SELECT)
}
func showParse(t *testing.T, v string) {
	fmt.Println(v)
	lexer := NewLexer(strings.NewReader(v))
	for {
		tok, str := lexer.Scan()
		if tok != WS {
			fmt.Printf("%v\t%v\n", string(tok), str)
		}
		if tok == EOF {
			break
		}
	}

}
