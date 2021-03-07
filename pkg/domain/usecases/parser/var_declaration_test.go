package parser

import (
	"testing"

	"github.com/alexgarzao/gogpt-interpreter/pkg/domain/usecases/lexer"
	"github.com/stretchr/testify/assert"
)

func TestValidVarDeclaration(t *testing.T) {
	alg := `		variáveis
	nome: literal;
fim-variáveis`
	l := lexer.New(alg)
	p := New(l)
	err := p.parserVarDeclBlock()
	assert.Nil(t, err)
}
