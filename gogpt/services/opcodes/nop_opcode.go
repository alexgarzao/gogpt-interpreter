package opcodes

import (
	"github.com/alexgarzao/gogpt-interpreter/gogpt/model/constant_pool"
	"github.com/alexgarzao/gogpt-interpreter/gogpt/model/stack"
	"github.com/alexgarzao/gogpt-interpreter/gogpt/model/vars"
)

// NopOpcode is responsible for nothing.
type NopOpcode struct {
	Instruction
}

// NewNopOpcode creates a new NopOpcode.
func NewNopOpcode() *NopOpcode {
	return &NopOpcode{Instruction{"NOP", Nop, 0}}
}

// GetOperandCount gets the numbers os opcode operands.
func (d *NopOpcode) GetOperandCount() int {
	return d.OperandCount
}

// FetchOperands gets the opcode operands.
func (d *NopOpcode) FetchOperands(op int) error {
	return nil
}

// Execute receives the context and runs the opcode.
func (d *NopOpcode) Execute(cp *constant_pool.CP, vars *vars.Vars, st *stack.Stack, stdin StdinInterface, stdout StdoutInterface) error {
	return nil
}
