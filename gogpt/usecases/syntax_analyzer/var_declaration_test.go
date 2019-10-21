package syntax

import (
	"testing"

	"github.com/stretchr/testify/assert"

	lexer "github.com/alexgarzao/gogpt-interpreter/gogpt/entities/lexical_analyzer"
)

func TestValidVarDeclaration(t *testing.T) {
	c := `		variáveis
	nome: literal;
fim-variáveis`
	l := lexer.NewLexer(c)
	p := NewAlgorithm(l)
	pr := p.ParserVarDeclBlock()
	assert.Equal(t, true, pr.Parsed)
}
