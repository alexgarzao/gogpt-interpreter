package parser

import (
	"testing"

	"github.com/alexgarzao/gogpt-interpreter/pkg/domain/usecases/lexer"
	"github.com/stretchr/testify/assert"
)

func TestValidEmptyMainBlock(t *testing.T) {
	l := lexer.New(`início fim`)
	p := New(l)
	err := p.parserStmBlock()
	assert.Nil(t, err)
}

func TestInvalidEmptyMainBlock(t *testing.T) {
	l := lexer.New(`início fimm`)
	p := New(l)
	err := p.parserStmBlock()
	assert.EqualError(t, err, "Expected FIM")
}

func TestValidMainBlockWithOneSentence(t *testing.T) {
	l := lexer.New(`início imprima("hello"); fim`)
	p := New(l)
	err := p.parserStmBlock()
	assert.Nil(t, err)
}

func TestValidMainBlockWithNSentences(t *testing.T) {
	l := lexer.New(`início imprima("hello"); imprima("hello again!"); myfunction(); fim`)
	p := New(l)
	err := p.parserStmBlock()
	assert.Nil(t, err)
}

func TestInvalidMainBlockWithOneSentence(t *testing.T) {
	l := lexer.New(`início imprima("hello") fim`)
	p := New(l)
	err := p.parserStmBlock()
	assert.EqualError(t, err, "Expected SEMICOLON")
}

func TestInvalidMainBlockWithNSentences(t *testing.T) {
	l := lexer.New(`início xxx("aa") imprima("hello"); fim`)
	p := New(l)
	err := p.parserStmBlock()
	assert.EqualError(t, err, "Expected SEMICOLON")
}
