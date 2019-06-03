package lexer

import (
	"strings"
)

const (
	EOF     = "EOF"
	INVALID = "INVALID"

	// Keywords
	ALGORITMO = "ALGORITMO"
	INICIO    = "INÍCIO"
	FIM       = "FIM"

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"
	LPAREN    = "("
	RPAREN    = ")"

	// Literals
	INT    = "INT"
	STRING = "STRING"

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
	return ch == ',' || ch == ';' || ch == '(' || ch == ')'
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
	}

	switch strings.ToUpper(token) {
	case ALGORITMO:
		return &Token{ALGORITMO, ALGORITMO}
	case INICIO:
		return &Token{INICIO, INICIO}
	case FIM:
		return &Token{FIM, FIM}
	}

	if token[0] == '"' {
		return &Token{STRING, token}
	}

	return &Token{IDENT, token}
}
