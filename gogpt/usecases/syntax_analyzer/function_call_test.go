package syntax

import (
	"testing"

	lexer "github.com/alexgarzao/gogpt-interpreter/gogpt/entities/lexical_analyzer"
	"github.com/stretchr/testify/assert"
)

func TestImprimaFunctionCall(t *testing.T) {
	l := lexer.NewLexer(`imprima("Hello!")`)
	s := NewAlgorithm()
	assert.Equal(t, s.ParserFunctionCall(l), true)
}

func TestFunctionCallWithoutArguments(t *testing.T) {
	l := lexer.NewLexer(`myfunction()`)
	s := NewAlgorithm()
	assert.Equal(t, s.ParserFunctionCall(l), true)
}

func TestInvalidFunctionCallWithoutArguments(t *testing.T) {
	l := lexer.NewLexer(`myfunction(,)`)
	s := NewAlgorithm()
	assert.Equal(t, s.ParserFunctionCall(l), false)
}

func TestFunctionCallWithNArguments(t *testing.T) {
	l := lexer.NewLexer(`myfunction("A", "B", "C")`)
	s := NewAlgorithm()
	assert.Equal(t, s.ParserFunctionCall(l), true)
}

func TestInvalidFunctionCallWithNArguments(t *testing.T) {
	l := lexer.NewLexer(`myfunction("A", "B", )`)
	s := NewAlgorithm()
	assert.Equal(t, s.ParserFunctionCall(l), false)

	l = lexer.NewLexer(`myfunction("A" "B", "C")`)
	assert.Equal(t, s.ParserFunctionCall(l), false)
}
