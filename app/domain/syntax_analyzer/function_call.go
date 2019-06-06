package analyzer

import (
	lexer "github.com/alexgarzao/gpt-interpreter/app/domain/lexical_analyzer"
)

type FunctionCall struct {
}

func NewFunctionCall() *FunctionCall {
	return &FunctionCall{}
}

func (fc *FunctionCall) IsValid(l *lexer.Lexer) bool {
	if l.NextToken().Type != lexer.IDENT {
		l.BackTracking()
		return false
	}

	if l.NextToken().Type != lexer.LPAREN {
		l.BackTracking()
		return false
	}

	if l.NextToken().Type != lexer.STRING {
		l.BackTracking()
	} else {
		for {
			if l.NextToken().Type != lexer.COMMA {
				l.BackTracking()
				break
			}
			if l.NextToken().Type != lexer.STRING {
				l.BackTracking()
				return false
			}
		}
	}

	if l.NextToken().Type != lexer.RPAREN {
		return false
	}

	return true
}
