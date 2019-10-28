package opcodes

import (
	"github.com/alexgarzao/gogpt-interpreter/gogpt/model/constant_pool"
	"github.com/alexgarzao/gogpt-interpreter/gogpt/model/stack"
	"github.com/alexgarzao/gogpt-interpreter/gogpt/model/vars"
)

// LdvOpcode is responsible for push a variable content into the stack.
type LdvOpcode struct {
	Instruction
	VarIndex int
}

// NewLdvOpcode creates a new LdvOpcode.
func NewLdvOpcode() *LdvOpcode {
	return &LdvOpcode{Instruction{"LDV", Ldv, 1}, 0}
}

// GetOperandCount gets the numbers os opcode operands.
func (d *LdvOpcode) GetOperandCount() int {
	return d.OperandCount
}

// FetchOperands gets the opcode operands.
func (d *LdvOpcode) FetchOperands(op int) error {
	d.VarIndex = op
	return nil
}

// Execute receives the context and runs the opcode.
func (d *LdvOpcode) Execute(cp *constant_pool.CP, vars *vars.Vars, st *stack.Stack, stdin StdinInterface, stdout StdoutInterface) error {
	value, err := vars.Get(d.VarIndex)
	if err != nil {
		return err
	}

	st.Push(value)

	return nil
}
