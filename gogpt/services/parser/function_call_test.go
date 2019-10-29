package parser

import (
	"testing"

	"github.com/alexgarzao/gogpt-interpreter/gogpt/services/lexer"
	"github.com/stretchr/testify/assert"
)

func TestImprimaFunctionCall(t *testing.T) {
	l := lexer.New(`imprima("Hello!")`)
	p := New(l)
	pr := p.parserFunctionCall()
	assert.Equal(t, true, pr.Parsed)
}

func TestFunctionCallWithoutArguments(t *testing.T) {
	l := lexer.New(`myfunction()`)
	p := New(l)
	pr := p.parserFunctionCall()
	assert.Equal(t, true, pr.Parsed)
}

func TestInvalidFunctionCallWithoutArguments(t *testing.T) {
	l := lexer.New(`myfunction(,)`)
	p := New(l)
	pr := p.parserFunctionCall()
	assert.Equal(t, false, pr.Parsed)
	assert.EqualError(t, pr.Err, "Expected RPAREN")
}

func TestFunctionCallWithNArguments(t *testing.T) {
	l := lexer.New(`myfunction("A", "B", "C")`)
	p := New(l)
	pr := p.parserFunctionCall()
	assert.Equal(t, true, pr.Parsed)
}

func TestInvalidFunctionCallWithNArguments(t *testing.T) {
	l := lexer.New(`myfunction("A", "B", )`)
	p := New(l)
	pr := p.parserFunctionCall()
	assert.Equal(t, false, pr.Parsed)
	assert.EqualError(t, pr.Err, "Expected EXPR")

	l = lexer.New(`myfunction("A" "B", "C")`)
	p = New(l)
	pr = p.parserFunctionCall()
	assert.Equal(t, false, pr.Parsed)
	assert.EqualError(t, pr.Err, "Expected RPAREN")
}

func TestFunctionCallWithStringAndIdentifiers(t *testing.T) {
	l := lexer.New(`imprima("Olá ", nome)`)
	p := New(l)
	pr := p.parserFunctionCall()
	assert.Equal(t, true, pr.Parsed)
}
