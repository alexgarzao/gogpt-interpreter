package opcodes

import (
	"github.com/alexgarzao/gogpt-interpreter/gogpt/model/constant_pool"
	"github.com/alexgarzao/gogpt-interpreter/gogpt/model/stack"
	"github.com/alexgarzao/gogpt-interpreter/gogpt/model/vars"
)

// STVOpcode is responsible for pop a value from the stack and put into a var.
type STVOpcode struct {
	Instruction
	VarIndex int
}

// NewSTVOpcode creates a new STVOpcode.
func NewSTVOpcode() *STVOpcode {
	return &STVOpcode{Instruction{"STV", STV, 1}, 0}
}

// GetOperandCount gets the numbers os opcode operands.
func (d *STVOpcode) GetOperandCount() int {
	return d.OperandCount
}

// FetchOperands gets the opcode operands.
func (d *STVOpcode) FetchOperands(op int) error {
	d.VarIndex = op
	return nil
}

// Execute receives the context and runs the opcode.
func (d *STVOpcode) Execute(cp *constant_pool.CP, vars *vars.Vars, st *stack.Stack, stdin StdinInterface, stdout StdoutInterface) error {
	value, err := st.Pop()
	if err != nil {
		return err
	}

	vars.Set(d.VarIndex, value)

	return nil
}
