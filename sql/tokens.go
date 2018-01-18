package sql

type Token string

const (
	ILLEGAL Token = "ILLEGAL"

	SELECT = "SELECT"
	DELETE = "DELETE"
	UPDATE = "UPDATE"
	INSERT = "INSERT"
	FROM   = "FROM"
	WHERE  = "WHERE"
	AND    = "AND"
	OR     = "OR"

	EOF = "EOF"
	WS  = "WS"

	IDENT  = "IDENT"
	NUMBER = "NUMBER"

	ASTERISK = "ASTERISK"
	EQUALS   = "EQUALS"
	COMMA    = "COMMA"
	PERIOD   = "PERIOD"
)

const (
	eof = rune(0)
)
