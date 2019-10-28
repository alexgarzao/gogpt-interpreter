package opcodes

import (
	"github.com/alexgarzao/gogpt-interpreter/gogpt/model/vars"
	"github.com/alexgarzao/gogpt-interpreter/gogpt/model/constant_pool"
	"github.com/alexgarzao/gogpt-interpreter/gogpt/model/stack"
)

type StvOpcode struct {
	Instruction
	VarIndex int
}

func NewStvOpcode() *StvOpcode {
	return &StvOpcode{Instruction{"STV", Stv, 1}, 0}
}

func (d *StvOpcode) GetOperandCount() int {
	return d.OperandCount
}

func (d *StvOpcode) FetchOperands(op int) error {
	d.VarIndex = op
	return nil
}

func (d *StvOpcode) Execute(cp *constant_pool.CP, vars *vars.Vars, st *stack.Stack, stdin StdinInterface, stdout StdoutInterface) error {
	value, err := st.Pop()
	if err != nil {
		return err
	}

	vars.Set(d.VarIndex, value)

	return nil
}
