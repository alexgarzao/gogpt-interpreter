package parser

import (
	"testing"

	"github.com/alexgarzao/gogpt-interpreter/pkg/domain/usecases/lexer"
	"github.com/stretchr/testify/assert"
)

func TestFunctionCallWithReturn(t *testing.T) {
	l := lexer.New(`nome := leia()`)
	p := New(l)
	err := p.parserStmAttr()
	assert.Nil(t, err)
}
