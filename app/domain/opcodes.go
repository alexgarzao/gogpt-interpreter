package opcodes

const (
	Nop int = iota
	Ldc
	Add
	Call
)

type Instruction struct {
	Name         string
	Opcode       int
	OperandCount int
}

type InstructionImplementation interface {
	FetchOperands(bytecode *Bytecode)
	Execute(cp *CP, stack *Stack)
}
