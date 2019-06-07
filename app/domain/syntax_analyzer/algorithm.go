package analyzer

import (
	lexer "github.com/alexgarzao/gpt-interpreter/app/domain/lexical_analyzer"
)

type Program struct {
}

func NewProgram() *Program {
	return &Program{}
}

func (p *Program) TryToParse(l *lexer.Lexer) bool {
	if p.isValid(l) {
		return true
	}

	return false
}

func (p *Program) isValid(l *lexer.Lexer) bool {
	if l.GetNextTokenIf(lexer.ALGORITMO) == nil {
		return false
	}

	if l.GetNextTokenIf(lexer.IDENT) == nil {
		return false
	}

	if l.GetNextTokenIf(lexer.SEMICOLON) == nil {
		return false
	}

	mb := NewMainBlock()

	return mb.TryToParse(l)
}
