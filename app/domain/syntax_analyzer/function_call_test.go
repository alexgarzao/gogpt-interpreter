package analyzer

import (
	"testing"

	"github.com/stretchr/testify/assert"

	lexer "github.com/alexgarzao/gpt-interpreter/app/domain/lexical_analyzer"
)

func TestImprimaFunctionCall(t *testing.T) {
	l := lexer.NewLexer(`imprima("Hello!")`)
	s := NewFunctionCall()
	assert.Equal(t, s.IsValid(l), true)
}

func TestFunctionCallWithoutArguments(t *testing.T) {
	l := lexer.NewLexer(`myfunction()`)
	s := NewFunctionCall()
	assert.Equal(t, s.IsValid(l), true)
}

func TestInvalidFunctionCallWithoutArguments(t *testing.T) {
	l := lexer.NewLexer(`myfunction(,)`)
	s := NewFunctionCall()
	assert.Equal(t, s.IsValid(l), false)
}

func TestFunctionCallWithNArguments(t *testing.T) {
	l := lexer.NewLexer(`myfunction("A", "B", "C")`)
	s := NewFunctionCall()
	assert.Equal(t, s.IsValid(l), true)
}

func TestInvalidFunctionCallWithNArguments(t *testing.T) {
	l := lexer.NewLexer(`myfunction("A", "B", )`)
	s := NewFunctionCall()
	assert.Equal(t, s.IsValid(l), false)

	l = lexer.NewLexer(`myfunction("A" "B", "C")`)
	assert.Equal(t, s.IsValid(l), false)
}
