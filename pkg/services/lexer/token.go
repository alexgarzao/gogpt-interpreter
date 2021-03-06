package lexer

import (
	"strings"
)

const (
	EOF     = "EOF"
	INVALID = "INVALID"

	// Keywords
	ALGORITHM  = "ALGORITMO"
	BLOCKBEGIN = "INÍCIO"
	BLOCKEND   = "FIM"
	VARSBEGIN  = "VARIÁVEIS"
	VARSEND    = "FIM-VARIÁVEIS"

	// Delimiters
	COMMA     = ","
	COLON     = ":"
	SEMICOLON = ";"
	LPAREN    = "("
	RPAREN    = ")"
	ATTR      = ":="

	// Literals
	INT    = "INTEIRO"
	STRING = "LITERAL"

	// Identifiers
	IDENT = "IDENT"
)

type Token struct {
	Type  string
	Value string
}

func isUselessChar(ch rune) bool {
	return ch == ' ' || ch == '\t' || ch == '\n' || ch == '\r'
}

func isDelimiterChar(ch rune) bool {
	return ch == ',' || ch == ';' || ch == '(' || ch == ')' || ch == ':' || ch == '='
}

func defineToken(token string) *Token {
	switch token {
	case "(":
		return &Token{LPAREN, LPAREN}
	case ")":
		return &Token{RPAREN, RPAREN}
	case ";":
		return &Token{SEMICOLON, SEMICOLON}
	case ",":
		return &Token{COMMA, COMMA}
	case ":":
		return &Token{COLON, COLON}
	case ATTR:
		return &Token{ATTR, ATTR}
	}

	switch strings.ToUpper(token) {
	case ALGORITHM:
		return &Token{ALGORITHM, ALGORITHM}
	case BLOCKBEGIN:
		return &Token{BLOCKBEGIN, BLOCKBEGIN}
	case BLOCKEND:
		return &Token{BLOCKEND, BLOCKEND}
	case VARSBEGIN:
		return &Token{VARSBEGIN, VARSBEGIN}
	case VARSEND:
		return &Token{VARSEND, VARSEND}
	case STRING:
		return &Token{STRING, STRING}
	case INT:
		return &Token{INT, INT}
	}

	if token[0] == '"' {
		return &Token{STRING, token}
	}

	return &Token{IDENT, token}
}
