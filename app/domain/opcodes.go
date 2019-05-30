package opcodes

const (
	Nop Opcode = iota
	Ldc
	Add
	Call
)

type Opcode byte

type Instruction struct {
	Name         string
	Opcode       Opcode
	OperandCount int
}

type InstructionImplementation interface {
	FetchOperands(bytecode *Bytecode)
	Execute(cp *CP, stack *Stack)
}
