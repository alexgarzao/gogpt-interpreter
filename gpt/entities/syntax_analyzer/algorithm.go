package analyzer

import (
	opcodes "github.com/alexgarzao/gpt-interpreter/gpt/entities"
	lexer "github.com/alexgarzao/gpt-interpreter/gpt/entities/lexical_analyzer"
)

type Program struct {
	cp *opcodes.CP
	bc *opcodes.Bytecode
}

func NewProgram() *Program {
	return &Program{
		cp: opcodes.NewCp(),
		bc: opcodes.NewBytecode(),
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

func (p *Program) GetCP() *opcodes.CP {
	return p.cp
}

func (p *Program) GetBC() *opcodes.Bytecode {
	return p.bc
}
