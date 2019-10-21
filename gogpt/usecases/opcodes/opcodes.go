package opcodes

import (
	"github.com/alexgarzao/gogpt-interpreter/gogpt/entities/constant_pool"
	"github.com/alexgarzao/gogpt-interpreter/gogpt/entities/stack"
	"github.com/alexgarzao/gogpt-interpreter/gogpt/entities/vars"
)

const (
	Nop int = iota
	Ldc
	Add
	Call
	Stv
	Ldv
)

type Instruction struct {
	Name         string
	Opcode       int
	OperandCount int
}

type InstructionImplementation interface {
	GetOperandCount() int
	FetchOperands(op int) error
	Execute(cp *constant_pool.CP, vars *vars.Vars, st *stack.Stack, stdout StdoutInterface) error
}
