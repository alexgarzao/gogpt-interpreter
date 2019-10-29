package instructions

import (
	"github.com/alexgarzao/gogpt-interpreter/gogpt/model/cp"
	"github.com/alexgarzao/gogpt-interpreter/gogpt/model/stack"
	"github.com/alexgarzao/gogpt-interpreter/gogpt/model/vars"
)

const (
	// NOP opcode
	NOP int = iota
	// LDC opcode
	LDC
	// ADD opcode
	ADD
	// CALL opcode
	CALL
	// STV opcode
	STV
	// LDV opcode
	LDV
)

// Instruction has the common data of all instructions.
type Instruction struct {
	Name         string
	Opcode       int
	OperandCount int
}

// InstructionImplementation has the minimal interface to be a valid instruction.
type InstructionImplementation interface {
	GetOperandCount() int
	FetchOperands(op int) error
	Execute(cp *cp.CP, vars *vars.Vars, st *stack.Stack, stdin StdinInterface, stdout StdoutInterface) error
}
