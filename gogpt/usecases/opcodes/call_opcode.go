package opcodes

import (
	"github.com/alexgarzao/gogpt-interpreter/gogpt/entities/constant_pool"
	"github.com/alexgarzao/gogpt-interpreter/gogpt/entities/stack"
	"github.com/alexgarzao/gogpt-interpreter/gogpt/entities/vars"
)

type CallOpcode struct {
	Instruction
	CpIndex int
}

func NewCallOpcode() *CallOpcode {
	return &CallOpcode{Instruction{"CALL", Call, 1}, 0}
}

func (i *CallOpcode) GetOperandCount() int {
	return i.OperandCount
}

func (i *CallOpcode) FetchOperands(op int) error {
	i.CpIndex = op
	return nil
}

func (i *CallOpcode) Execute(cp *constant_pool.CP, vars *vars.Vars, st *stack.Stack, stdout StdoutInterface) error {
	cpv, err := cp.Get(i.CpIndex)
	if err != nil {
		return err
	}

	if cpv == "io.println" {
		stv, err := st.Pop()
		if err != nil {
			return err
		}

		stdout.Println(stv)
	} else if cpv == "io.readln" {
		text := stdout.Readln()
		st.Push(text)
	}

	return nil
}
