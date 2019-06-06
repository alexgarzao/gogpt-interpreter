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
