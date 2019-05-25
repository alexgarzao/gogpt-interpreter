package opcodes

const (
	Nop Opcode = iota
	Aic
	Lic
)

type Opcode byte

type Instruction struct {
	Name         string
	Opcode       Opcode
	OperandCount int
	CP           *Cp
	Stack        *Stack
}

type InstructionImplementation interface {
	Execute(op int)
}
