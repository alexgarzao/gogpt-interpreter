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
	CP           *int
	Stack        *int
}

type InstructionImplementation interface {
	Execute(op int)
}
