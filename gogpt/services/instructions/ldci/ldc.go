package ldci

import (
	"github.com/alexgarzao/gogpt-interpreter/gogpt/model/cp"
	"github.com/alexgarzao/gogpt-interpreter/gogpt/model/stack"
	"github.com/alexgarzao/gogpt-interpreter/gogpt/model/vars"
	"github.com/alexgarzao/gogpt-interpreter/gogpt/services/instructions"
)

// LDCInst is responsible for push a constant into the stack.
type LDCInst struct {
	instructions.Instruction
	CpIndex int
}

// New creates a new LDCInst.
func New() *LDCInst {
	return &LDCInst{
		instructions.Instruction{
			Name:   "LDC",
			Opcode: instructions.LDC,
		},
		0,
	}
}

// FetchOperands gets the opcode operands.
func (i *LDCInst) FetchOperands(fetch instructions.FetchOperandsImplementation) error {
	var err error
	i.CpIndex, err = fetch.Next()

	return err
}

// Execute receives the context and runs the opcode.
func (i *LDCInst) Execute(cp *cp.CP, vars *vars.Vars, st *stack.Stack, stdin instructions.StdinInterface, stdout instructions.StdoutInterface) error {
	cpv, err := cp.Get(i.CpIndex)
	if err != nil {
		return err
	}

	st.Push(cpv)

	return nil
}