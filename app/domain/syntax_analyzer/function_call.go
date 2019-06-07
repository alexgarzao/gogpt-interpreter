package analyzer

import (
	lexer "github.com/alexgarzao/gpt-interpreter/app/domain/lexical_analyzer"
)

type FunctionCall struct {
}

func NewFunctionCall() *FunctionCall {
	return &FunctionCall{}
}

func (fc *FunctionCall) TryToParse(l *lexer.Lexer) bool {
	l.SaveBacktrackingPoint()
	if fc.isValid(l) {
		return true
	}

	l.BackTracking()
	return false
}

func (fc *FunctionCall) isValid(l *lexer.Lexer) bool {
	if l.GetNextTokenIf(lexer.IDENT) == nil {
		return false
	}

	if l.GetNextTokenIf(lexer.LPAREN) == nil {
		return false
	}

	if l.GetNextTokenIf(lexer.STRING) != nil {
		for {
			if l.GetNextTokenIf(lexer.COMMA) == nil {
				break
			}
			if l.GetNextTokenIf(lexer.STRING) == nil {
				return false
			}
		}
	}

	if l.GetNextTokenIf(lexer.RPAREN) == nil {
		return false
	}

	return true
}
