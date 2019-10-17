package syntax

import (
	"testing"

	lexer "github.com/alexgarzao/gogpt-interpreter/gogpt/entities/lexical_analyzer"
	"github.com/stretchr/testify/assert"
)

func TestValidEmptyMainBlock(t *testing.T) {
	l := lexer.NewLexer(`início fim`)
	s := NewAlgorithm(l)
	assert.Equal(t, true, s.ParserStmBlock())
}

func TestInvalidEmptyMainBlock(t *testing.T) {
	l := lexer.NewLexer(`início fimm`)
	s := NewAlgorithm(l)
	assert.Equal(t, false, s.ParserStmBlock())
}

func TestValidMainBlockWithOneSentence(t *testing.T) {
	l := lexer.NewLexer(`início imprima("hello"); fim`)
	s := NewAlgorithm(l)
	assert.Equal(t, true, s.ParserStmBlock())
}

func TestValidMainBlockWithNSentences(t *testing.T) {
	l := lexer.NewLexer(`início imprima("hello"); imprima("hello again!"); myfunction(); fim`)
	s := NewAlgorithm(l)
	assert.Equal(t, true, s.ParserStmBlock())
}

func TestInvalidMainBlockWithOneSentence(t *testing.T) {
	l := lexer.NewLexer(`início imprima("hello") fim`)
	s := NewAlgorithm(l)
	assert.Equal(t, false, s.ParserStmBlock())
}

func TestInvalidMainBlockWithNSentences(t *testing.T) {
	l := lexer.NewLexer(`início xxx("aa") imprima("hello"); fim`)
	s := NewAlgorithm(l)
	assert.Equal(t, false, s.ParserStmBlock())
}
