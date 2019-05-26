package vm

import (
	"github.com/alexgarzao/gpt-interpreter/app/domain"
)

type BytecodeExecutor struct {
	instructions map[opcodes.Opcode]opcodes.InstructionImplementation
}

func NewBytecodeExecutor() *BytecodeExecutor {
	bce := &BytecodeExecutor{}
	bce.instructions = make(map[opcodes.Opcode]opcodes.InstructionImplementation)
	bce.instructions[opcodes.Aic] = opcodes.NewAicOpcode()
	bce.instructions[opcodes.Lic] = opcodes.NewLicOpcode()

	return bce
}

func (bce *BytecodeExecutor) Run(cp *opcodes.CP, st *opcodes.Stack, bc *opcodes.Bytecode) {
	ip := 0
	for ip < bc.Len() {
		opcode, _ := bc.Get(ip)
		ip++
		operand, _ := bc.Get(ip)
		ip++
		bce.instructions[(opcodes.Opcode)(opcode)].Execute(cp, st, (int)(operand))
	}
}
