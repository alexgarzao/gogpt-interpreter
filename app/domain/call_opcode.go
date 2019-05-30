package opcodes

import (
	"fmt"
)

type CallOpcode struct {
	Instruction
	CpIndex BytecodeItem
}

func NewCallOpcode() *CallOpcode {
	return &CallOpcode{Instruction{"CALL", Call, 1}, 0}
}

func (i *CallOpcode) FetchOperands(bc *Bytecode) {
	i.CpIndex, _ = bc.Next()
}

func (i *CallOpcode) Execute(cp *CP, stack *Stack) {
	cpv, _ := cp.Get(int(i.CpIndex))
	if cpv == "io.println" {
		stv, _ := stack.Pop()
		fmt.Printf("%v\n", stv)
	}
}
