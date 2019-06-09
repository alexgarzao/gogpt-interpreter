package opcodes

import (
	"github.com/alexgarzao/gpt-interpreter/gpt/entities/constant_pool"
	"github.com/alexgarzao/gpt-interpreter/gpt/entities/stack"
)

const (
	Nop int = iota
	Ldc
	Add
	Call
)

type Instruction struct {
	Name         string
	Opcode       int
	OperandCount int
}

type InstructionImplementation interface {
	GetOperandCount() int
	FetchOperands(op int) error
	Execute(cp *constant_pool.CP, st *stack.Stack, stdout StdoutInterface) error
}
