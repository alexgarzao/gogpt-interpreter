package opcodes

import (
	"fmt"
	"log"
)

type AddOpcode struct {
	Instruction
}

func NewAddOpcode() *AddOpcode {
	return &AddOpcode{Instruction{"ADD", Add, 0}}
}

func (i *AddOpcode) FetchOperands(bc *Bytecode) {
}

func (i *AddOpcode) Execute(cp *CP, stack *Stack) {
	op2, _ := stack.Pop()
	op1, _ := stack.Pop()
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

	stack.Push(StackItem(res))
}
