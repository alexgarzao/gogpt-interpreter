package opcodes

import (
	"github.com/alexgarzao/gogpt-interpreter/gogpt/model/constant_pool"
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
func (d *NOPOpcode) GetOperandCount() int {
	return d.OperandCount
}

// FetchOperands gets the opcode operands.
func (d *NOPOpcode) FetchOperands(op int) error {
	return nil
}

// Execute receives the context and runs the opcode.
func (d *NOPOpcode) Execute(cp *constant_pool.CP, vars *vars.Vars, st *stack.Stack, stdin StdinInterface, stdout StdoutInterface) error {
	return nil
}
