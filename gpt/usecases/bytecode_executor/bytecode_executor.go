package bce

import (
	"fmt"

	"github.com/alexgarzao/gpt-interpreter/gpt/entities/bytecode"
	"github.com/alexgarzao/gpt-interpreter/gpt/entities/constant_pool"
	"github.com/alexgarzao/gpt-interpreter/gpt/entities/stack"
	"github.com/alexgarzao/gpt-interpreter/gpt/usecases/opcodes"
)

type BytecodeExecutor struct {
	instructions map[int]opcodes.InstructionImplementation
	IP           int
}

func NewBytecodeExecutor() *BytecodeExecutor {
	bce := &BytecodeExecutor{}
	bce.instructions = make(map[int]opcodes.InstructionImplementation)
	bce.instructions[opcodes.Nop] = opcodes.NewNopOpcode()
	bce.instructions[opcodes.Ldc] = opcodes.NewLdcOpcode()
	bce.instructions[opcodes.Call] = opcodes.NewCallOpcode()
	bce.IP = 0

	return bce
}

func (bce *BytecodeExecutor) Run(cp *constant_pool.CP, st *stack.Stack, stdout opcodes.StdoutInterface, bc *bytecode.Bytecode) error {
	for bce.IP < bc.Len() {
		opcode, err := bce.Next(bc)
		if err != nil {
			return err
		}
		instruction, exist := bce.instructions[opcode]
		if !exist {
			return fmt.Errorf("Invalid opcode %d", opcode)
		}
		if instruction.GetOperandCount() == 1 {
			operand, err := bce.Next(bc)
			if err != nil {
				return err
			}
			err = instruction.FetchOperands(operand)
			if err != nil {
				return err
			}
		}
		err = instruction.Execute(cp, st, stdout)
		if err != nil {
			return err
		}
	}

	return nil
}

func (bce *BytecodeExecutor) Next(bc *bytecode.Bytecode) (code int, err error) {
	code, err = bc.Get(bce.IP)
	bce.IP += 1
	return
}
