package nopi

import (
	"github.com/alexgarzao/gogpt-interpreter/pkg/model/cp"
	"github.com/alexgarzao/gogpt-interpreter/pkg/model/stack"
	"github.com/alexgarzao/gogpt-interpreter/pkg/model/vars"
	"github.com/alexgarzao/gogpt-interpreter/pkg/services/instructions"
)

// NOPInst is responsible for nothing.
type NOPInst struct {
	instructions.Instruction
}

// New creates a new NOPInst.
func New() *NOPInst {
	return &NOPInst{
		instructions.Instruction{
			Opcode: instructions.NOP,
		},
	}
}

// FetchOperands gets the opcode operands.
func (i *NOPInst) FetchOperands(fetch instructions.FetchOperandsImplementation) error {
	return nil
}

// Execute receives the context and runs the opcode.
func (i *NOPInst) Execute(cp *cp.CP, vars *vars.Vars, st *stack.Stack, stdin instructions.StdinInterface, stdout instructions.StdoutInterface) error {
	return nil
}
