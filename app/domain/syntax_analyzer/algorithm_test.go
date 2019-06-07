package analyzer

import (
	"testing"

	"github.com/stretchr/testify/assert"

	lexer "github.com/alexgarzao/gpt-interpreter/app/domain/lexical_analyzer"
)

func TestValidEmptyProgram(t *testing.T) {
	c :=
		`algoritmo olá_mundo;
início
fim`
	l := lexer.NewLexer(c)
	p := NewProgram()
	assert.Equal(t, true, p.TryToParse(l))
}

func TestValidHelloWorldAlgorithm(t *testing.T) {
	c :=
		`algoritmo olá_mundo;
início
	imprima("Olá mundo!");
fim`
	l := lexer.NewLexer(c)
	p := NewProgram()
	assert.Equal(t, true, p.TryToParse(l))
}

func TestValidHelloWorldWithTwoSentences(t *testing.T) {
	c :=
		`algoritmo olá_mundo;
início
	imprima("Olá...");
	imprima("Mundo!");
fim`
	l := lexer.NewLexer(c)
	p := NewProgram()
	assert.Equal(t, true, p.TryToParse(l))
}
