package bce

import (
	"fmt"

	"github.com/alexgarzao/gogpt-interpreter/gogpt/model/bytecode"
	"github.com/alexgarzao/gogpt-interpreter/gogpt/model/cp"
	"github.com/alexgarzao/gogpt-interpreter/gogpt/model/stack"
	"github.com/alexgarzao/gogpt-interpreter/gogpt/model/vars"
	"github.com/alexgarzao/gogpt-interpreter/gogpt/services/instructions"
	"github.com/alexgarzao/gogpt-interpreter/gogpt/services/instructions/calli"
	"github.com/alexgarzao/gogpt-interpreter/gogpt/services/instructions/ldci"
	"github.com/alexgarzao/gogpt-interpreter/gogpt/services/instructions/ldvi"
	"github.com/alexgarzao/gogpt-interpreter/gogpt/services/instructions/nopi"
	"github.com/alexgarzao/gogpt-interpreter/gogpt/services/instructions/stvi"
)

// BytecodeExecutor is responsible for execute the bytecode.
type BytecodeExecutor struct {
	bc           *bytecode.Bytecode
	instructions map[int]instructions.InstructionImplementation
	ip           int
}

// New creates a new BytecodeExecutor.
func New(bc *bytecode.Bytecode) *BytecodeExecutor {
	bce := &BytecodeExecutor{
		ip:           0,
		bc:           bc,
		instructions: make(map[int]instructions.InstructionImplementation),
	}

	bce.instructions[instructions.NOP] = nopi.New()
	bce.instructions[instructions.LDC] = ldci.New()
	bce.instructions[instructions.CALL] = calli.New()
	bce.instructions[instructions.LDV] = ldvi.New()
	bce.instructions[instructions.STV] = stvi.New()

	return bce
}

// Run receives the context (constant pool, vars, stack, stdin and stdout) and runs the bytecode.
func (bce *BytecodeExecutor) Run(cp *cp.CP, vars *vars.Vars, st *stack.Stack, stdin instructions.StdinInterface, stdout instructions.StdoutInterface) error {
	for {
		opcode, err := bce.next()
		if err != nil {
			return nil
		}
		instruction, exist := bce.instructions[opcode]
		if !exist {
			return fmt.Errorf("Invalid opcode %d", opcode)
		}
		if instruction.GetOperandCount() == 1 {
			operand, err := bce.next()
			if err != nil {
				return err
			}
			err = instruction.FetchOperands(operand)
			if err != nil {
				return err
			}
		}
		err = instruction.Execute(cp, vars, st, stdin, stdout)
		if err != nil {
			return err
		}
	}
}

func (bce *BytecodeExecutor) next() (code int, err error) {
	code, err = bce.bc.Get(bce.ip)
	bce.ip++
	return
}
