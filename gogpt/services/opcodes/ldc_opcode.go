package opcodes

import (
	"github.com/alexgarzao/gogpt-interpreter/gogpt/model/cp"
	"github.com/alexgarzao/gogpt-interpreter/gogpt/model/stack"
	"github.com/alexgarzao/gogpt-interpreter/gogpt/model/vars"
)

// LDCOpcode is responsible for push a constant into the stack.
type LDCOpcode struct {
	Instruction
	CpIndex int
}

// NewLDCOpcode creates a new LDCOpcode.
func NewLDCOpcode() *LDCOpcode {
	return &LDCOpcode{Instruction{"LDC", LDC, 1}, 0}
}

// GetOperandCount gets the numbers os opcode operands.
func (i *LDCOpcode) GetOperandCount() int {
	return i.OperandCount
}

// FetchOperands gets the opcode operands.
func (i *LDCOpcode) FetchOperands(op int) error {
	i.CpIndex = op
	return nil
}

// Execute receives the context and runs the opcode.
func (i *LDCOpcode) Execute(cp *cp.CP, vars *vars.Vars, st *stack.Stack, stdin StdinInterface, stdout StdoutInterface) error {
	cpv, err := cp.Get(i.CpIndex)
	if err != nil {
		return err
	}

	st.Push(cpv)

	return nil
}
