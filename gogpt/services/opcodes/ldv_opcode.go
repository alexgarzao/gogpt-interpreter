package opcodes

import (
	"github.com/alexgarzao/gogpt-interpreter/gogpt/model/cp"
	"github.com/alexgarzao/gogpt-interpreter/gogpt/model/stack"
	"github.com/alexgarzao/gogpt-interpreter/gogpt/model/vars"
)

// LDVOpcode is responsible for push a variable content into the stack.
type LDVOpcode struct {
	Instruction
	VarIndex int
}

// NewLDVOpcode creates a new LDVOpcode.
func NewLDVOpcode() *LDVOpcode {
	return &LDVOpcode{Instruction{"LDV", LDV, 1}, 0}
}

// GetOperandCount gets the numbers os opcode operands.
func (i *LDVOpcode) GetOperandCount() int {
	return i.OperandCount
}

// FetchOperands gets the opcode operands.
func (i *LDVOpcode) FetchOperands(op int) error {
	i.VarIndex = op
	return nil
}

// Execute receives the context and runs the opcode.
func (i *LDVOpcode) Execute(cp *cp.CP, vars *vars.Vars, st *stack.Stack, stdin StdinInterface, stdout StdoutInterface) error {
	value, err := vars.Get(i.VarIndex)
	if err != nil {
		return err
	}

	st.Push(value)

	return nil
}
