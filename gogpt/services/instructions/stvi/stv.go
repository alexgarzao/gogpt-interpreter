package stvi

import (
	"github.com/alexgarzao/gogpt-interpreter/gogpt/model/cp"
	"github.com/alexgarzao/gogpt-interpreter/gogpt/model/stack"
	"github.com/alexgarzao/gogpt-interpreter/gogpt/model/vars"
	"github.com/alexgarzao/gogpt-interpreter/gogpt/services/instructions"
)

// STVInst is responsible for pop a value from the stack and put into a var.
type STVInst struct {
	instructions.Instruction
	VarIndex int
}

// New creates a new STVInst.
func New() *STVInst {
	return &STVInst{
		instructions.Instruction{
			Opcode: instructions.STV,
		},
		0,
	}
}

// FetchOperands gets the opcode operands.
func (i *STVInst) FetchOperands(fetch instructions.FetchOperandsImplementation) error {
	var err error
	i.VarIndex, err = fetch.Next()

	return err
}

// Execute receives the context and runs the opcode.
func (i *STVInst) Execute(cp *cp.CP, vars *vars.Vars, st *stack.Stack, stdin instructions.StdinInterface, stdout instructions.StdoutInterface) error {
	value, err := st.Pop()
	if err != nil {
		return err
	}

	vars.Set(i.VarIndex, value)

	return nil
}
