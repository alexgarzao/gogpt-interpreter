package stvi

import (
	"github.com/alexgarzao/gogpt-interpreter/gogpt/model/cp"
	"github.com/alexgarzao/gogpt-interpreter/gogpt/model/stack"
	"github.com/alexgarzao/gogpt-interpreter/gogpt/model/vars"
	"github.com/alexgarzao/gogpt-interpreter/gogpt/services/instructions"
)

// STVInst is responsible for pop a value from the stack and put into a var.
type STVInst struct {
	instructions.Instruction
	VarIndex int
}

// New creates a new STVInst.
func New() *STVInst {
	return &STVInst{instructions.Instruction{"STV", instructions.STV, 1}, 0}
}

// GetOperandCount gets the numbers os opcode operands.
func (i *STVInst) GetOperandCount() int {
	return i.OperandCount
}

// FetchOperands gets the opcode operands.
func (i *STVInst) FetchOperands(op int) error {
	i.VarIndex = op
	return nil
}

// Execute receives the context and runs the opcode.
func (i *STVInst) Execute(cp *cp.CP, vars *vars.Vars, st *stack.Stack, stdin instructions.StdinInterface, stdout instructions.StdoutInterface) error {
	value, err := st.Pop()
	if err != nil {
		return err
	}

	vars.Set(i.VarIndex, value)

	return nil
}
