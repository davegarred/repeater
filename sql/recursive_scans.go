package sql

import (
	"bytes"
	"strings"
)

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
