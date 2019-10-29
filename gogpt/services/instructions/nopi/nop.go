package nopi

import (
	"github.com/alexgarzao/gogpt-interpreter/gogpt/model/cp"
	"github.com/alexgarzao/gogpt-interpreter/gogpt/model/stack"
	"github.com/alexgarzao/gogpt-interpreter/gogpt/model/vars"
	"github.com/alexgarzao/gogpt-interpreter/gogpt/services/instructions"
)

// NOPInst is responsible for nothing.
type NOPInst struct {
	instructions.Instruction
}

// New creates a new NOPInst.
func New() *NOPInst {
	return &NOPInst{instructions.Instruction{"NOP", instructions.NOP, 0}}
}

// GetOperandCount gets the numbers os opcode operands.
func (i *NOPInst) GetOperandCount() int {
	return i.OperandCount
}

// FetchOperands gets the opcode operands.
func (i *NOPInst) FetchOperands(op int) error {
	return nil
}

// Execute receives the context and runs the opcode.
func (i *NOPInst) Execute(cp *cp.CP, vars *vars.Vars, st *stack.Stack, stdin instructions.StdinInterface, stdout instructions.StdoutInterface) error {
	return nil
}
