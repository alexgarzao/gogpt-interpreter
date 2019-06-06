package lexer

import (
	"strings"
	"unicode"
	"unicode/utf8"
)

type Lexer struct {
	input      string
	currentPos int
}

func NewLexer(input string) *Lexer {
	return &Lexer{
		input:      input,
		currentPos: 0,
	}
}

func (l *Lexer) NextToken() *Token {
	var ch rune

	// Ignore useless chars.
	for {
		ch, _ = l.nextChar()
		if ch == 0 {
			return &Token{EOF, EOF}
		}
		if isDelimiterChar(ch) {
			return defineToken(string(ch))
		}
		if !isUselessChar(ch) {
			break
		}
	}

	// Test if is an Identifier or Keyword.
	if token := l.tryIdOrKeyword(ch); token != nil {
		return token
	}

	// Test if is a string literal.
	if token := l.tryString(ch); token != nil {
		return token
	}

	// Test if is a integer literal.
	if token := l.tryInt(ch); token != nil {
		return token
	}

	return &Token{INVALID, ""}
}

func (l *Lexer) tryIdOrKeyword(ch rune) *Token {
	// [A-Z|_] ([A-Z]|[0-9]|_)*
	token := string(ch)
	if (strings.ToUpper(token) < "A" || strings.ToUpper(token) > "Z") && token != "_" {
		return nil
	}

	for {
		ch, size := l.nextChar()
		if ch == 0 {
			break
		}
		if !unicode.IsLetter(ch) && string(ch) != "_" && !(string(ch) >= "0" && string(ch) <= "9") {
			l.currentPos -= size
			break
		}
		token += string(ch)
	}
	return defineToken(token)
}

func (l *Lexer) tryString(ch rune) *Token {
	// '"' .* '""
	if ch != '"' {
		return nil
	}

	token := string(ch)
	for {
		ch, _ := l.nextChar()
		if ch == 0 || ch == '\n' || ch == '\r' {
			return &Token{INVALID, token}
		}
		token += string(ch)
		if ch == '"' {
			return &Token{STRING, token[1 : len(token)-1]}
		}
	}
}

func (l *Lexer) tryInt(ch rune) *Token {
	// [0-9]+
	if ch < '0' || ch > '9' {
		return nil
	}

	token := string(ch)
	for {
		ch, size := l.nextChar()
		if ch == 0 || isUselessChar(ch) {
			return &Token{INT, strings.TrimPrefix(token, "0")}
		}
		if isDelimiterChar(ch) {
			l.currentPos -= size
			return &Token{INT, strings.TrimPrefix(token, "0")}
		}
		token += string(ch)
		if ch < '0' || ch > '9' {
			return &Token{INVALID, token}
		}
	}
}

func (l *Lexer) nextChar() (rune, int) {
	if l.currentPos == len(l.input) {
		return 0, 0
	}

	ret, size := utf8.DecodeRuneInString(l.input[l.currentPos:])
	l.currentPos += size
	return ret, size
}