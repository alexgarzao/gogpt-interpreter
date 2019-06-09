package syntax

import (
	"github.com/alexgarzao/gpt-interpreter/gpt/entities/bytecode"
	"github.com/alexgarzao/gpt-interpreter/gpt/entities/constant_pool"
	"github.com/alexgarzao/gpt-interpreter/gpt/entities/lexical_analyzer"
)

type Program struct {
	cp *constant_pool.CP
	bc *bytecode.Bytecode
}

func NewProgram() *Program {
	return &Program{
		cp: constant_pool.NewCp(),
		bc: bytecode.NewBytecode(),
	}
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

	mb := NewMainBlock().
		SetBytecodeGenRequirements(p.cp, p.bc)

	return mb.TryToParse(l)
}

func (p *Program) GetCP() *constant_pool.CP {
	return p.cp
}

func (p *Program) GetBC() *bytecode.Bytecode {
	return p.bc
}
