package opcodes

import (
	"github.com/alexgarzao/gogpt-interpreter/gogpt/model/vars"
	"github.com/alexgarzao/gogpt-interpreter/gogpt/model/constant_pool"
	"github.com/alexgarzao/gogpt-interpreter/gogpt/model/stack"
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
	Execute(cp *constant_pool.CP, vars *vars.Vars, st *stack.Stack, stdin StdinInterface, stdout StdoutInterface) error
}
