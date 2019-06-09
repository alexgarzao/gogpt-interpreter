package opcodes

import (
	"github.com/alexgarzao/gpt-interpreter/gpt/entities/bytecode"
	"github.com/alexgarzao/gpt-interpreter/gpt/entities/constant_pool"
	"github.com/alexgarzao/gpt-interpreter/gpt/entities/stack"
)

type LdcOpcode struct {
	Instruction
	CpIndex int
}

func NewLdcOpcode() *LdcOpcode {
	return &LdcOpcode{Instruction{"LDC", Ldc, 1}, 0}
}

func (d *LdcOpcode) FetchOperands(bc *bytecode.Bytecode) error {
	var err error
	d.CpIndex, err = bc.Next()
	return err
}

func (d *LdcOpcode) Execute(cp *constant_pool.CP, st *stack.Stack, stdout StdoutInterface) error {
	cpv, err := cp.Get(d.CpIndex)
	if err != nil {
		return err
	}

	st.Push(stack.StackItem(cpv))

	return nil
}
