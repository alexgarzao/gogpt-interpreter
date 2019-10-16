package syntax

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/alexgarzao/gogpt-interpreter/gogpt/entities/lexical_analyzer"
)

func TestValidEmptyMainBlock(t *testing.T) {
	l := lexer.NewLexer(`início fim`)
	s := NewAlgorithm() 
	assert.Equal(t, true, s.ParserStmBlock(l))
}

func TestInvalidEmptyMainBlock(t *testing.T) {
	l := lexer.NewLexer(`início fimm`)
	s := NewAlgorithm()
	assert.Equal(t, false, s.ParserStmBlock(l))
}

func TestValidMainBlockWithOneSentence(t *testing.T) {
	l := lexer.NewLexer(`início imprima("hello"); fim`)
	s := NewAlgorithm()
	assert.Equal(t, true, s.ParserStmBlock(l))
}

func TestValidMainBlockWithNSentences(t *testing.T) {
	l := lexer.NewLexer(`início imprima("hello"); imprima("hello again!"); myfunction(); fim`)
	s := NewAlgorithm()
	assert.Equal(t, true, s.ParserStmBlock(l))
}

func TestInvalidMainBlockWithOneSentence(t *testing.T) {
	l := lexer.NewLexer(`início imprima("hello") fim`)
	s := NewAlgorithm()
	assert.Equal(t, false, s.ParserStmBlock(l))
}

func TestInvalidMainBlockWithNSentences(t *testing.T) {
	l := lexer.NewLexer(`início xxx("aa") imprima("hello"); fim`)
	s := NewAlgorithm()
	assert.Equal(t, false, s.ParserStmBlock(l))
}
