package parser

import (
	"testing"

	"github.com/alexgarzao/gogpt-interpreter/gogpt/services/lexer"
	"github.com/stretchr/testify/assert"
)

func TestFunctionCallWithReturn(t *testing.T) {
	l := lexer.New(`nome := leia()`)
	p := New(l)
	pr := p.parserStmAttr()
	assert.Equal(t, true, pr.Parsed)
}
