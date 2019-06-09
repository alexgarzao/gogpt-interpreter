package opcodes

import (
	"fmt"
	"log"

	"github.com/alexgarzao/gpt-interpreter/gpt/entities/constant_pool"
	"github.com/alexgarzao/gpt-interpreter/gpt/entities/stack"
)

type AddOpcode struct {
	Instruction
}

func NewAddOpcode() *AddOpcode {
	return &AddOpcode{Instruction{"ADD", Add, 0}}
}

func (i *AddOpcode) GetOperandCount() int {
	return i.OperandCount
}

func (i *AddOpcode) FetchOperands(op int) error {
	return nil
}

func (i *AddOpcode) Execute(cp *constant_pool.CP, st *stack.Stack) error {
	op2, err := st.Pop()
	if err != nil {
		return err
	}
	op1, err := st.Pop()
	if err != nil {
		return err
	}
	if fmt.Sprintf("%T", op1) != fmt.Sprintf("%T", op2) {
		log.Fatalln("Invalid types in ADD opcode: ")
	}

	var res interface{}
	switch op1.(type) {
	case int:
		res = op1.(int) + op2.(int)
	case string:
		res = op1.(string) + op2.(string)
	}

	st.Push(stack.StackItem(res))

	return nil
}
