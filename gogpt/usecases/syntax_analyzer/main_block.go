package syntax

import (
	"github.com/alexgarzao/gogpt-interpreter/gogpt/entities/bytecode"
	"github.com/alexgarzao/gogpt-interpreter/gogpt/entities/constant_pool"
	"github.com/alexgarzao/gogpt-interpreter/gogpt/entities/lexical_analyzer"
)

type MainBlock struct {
	cp *constant_pool.CP
	bc *bytecode.Bytecode
}

func NewMainBlock() *MainBlock {
	return &MainBlock{}
}

func (mb *MainBlock) SetBytecodeGenRequirements(cp *constant_pool.CP, bc *bytecode.Bytecode) *MainBlock {
	mb.cp = cp
	mb.bc = bc

	return mb
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

	fc := NewFunctionCall().
		SetBytecodeGenRequirements(mb.cp, mb.bc)

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
