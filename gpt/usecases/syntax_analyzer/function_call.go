package syntax

import (
	"github.com/alexgarzao/gpt-interpreter/gpt/entities/bytecode"
	"github.com/alexgarzao/gpt-interpreter/gpt/entities/constant_pool"
	"github.com/alexgarzao/gpt-interpreter/gpt/entities/lexical_analyzer"
	"github.com/alexgarzao/gpt-interpreter/gpt/usecases/opcodes"
)

type FunctionCall struct {
	cp *constant_pool.CP
	bc *bytecode.Bytecode
}

func NewFunctionCall() *FunctionCall {
	return &FunctionCall{}
}

func (fc *FunctionCall) SetBytecodeGenRequirements(cp *constant_pool.CP, bc *bytecode.Bytecode) *FunctionCall {
	fc.cp = cp
	fc.bc = bc

	return fc
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
	var token *lexer.Token
	if token = l.GetNextTokenIf(lexer.IDENT); token == nil {
		return false
	}

	if l.GetNextTokenIf(lexer.LPAREN) == nil {
		return false
	}

	funcIndex := -1
	if token.Value == "imprima" {
		funcIndex = fc.cp.Add("io.println")
	}

	if token = l.GetNextTokenIf(lexer.STRING); token != nil {
		for {
			cpIndex := fc.cp.Add(token.Value)
			fc.bc.Add(opcodes.Ldc, cpIndex)
			if l.GetNextTokenIf(lexer.COMMA) == nil {
				break
			}
			if l.GetNextTokenIf(lexer.STRING) == nil {
				return false
			}
		}
	}

	fc.bc.Add(opcodes.Call, funcIndex)

	if l.GetNextTokenIf(lexer.RPAREN) == nil {
		return false
	}

	return true
}
