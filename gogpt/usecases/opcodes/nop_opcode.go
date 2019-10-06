package opcodes

import (
	"github.com/alexgarzao/gogpt-interpreter/gogpt/entities/constant_pool"
	"github.com/alexgarzao/gogpt-interpreter/gogpt/entities/stack"
)

type NopOpcode struct {
	Instruction
}

func NewNopOpcode() *NopOpcode {
	return &NopOpcode{Instruction{"NOP", Nop, 0}}
}

func (d *NopOpcode) GetOperandCount() int {
	return d.OperandCount
}

func (d *NopOpcode) FetchOperands(op int) error {
	return nil
}

func (d *NopOpcode) Execute(cp *constant_pool.CP, st *stack.Stack, stdout StdoutInterface) error {
	return nil
}