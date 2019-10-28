package opcodes

import (
	"github.com/alexgarzao/gogpt-interpreter/gogpt/model/constant_pool"
	"github.com/alexgarzao/gogpt-interpreter/gogpt/model/stack"
	"github.com/alexgarzao/gogpt-interpreter/gogpt/model/vars"
)

// LdcOpcode is responsible for push a constant into the stack.
type LdcOpcode struct {
	Instruction
	CpIndex int
}

// NewLdcOpcode creates a new LdcOpcode.
func NewLdcOpcode() *LdcOpcode {
	return &LdcOpcode{Instruction{"LDC", Ldc, 1}, 0}
}

// GetOperandCount gets the numbers os opcode operands.
func (d *LdcOpcode) GetOperandCount() int {
	return d.OperandCount
}

// FetchOperands gets the opcode operands.
func (d *LdcOpcode) FetchOperands(op int) error {
	d.CpIndex = op
	return nil
}

// Execute receives the context and runs the opcode.
func (d *LdcOpcode) Execute(cp *constant_pool.CP, vars *vars.Vars, st *stack.Stack, stdin StdinInterface, stdout StdoutInterface) error {
	cpv, err := cp.Get(d.CpIndex)
	if err != nil {
		return err
	}

	st.Push(stack.StackItem(cpv))

	return nil
}
