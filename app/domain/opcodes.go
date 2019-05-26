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
}

type InstructionImplementation interface {
	Execute(cp *CP, stack *Stack, op int)
}
