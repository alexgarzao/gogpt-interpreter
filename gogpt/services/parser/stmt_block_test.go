package parser

import (
	"testing"

	"github.com/alexgarzao/gogpt-interpreter/gogpt/services/lexer"
	"github.com/stretchr/testify/assert"
)

func TestValidEmptyMainBlock(t *testing.T) {
	l := lexer.New(`início fim`)
	s := New(l)
	pr := s.parserStmBlock()
	assert.Equal(t, true, pr.Parsed)
}

func TestInvalidEmptyMainBlock(t *testing.T) {
	l := lexer.New(`início fimm`)
	s := New(l)
	pr := s.parserStmBlock()
	assert.Equal(t, false, pr.Parsed)
	assert.EqualError(t, pr.Err, "Expected FIM")
}

func TestValidMainBlockWithOneSentence(t *testing.T) {
	l := lexer.New(`início imprima("hello"); fim`)
	s := New(l)
	pr := s.parserStmBlock()
	assert.Equal(t, true, pr.Parsed)
}

func TestValidMainBlockWithNSentences(t *testing.T) {
	l := lexer.New(`início imprima("hello"); imprima("hello again!"); myfunction(); fim`)
	s := New(l)
	pr := s.parserStmBlock()
	assert.Equal(t, true, pr.Parsed)
}

func TestInvalidMainBlockWithOneSentence(t *testing.T) {
	l := lexer.New(`início imprima("hello") fim`)
	s := New(l)
	pr := s.parserStmBlock()
	assert.Equal(t, false, pr.Parsed)
	assert.EqualError(t, pr.Err, "Expected SEMICOLON")
}

func TestInvalidMainBlockWithNSentences(t *testing.T) {
	l := lexer.New(`início xxx("aa") imprima("hello"); fim`)
	s := New(l)
	pr := s.parserStmBlock()
	assert.Equal(t, false, pr.Parsed)
	assert.EqualError(t, pr.Err, "Expected SEMICOLON")
}
