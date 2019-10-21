package opcodes

import (
	"github.com/alexgarzao/gogpt-interpreter/gogpt/entities/constant_pool"
	"github.com/alexgarzao/gogpt-interpreter/gogpt/entities/stack"
	"github.com/alexgarzao/gogpt-interpreter/gogpt/entities/vars"
)

type LdvOpcode struct {
	Instruction
	VarIndex int
}

func NewLdvOpcode() *LdvOpcode {
	return &LdvOpcode{Instruction{"LDV", Ldv, 1}, 0}
}

func (d *LdvOpcode) GetOperandCount() int {
	return d.OperandCount
}

func (d *LdvOpcode) FetchOperands(op int) error {
	d.VarIndex = op
	return nil
}

func (d *LdvOpcode) Execute(cp *constant_pool.CP, vars *vars.Vars, st *stack.Stack, stdout StdoutInterface) error {
	value, err := vars.Get(d.VarIndex)
	if err != nil {
		return err
	}

	st.Push(value)

	return nil
}
