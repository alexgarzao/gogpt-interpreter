package syntax

import (
	"testing"

	lexer "github.com/alexgarzao/gogpt-interpreter/gogpt/entities/lexical_analyzer"
	"github.com/stretchr/testify/assert"
)

func TestImprimaFunctionCall(t *testing.T) {
	l := lexer.NewLexer(`imprima("Hello!")`)
	s := NewAlgorithm(l)
	assert.Equal(t, s.ParserFunctionCall(), true)
}

func TestFunctionCallWithoutArguments(t *testing.T) {
	l := lexer.NewLexer(`myfunction()`)
	s := NewAlgorithm(l)
	assert.Equal(t, s.ParserFunctionCall(), true)
}

func TestInvalidFunctionCallWithoutArguments(t *testing.T) {
	l := lexer.NewLexer(`myfunction(,)`)
	s := NewAlgorithm(l)
	assert.Equal(t, s.ParserFunctionCall(), false)
}

func TestFunctionCallWithNArguments(t *testing.T) {
	l := lexer.NewLexer(`myfunction("A", "B", "C")`)
	s := NewAlgorithm(l)
	assert.Equal(t, s.ParserFunctionCall(), true)
}

func TestInvalidFunctionCallWithNArguments(t *testing.T) {
	l := lexer.NewLexer(`myfunction("A", "B", )`)
	s := NewAlgorithm(l)
	assert.Equal(t, s.ParserFunctionCall(), false)

	l = lexer.NewLexer(`myfunction("A" "B", "C")`)
	assert.Equal(t, s.ParserFunctionCall(), false)
}
