package parser

import (
	"testing"

	"github.com/alexgarzao/gogpt-interpreter/gogpt/services/lexer"
	"github.com/stretchr/testify/assert"
)

func TestValidVarDeclaration(t *testing.T) {
	c := `		variáveis
	nome: literal;
fim-variáveis`
	l := lexer.NewLexer(c)
	p := NewAlgorithm(l)
	pr := p.parserVarDeclBlock()
	assert.Equal(t, true, pr.Parsed)
}
