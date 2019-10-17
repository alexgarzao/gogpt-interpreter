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
	assert.Equal(t, true, p.ParserAlgorithmDeclaration())
}

func TestInvalidAlgorithmDeclarationWithoutTokenAlgoritmo(t *testing.T) {
	c := `olá_mundo;`
	l := lexer.NewLexer(c)
	p := NewAlgorithm(l)
	assert.Equal(t, false, p.Parser())
}

func TestInvalidAlgorithmDeclarationWithoutId(t *testing.T) {
	c := `algoritmo ;`
	l := lexer.NewLexer(c)
	p := NewAlgorithm(l)
	assert.Equal(t, false, p.Parser())
}

func TestInvalidAlgorithmDeclarationWithoutSemicolon(t *testing.T) {
	c := `algoritmo olá_mundo`
	l := lexer.NewLexer(c)
	p := NewAlgorithm(l)
	assert.Equal(t, false, p.Parser())
}
