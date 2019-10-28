package opcodes

import (
	"fmt"
	"log"

	"github.com/alexgarzao/gogpt-interpreter/gogpt/model/constant_pool"
	"github.com/alexgarzao/gogpt-interpreter/gogpt/model/stack"
)

// AddOpcode is an opcode responsible for get two elements from the stack, add, and push onto the stack again.
type AddOpcode struct {
	Instruction
}

// NewAddOpcode creates a new AddOpcode.
func NewAddOpcode() *AddOpcode {
	return &AddOpcode{Instruction{"ADD", Add, 0}}
}

// GetOperandCount gets the numbers os opcode operands.
func (i *AddOpcode) GetOperandCount() int {
	return i.OperandCount
}

// FetchOperands gets the opcode operands.
func (i *AddOpcode) FetchOperands(op int) error {
	return nil
}

// Execute receives the context and runs the opcode.
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
