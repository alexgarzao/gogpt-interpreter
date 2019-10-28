package bce

import (
	"fmt"

	"github.com/alexgarzao/gogpt-interpreter/gogpt/model/bytecode"
	"github.com/alexgarzao/gogpt-interpreter/gogpt/model/constant_pool"
	"github.com/alexgarzao/gogpt-interpreter/gogpt/model/stack"
	"github.com/alexgarzao/gogpt-interpreter/gogpt/model/vars"
	"github.com/alexgarzao/gogpt-interpreter/gogpt/services/opcodes"
)

// BytecodeExecutor is responsible for execute the bytecode.
type BytecodeExecutor struct {
	bc           *bytecode.Bytecode
	instructions map[int]opcodes.InstructionImplementation
	ip           int
}

// NewBytecodeExecutor creates a new BytecodeExecutor.
func NewBytecodeExecutor(bc *bytecode.Bytecode) *BytecodeExecutor {
	bce := &BytecodeExecutor{
		ip:           0,
		bc:           bc,
		instructions: make(map[int]opcodes.InstructionImplementation),
	}

	bce.instructions[opcodes.NOP] = opcodes.NewNOPOpcode()
	bce.instructions[opcodes.LDC] = opcodes.NewLDCOpcode()
	bce.instructions[opcodes.CALL] = opcodes.NewCALLOpcode()
	bce.instructions[opcodes.LDV] = opcodes.NewLDVOpcode()
	bce.instructions[opcodes.STV] = opcodes.NewSTVOpcode()

	return bce
}

// Run receives the context (constant pool, vars, stack, stdin and stdout) and runs the bytecode.
func (bce *BytecodeExecutor) Run(cp *constant_pool.CP, vars *vars.Vars, st *stack.Stack, stdin opcodes.StdinInterface, stdout opcodes.StdoutInterface) error {
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
