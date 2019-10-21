package syntax

import (
	"testing"

	lexer "github.com/alexgarzao/gogpt-interpreter/gogpt/entities/lexical_analyzer"
	"github.com/stretchr/testify/assert"
)

func TestFunctionCallWithReturn(t *testing.T) {
	l := lexer.NewLexer(`nome = leia()`)
	s := NewAlgorithm(l)
	pr := s.ParserStmAttr()
	assert.Equal(t, true, pr.Parsed)
}
