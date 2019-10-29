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
			Name:         "LDC",
			Opcode:       instructions.LDC,
			OperandCount: 1},
		0,
	}
}

// GetOperandCount gets the numbers os opcode operands.
func (i *LDCInst) GetOperandCount() int {
	return i.OperandCount
}

// FetchOperands gets the opcode operands.
func (i *LDCInst) FetchOperands(op int) error {
	i.CpIndex = op
	return nil
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
