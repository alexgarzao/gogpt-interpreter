package lexer

import (
	"strings"
	"unicode"
	"unicode/utf8"
)

// Lexer keeps the infos to do the lexical analysis.
type Lexer struct {
	input        string
	currentPos   int
	backTracking int
}

// New creates a new Lexer.
func New(input string) *Lexer {
	return &Lexer{
		input:      input,
		currentPos: 0,
	}
}

// NextToken returns the next token.
func (l *Lexer) NextToken() *Token {
	var ch rune

	// Ignore useless chars.
	for {
		ch, _ = l.nextChar()
		if ch == 0 {
			return &Token{EOF, EOF}
		}
		if isDelimiterChar(ch) {
			token := string(ch)
			if ch == ':' {
				nextCh, _ := l.nextChar()
				if nextCh == '=' {
					token = ":="
				} else {
					l.currentPos--
				}
			}
			return defineToken(token)
		}
		if !isUselessChar(ch) {
			break
		}
	}

	// Test if is an Identifier or Keyword.
	if token := l.tryIDOrKeyword(ch); token != nil {
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

// GetNextTokenIf get the next token if match a specific token type.
func (l *Lexer) GetNextTokenIf(expectedType string) *Token {
	backTracking := l.currentPos
	token := l.NextToken()
	if token.Type == expectedType {
		return token
	}

	l.currentPos = backTracking
	return nil
}

// GetNextsTokensIf get the nexts tokens if both match a specifics tokens types.
func (l *Lexer) GetNextsTokensIf(expectedType1, expectedType2 string) (*Token, *Token) {
	backTracking := l.currentPos

	token1 := l.NextToken()
	if token1.Type != expectedType1 {
		l.currentPos = backTracking
		return nil, nil
	}

	token2 := l.NextToken()
	if token2.Type != expectedType2 {
		l.currentPos = backTracking
		return nil, nil
	}

	return token1, token2
}

func (l *Lexer) tryIDOrKeyword(ch rune) *Token {
	// [A-Z|_|-] ([A-Z]|[0-9]|_|-)*
	token := string(ch)
	if (strings.ToUpper(token) < "A" || strings.ToUpper(token) > "Z") && token != "_" && token != "-" {
		return nil
	}

	for {
		ch, size := l.nextChar()
		if ch == 0 {
			break
		}
		if !unicode.IsLetter(ch) && string(ch) != "_" && string(ch) != "-" && !(string(ch) >= "0" && string(ch) <= "9") {
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
