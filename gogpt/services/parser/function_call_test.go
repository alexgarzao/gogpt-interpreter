package parser

import (
	"testing"

	"github.com/alexgarzao/gogpt-interpreter/gogpt/services/lexer"
	"github.com/stretchr/testify/assert"
)

func TestImprimaFunctionCall(t *testing.T) {
	l := lexer.NewLexer(`imprima("Hello!")`)
	s := NewAlgorithm(l)
	pr := s.ParserFunctionCall()
	assert.Equal(t, true, pr.Parsed)
}

func TestFunctionCallWithoutArguments(t *testing.T) {
	l := lexer.NewLexer(`myfunction()`)
	s := NewAlgorithm(l)
	pr := s.ParserFunctionCall()
	assert.Equal(t, true, pr.Parsed)
}

func TestInvalidFunctionCallWithoutArguments(t *testing.T) {
	l := lexer.NewLexer(`myfunction(,)`)
	s := NewAlgorithm(l)
	pr := s.ParserFunctionCall()
	assert.Equal(t, false, pr.Parsed)
	assert.EqualError(t, pr.Err, "Expected RPAREN")
}

func TestFunctionCallWithNArguments(t *testing.T) {
	l := lexer.NewLexer(`myfunction("A", "B", "C")`)
	s := NewAlgorithm(l)
	pr := s.ParserFunctionCall()
	assert.Equal(t, true, pr.Parsed)
}

func TestInvalidFunctionCallWithNArguments(t *testing.T) {
	l := lexer.NewLexer(`myfunction("A", "B", )`)
	s := NewAlgorithm(l)
	pr := s.ParserFunctionCall()
	assert.Equal(t, false, pr.Parsed)
	assert.EqualError(t, pr.Err, "Expected EXPR")

	l = lexer.NewLexer(`myfunction("A" "B", "C")`)
	s = NewAlgorithm(l)
	pr = s.ParserFunctionCall()
	assert.Equal(t, false, pr.Parsed)
	assert.EqualError(t, pr.Err, "Expected RPAREN")
}

func TestFunctionCallWithStringAndIdentifiers(t *testing.T) {
	l := lexer.NewLexer(`imprima("Olá ", nome)`)
	s := NewAlgorithm(l)
	pr := s.ParserFunctionCall()
	assert.Equal(t, true, pr.Parsed)
}
