package sql

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"strings"
)

type Token string

const (
	ILLEGAL Token = "ILLEGAL"
	SELECT  Token = "SELECT"
	FROM    Token = "FROM"
	WHERE   Token = "WHERE"
	AND     Token = "AND"
	OR      Token = "OR"

	EOF Token = "EOF"
	WS  Token = "WS"

	IDENT  Token = "IDENT"
	NUMBER Token = "NUMBER"

	ASTERISK Token = "ASTERISK"
	EQUALS   Token = "EQUALS"
	COMMA    Token = "COMMA"
	PERIOD   Token = "PERIOD"

	eof = rune(0)
)

type Lexer struct {
	r *bufio.Reader
}

func NewLexer(r io.Reader) *Lexer {
	fmt.Println("lexer starts")
	return &Lexer{bufio.NewReader(r)}
}

func (s *Lexer) read() rune {
	rune, _, err := s.r.ReadRune()
	if err != nil {
		return eof
	}
	return rune
}

func (s *Lexer) unread() {
	s.r.UnreadRune()
}

func (s *Lexer) Scan() (Token, string) {
	ch := s.read()

	if isWhitespace(ch) {
		s.unread()
		return s.scanWhiteSpace()
	} else if isLetter(ch) {
		s.unread()
		return s.scanIdent()
	} else if isDigit(ch) {
		s.unread()
		return s.scanNumber()
	}
	switch ch {
	case eof:
		return EOF, ""
	case '=':
		return EQUALS, string(ch)
	case '*':
		return ASTERISK, string(ch)
	case ',':
		return COMMA, string(ch)
	}

	return ILLEGAL, string(ch)
}

func (s *Lexer) scanWhiteSpace() (Token, string) {
	var buf bytes.Buffer
	buf.WriteRune(s.read())

	for {
		if ch := s.read(); ch == eof {
			break
		} else if isWhitespace(ch) {
			buf.WriteRune(ch)
		} else {
			s.unread()
			break
		}
	}
	return WS, buf.String()
}

func (s *Lexer) scanIdent() (Token, string) {
	var buf bytes.Buffer
	buf.WriteRune(s.read())

	for {
		if ch := s.read(); ch == eof {
			break
		} else if isLetter(ch) || isDigit(ch) {
			buf.WriteRune(ch)
		} else {
			s.unread()
			break
		}
	}

	stringVal := buf.String()
	switch strings.ToUpper(stringVal) {
	case "SELECT":
		return SELECT, stringVal
	case "FROM":
		return FROM, stringVal
	case "WHERE":
		return WHERE, stringVal
	case "AND":
		return AND, stringVal
	case "OR":
		return OR, stringVal
	}
	return IDENT, stringVal
}

func (s *Lexer) scanNumber() (Token, string) {
	var buf bytes.Buffer
	buf.WriteRune(s.read())

	for {
		if ch := s.read(); ch == eof {
			break
		} else if isDigit(ch) || ch == '.' {
			buf.WriteRune(ch)
		} else {
			s.unread()
			break
		}
	}
	return NUMBER, buf.String()
}

func isWhitespace(ch rune) bool {
	return ch == ' ' || ch == '\t' || ch == '\n'
}
func isLetter(ch rune) bool {
	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') || ch == '_'
}
func isDigit(ch rune) bool {
	return (ch >= '0' && ch <= '9')
}
