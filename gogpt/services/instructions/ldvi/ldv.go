package ldvi

import (
	"github.com/alexgarzao/gogpt-interpreter/gogpt/model/cp"
	"github.com/alexgarzao/gogpt-interpreter/gogpt/model/stack"
	"github.com/alexgarzao/gogpt-interpreter/gogpt/model/vars"
	"github.com/alexgarzao/gogpt-interpreter/gogpt/services/instructions"
)

// LDVInst is responsible for push a variable content into the stack.
type LDVInst struct {
	instructions.Instruction
	VarIndex int
}

// New creates a new LDVInst.
func New() *LDVInst {
	return &LDVInst{
		instructions.Instruction{
			Name:   "LDV",
			Opcode: instructions.LDV,
		},
		0,
	}
}

// FetchOperands gets the opcode operands.
func (i *LDVInst) FetchOperands(fetch instructions.FetchOperandsImplementation) error {
	var err error
	i.VarIndex, err = fetch.Next()

	return err
}

// Execute receives the context and runs the opcode.
func (i *LDVInst) Execute(cp *cp.CP, vars *vars.Vars, st *stack.Stack, stdin instructions.StdinInterface, stdout instructions.StdoutInterface) error {
	value, err := vars.Get(i.VarIndex)
	if err != nil {
		return err
	}

	st.Push(value)

	return nil
}