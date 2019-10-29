package parser

import (
	"testing"

	"github.com/alexgarzao/gogpt-interpreter/gogpt/services/lexer"
	"github.com/stretchr/testify/assert"
)

func TestValidAlgorithmDeclaration(t *testing.T) {
	c := `algoritmo olá_mundo;`
	l := lexer.New(c)
	p := New(l)
	pr := p.parserAlgorithmDeclaration()
	assert.Equal(t, true, pr.Parsed)
}

func TestInvalidAlgorithmDeclarationWithoutTokenAlgoritmo(t *testing.T) {
	c := `olá_mundo;`
	l := lexer.New(c)
	p := New(l)
	pr := p.parserAlgorithmDeclaration()
	assert.Equal(t, false, pr.Parsed)
	assert.NoError(t, pr.Err)
}

func TestInvalidAlgorithmDeclarationWithoutId(t *testing.T) {
	c := `algoritmo ;`
	l := lexer.New(c)
	p := New(l)
	pr := p.parserAlgorithmDeclaration()
	assert.Equal(t, false, pr.Parsed)
	assert.EqualError(t, pr.Err, "Expected IDENT")
}

func TestInvalidAlgorithmDeclarationWithoutSemicolon(t *testing.T) {
	c := `algoritmo olá_mundo`
	l := lexer.New(c)
	p := New(l)
	pr := p.parserAlgorithmDeclaration()
	assert.Equal(t, false, pr.Parsed)
	assert.EqualError(t, pr.Err, "Expected SEMICOLON")
}
