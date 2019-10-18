package syntax

import (
	"testing"

	"github.com/stretchr/testify/assert"

	lexer "github.com/alexgarzao/gogpt-interpreter/gogpt/entities/lexical_analyzer"
)

func TestValidAlgorithmDeclaration(t *testing.T) {
	c := `algoritmo olá_mundo;`
	l := lexer.NewLexer(c)
	p := NewAlgorithm(l)
	pr := p.ParserAlgorithmDeclaration()
	assert.Equal(t, true, pr.Parsed)
}

func TestInvalidAlgorithmDeclarationWithoutTokenAlgoritmo(t *testing.T) {
	c := `olá_mundo;`
	l := lexer.NewLexer(c)
	p := NewAlgorithm(l)
	pr := p.ParserAlgorithmDeclaration()
	assert.Equal(t, false, pr.Parsed)
	assert.NoError(t, pr.Err)
}

func TestInvalidAlgorithmDeclarationWithoutId(t *testing.T) {
	c := `algoritmo ;`
	l := lexer.NewLexer(c)
	p := NewAlgorithm(l)
	pr := p.ParserAlgorithmDeclaration()
	assert.Equal(t, false, pr.Parsed)
	assert.EqualError(t, pr.Err, "Expected IDENT")
}

func TestInvalidAlgorithmDeclarationWithoutSemicolon(t *testing.T) {
	c := `algoritmo olá_mundo`
	l := lexer.NewLexer(c)
	p := NewAlgorithm(l)
	pr := p.ParserAlgorithmDeclaration()
	assert.Equal(t, false, pr.Parsed)
	assert.EqualError(t, pr.Err, "Expected SEMICOLON")
}