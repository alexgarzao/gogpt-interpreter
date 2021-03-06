package addi

import (
	"fmt"
	"log"

	"github.com/alexgarzao/gogpt-interpreter/pkg/model/cp"
	"github.com/alexgarzao/gogpt-interpreter/pkg/model/stack"
	"github.com/alexgarzao/gogpt-interpreter/pkg/services/instructions"
)

// ADDInst is an opcode responsible for get two elements from the stack, add, and push onto the stack again.
type ADDInst struct {
	instructions.Instruction
}

// New creates a new ADDInst.
func New() *ADDInst {
	return &ADDInst{
		instructions.Instruction{
			Opcode: instructions.ADD,
		},
	}
}

// FetchOperands gets the opcode operands.
func (i *ADDInst) FetchOperands(fetch instructions.FetchOperandsImplementation) error {
	return nil
}

// Execute receives the context and runs the opcode.
func (i *ADDInst) Execute(cp *cp.CP, st *stack.Stack) error {
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

	st.Push(res)

	return nil
}
