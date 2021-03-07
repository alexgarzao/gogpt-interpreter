package parser

import (
	"testing"

	"github.com/alexgarzao/gogpt-interpreter/pkg/domain/usecases/lexer"
	"github.com/stretchr/testify/assert"
)

func TestValidAlgorithmDeclaration(t *testing.T) {
	alg := `algoritmo olá_mundo;`
	l := lexer.New(alg)
	p := New(l)
	err := p.parserAlgorithmDeclaration()
	assert.Nil(t, err)
}

func TestInvalidAlgorithmDeclarationWithoutTokenAlgoritmo(t *testing.T) {
	alg := `olá_mundo;`
	l := lexer.New(alg)
	p := New(l)
	err := p.parserAlgorithmDeclaration()
	assert.EqualError(t, err, "Expected ALGORITHM")
}

func TestInvalidAlgorithmDeclarationWithoutId(t *testing.T) {
	alg := `algoritmo ;`
	l := lexer.New(alg)
	p := New(l)
	err := p.parserAlgorithmDeclaration()
	assert.EqualError(t, err, "Expected IDENT")
}

func TestInvalidAlgorithmDeclarationWithoutSemicolon(t *testing.T) {
	alg := `algoritmo olá_mundo`
	l := lexer.New(alg)
	p := New(l)
	err := p.parserAlgorithmDeclaration()
	assert.EqualError(t, err, "Expected SEMICOLON")
}
