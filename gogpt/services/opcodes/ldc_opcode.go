package opcodes

import (
	"github.com/alexgarzao/gogpt-interpreter/gogpt/model/constant_pool"
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
func (d *LDCOpcode) GetOperandCount() int {
	return d.OperandCount
}

// FetchOperands gets the opcode operands.
func (d *LDCOpcode) FetchOperands(op int) error {
	d.CpIndex = op
	return nil
}

// Execute receives the context and runs the opcode.
func (d *LDCOpcode) Execute(cp *constant_pool.CP, vars *vars.Vars, st *stack.Stack, stdin StdinInterface, stdout StdoutInterface) error {
	cpv, err := cp.Get(d.CpIndex)
	if err != nil {
		return err
	}

	st.Push(stack.StackItem(cpv))

	return nil
}
