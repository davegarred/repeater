package sql

import (
	"bufio"
	"io"
)

type Lexer struct {
	r *bufio.Reader
}

func NewLexer(r io.Reader) *Lexer {
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
