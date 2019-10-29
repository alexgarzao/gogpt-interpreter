package opcodes

import (
	"fmt"
	"log"

	"github.com/alexgarzao/gogpt-interpreter/gogpt/model/cp"
	"github.com/alexgarzao/gogpt-interpreter/gogpt/model/stack"
)

// ADDOpcode is an opcode responsible for get two elements from the stack, add, and push onto the stack again.
type ADDOpcode struct {
	Instruction
}

// NewADDOpcode creates a new ADDOpcode.
func NewADDOpcode() *ADDOpcode {
	return &ADDOpcode{Instruction{"ADD", ADD, 0}}
}

// GetOperandCount gets the numbers os opcode operands.
func (i *ADDOpcode) GetOperandCount() int {
	return i.OperandCount
}

// FetchOperands gets the opcode operands.
func (i *ADDOpcode) FetchOperands(op int) error {
	return nil
}

// Execute receives the context and runs the opcode.
func (i *ADDOpcode) Execute(cp *cp.CP, st *stack.Stack) error {
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
