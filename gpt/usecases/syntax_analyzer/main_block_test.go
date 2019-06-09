package syntax

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/alexgarzao/gpt-interpreter/gpt/entities/lexical_analyzer"
)

func TestValidEmptyMainBlock(t *testing.T) {
	l := lexer.NewLexer(`início fim`)
	s := NewMainBlock()
	assert.Equal(t, true, s.TryToParse(l))
}

func TestInvalidEmptyMainBlock(t *testing.T) {
	l := lexer.NewLexer(`início fimm`)
	s := NewMainBlock()
	assert.Equal(t, false, s.TryToParse(l))
}

func TestValidMainBlockWithOneSentence(t *testing.T) {
	l := lexer.NewLexer(`início imprima("hello"); fim`)
	s := NewMainBlock()
	assert.Equal(t, true, s.TryToParse(l))
}

func TestValidMainBlockWithNSentences(t *testing.T) {
	l := lexer.NewLexer(`início imprima("hello"); imprima("hello again!"); myfunction(); fim`)
	s := NewMainBlock()
	assert.Equal(t, true, s.TryToParse(l))
}

func TestInvalidMainBlockWithOneSentence(t *testing.T) {
	l := lexer.NewLexer(`início imprima("hello") fim`)
	s := NewMainBlock()
	assert.Equal(t, false, s.TryToParse(l))
}

func TestInvalidMainBlockWithNSentences(t *testing.T) {
	l := lexer.NewLexer(`início xxx("aa") imprima("hello"); fim`)
	s := NewMainBlock()
	assert.Equal(t, false, s.TryToParse(l))
}
