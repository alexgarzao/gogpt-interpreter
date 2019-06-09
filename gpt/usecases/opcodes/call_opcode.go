package opcodes

import (
	cp "github.com/alexgarzao/gpt-interpreter/gpt/entities/constant_pool"
	"github.com/alexgarzao/gpt-interpreter/gpt/entities/stack"
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

func (i *CallOpcode) Execute(cp *cp.CP, stack *stack.Stack, stdout StdoutInterface) error {
	cpv, err := cp.Get(i.CpIndex)
	if err != nil {
		return err
	}

	if cpv == "io.println" {
		stv, err := stack.Pop()
		if err != nil {
			return err
		}

		stdout.Println(stv)
	}

	return nil
}
