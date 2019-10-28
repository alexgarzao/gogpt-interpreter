package opcodes

import (
	"github.com/alexgarzao/gogpt-interpreter/gogpt/model/constant_pool"
	"github.com/alexgarzao/gogpt-interpreter/gogpt/model/stack"
	"github.com/alexgarzao/gogpt-interpreter/gogpt/model/vars"
)

// StvOpcode is responsible for pop a value from the stack and put into a var.
type StvOpcode struct {
	Instruction
	VarIndex int
}

// NewStvOpcode creates a new StvOpcode.
func NewStvOpcode() *StvOpcode {
	return &StvOpcode{Instruction{"STV", Stv, 1}, 0}
}

// GetOperandCount gets the numbers os opcode operands.
func (d *StvOpcode) GetOperandCount() int {
	return d.OperandCount
}

// FetchOperands gets the opcode operands.
func (d *StvOpcode) FetchOperands(op int) error {
	d.VarIndex = op
	return nil
}

// Execute receives the context and runs the opcode.
func (d *StvOpcode) Execute(cp *constant_pool.CP, vars *vars.Vars, st *stack.Stack, stdin StdinInterface, stdout StdoutInterface) error {
	value, err := st.Pop()
	if err != nil {
		return err
	}

	vars.Set(d.VarIndex, value)

	return nil
}
