package analyzer

import (
	lexer "github.com/alexgarzao/gpt-interpreter/app/domain/lexical_analyzer"
)

type MainBlock struct {
}

func NewMainBlock() *MainBlock {
	return &MainBlock{}
}

func (mb *MainBlock) TryToParse(l *lexer.Lexer) bool {
	l.SaveBacktrackingPoint()
	if mb.isValid(l) {
		return true
	}

	l.BackTracking()
	return false
}

func (mb *MainBlock) isValid(l *lexer.Lexer) bool {
	if l.GetNextTokenIf(lexer.INICIO) == nil {
		return false
	}

	fc := NewFunctionCall()

	for fc.TryToParse(l) {
		if l.GetNextTokenIf(lexer.SEMICOLON) == nil {
			return false
		}
	}

	if l.GetNextTokenIf(lexer.FIM) == nil {
		return false
	}

	return true
}
