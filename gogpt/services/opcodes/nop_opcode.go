package opcodes

import (
	"github.com/alexgarzao/gogpt-interpreter/gogpt/model/cp"
	"github.com/alexgarzao/gogpt-interpreter/gogpt/model/stack"
	"github.com/alexgarzao/gogpt-interpreter/gogpt/model/vars"
)

// NOPOpcode is responsible for nothing.
type NOPOpcode struct {
	Instruction
}

// NewNOPOpcode creates a new NOPOpcode.
func NewNOPOpcode() *NOPOpcode {
	return &NOPOpcode{Instruction{"NOP", NOP, 0}}
}

// GetOperandCount gets the numbers os opcode operands.
func (i *NOPOpcode) GetOperandCount() int {
	return i.OperandCount
}

// FetchOperands gets the opcode operands.
func (i *NOPOpcode) FetchOperands(op int) error {
	return nil
}

// Execute receives the context and runs the opcode.
func (i *NOPOpcode) Execute(cp *cp.CP, vars *vars.Vars, st *stack.Stack, stdin StdinInterface, stdout StdoutInterface) error {
	return nil
}
