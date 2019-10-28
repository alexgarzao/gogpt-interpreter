package bce

import (
	"fmt"

	"github.com/alexgarzao/gogpt-interpreter/gogpt/model/bytecode"
	"github.com/alexgarzao/gogpt-interpreter/gogpt/model/constant_pool"
	"github.com/alexgarzao/gogpt-interpreter/gogpt/model/stack"
	"github.com/alexgarzao/gogpt-interpreter/gogpt/model/vars"
	"github.com/alexgarzao/gogpt-interpreter/gogpt/services/opcodes"
)

type BytecodeExecutor struct {
	bc           *bytecode.Bytecode
	instructions map[int]opcodes.InstructionImplementation
	ip           int
}

func NewBytecodeExecutor(bc *bytecode.Bytecode) *BytecodeExecutor {
	bce := &BytecodeExecutor{
		ip:           0,
		bc:           bc,
		instructions: make(map[int]opcodes.InstructionImplementation),
	}

	bce.instructions[opcodes.Nop] = opcodes.NewNopOpcode()
	bce.instructions[opcodes.Ldc] = opcodes.NewLdcOpcode()
	bce.instructions[opcodes.Call] = opcodes.NewCallOpcode()
	bce.instructions[opcodes.Ldv] = opcodes.NewLdvOpcode()
	bce.instructions[opcodes.Stv] = opcodes.NewStvOpcode()

	return bce
}

func (bce *BytecodeExecutor) Run(cp *constant_pool.CP, vars *vars.Vars, st *stack.Stack, stdin opcodes.StdinInterface, stdout opcodes.StdoutInterface) error {
	for {
		opcode, err := bce.Next()
		if err != nil {
			return nil
		}
		instruction, exist := bce.instructions[opcode]
		if !exist {
			return fmt.Errorf("Invalid opcode %d", opcode)
		}
		if instruction.GetOperandCount() == 1 {
			operand, err := bce.Next()
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

	return nil
}

func (bce *BytecodeExecutor) Next() (code int, err error) {
	code, err = bce.bc.Get(bce.ip)
	bce.ip += 1
	return
}