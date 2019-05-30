package vm

import (
	opcodes "github.com/alexgarzao/gpt-interpreter/app/domain"
)

type BytecodeExecutor struct {
	instructions map[opcodes.Opcode]opcodes.InstructionImplementation
}

func NewBytecodeExecutor() *BytecodeExecutor {
	bce := &BytecodeExecutor{}
	bce.instructions = make(map[opcodes.Opcode]opcodes.InstructionImplementation)
	bce.instructions[opcodes.Nop] = opcodes.NewNopOpcode()
	bce.instructions[opcodes.Ldc] = opcodes.NewLdcOpcode()
	bce.instructions[opcodes.Call] = opcodes.NewCallOpcode()

	return bce
}

func (bce *BytecodeExecutor) Run(cp *opcodes.CP, st *opcodes.Stack, bc *opcodes.Bytecode) {
	for bc.IP < bc.Len() {
		opcode, _ := bc.Next()
		instruction := bce.instructions[(opcodes.Opcode)(opcode)]
		instruction.FetchOperands(bc)
		instruction.Execute(cp, st)
	}
}
