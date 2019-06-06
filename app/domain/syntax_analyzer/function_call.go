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
	// TODO: backtracking.
	if l.NextToken().Type != lexer.IDENT {
		return false
	}

	if l.NextToken().Type != lexer.LPAREN {
		return false
	}

	if l.NextToken().Type != lexer.STRING {
		return false
	}

	if l.NextToken().Type != lexer.RPAREN {
		return false
	}

	return true
}
