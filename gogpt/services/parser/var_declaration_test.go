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
	l := lexer.New(c)
	p := New(l)
	pr := p.parserVarDeclBlock()
	assert.Equal(t, true, pr.Parsed)
}
