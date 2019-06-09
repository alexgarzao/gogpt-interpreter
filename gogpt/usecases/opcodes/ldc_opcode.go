package opcodes

import (
	"github.com/alexgarzao/gogpt-interpreter/gogpt/entities/constant_pool"
	"github.com/alexgarzao/gogpt-interpreter/gogpt/entities/stack"
)

type LdcOpcode struct {
	Instruction
	CpIndex int
}

func NewLdcOpcode() *LdcOpcode {
	return &LdcOpcode{Instruction{"LDC", Ldc, 1}, 0}
}

func (d *LdcOpcode) GetOperandCount() int {
	return d.OperandCount
}

func (d *LdcOpcode) FetchOperands(op int) error {
	d.CpIndex = op
	return nil
}

func (d *LdcOpcode) Execute(cp *constant_pool.CP, st *stack.Stack, stdout StdoutInterface) error {
	cpv, err := cp.Get(d.CpIndex)
	if err != nil {
		return err
	}

	st.Push(stack.StackItem(cpv))

	return nil
}
