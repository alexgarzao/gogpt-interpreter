package parser

import (
	"testing"

	"github.com/alexgarzao/gogpt-interpreter/pkg/domain/usecases/lexer"
	"github.com/stretchr/testify/assert"
)

func TestImprimaFunctionCall(t *testing.T) {
	l := lexer.New(`imprima("Hello!")`)
	p := New(l)
	err := p.parserFunctionCall()
	assert.Nil(t, err)
}

func TestFunctionCallWithoutArguments(t *testing.T) {
	l := lexer.New(`myfunction()`)
	p := New(l)
	err := p.parserFunctionCall()
	assert.Nil(t, err)
}

func TestInvalidFunctionCallWithoutArguments(t *testing.T) {
	l := lexer.New(`myfunction(,)`)
	p := New(l)
	err := p.parserFunctionCall()
	assert.EqualError(t, err, "Expected RPAREN")
}

func TestFunctionCallWithNArguments(t *testing.T) {
	l := lexer.New(`myfunction("A", "B", "C")`)
	p := New(l)
	err := p.parserFunctionCall()
	assert.Nil(t, err)
}

func TestInvalidFunctionCallWithNArguments(t *testing.T) {
	l := lexer.New(`myfunction("A", "B", )`)
	p := New(l)
	err := p.parserFunctionCall()
	assert.EqualError(t, err, "Expected EXPR")

	l = lexer.New(`myfunction("A" "B", "C")`)
	p = New(l)
	err = p.parserFunctionCall()
	assert.EqualError(t, err, "Expected RPAREN")
}

func TestFunctionCallWithStringAndIdentifiers(t *testing.T) {
	l := lexer.New(`imprima("Olá ", nome)`)
	p := New(l)
	err := p.parserFunctionCall()
	assert.Nil(t, err)
}
