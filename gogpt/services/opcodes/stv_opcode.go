package opcodes

import (
	"github.com/alexgarzao/gogpt-interpreter/gogpt/model/cp"
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
func (i *STVOpcode) GetOperandCount() int {
	return i.OperandCount
}

// FetchOperands gets the opcode operands.
func (i *STVOpcode) FetchOperands(op int) error {
	i.VarIndex = op
	return nil
}

// Execute receives the context and runs the opcode.
func (i *STVOpcode) Execute(cp *cp.CP, vars *vars.Vars, st *stack.Stack, stdin StdinInterface, stdout StdoutInterface) error {
	value, err := st.Pop()
	if err != nil {
		return err
	}

	vars.Set(i.VarIndex, value)

	return nil
}
