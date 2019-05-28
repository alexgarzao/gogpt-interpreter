package opcodes

type AddOpcode struct {
	Instruction
}

func NewAddOpcode() *AddOpcode {
	return &AddOpcode{Instruction{"ADD", Add, 0}}
}

func (i *AddOpcode) Execute(cp *CP, stack *Stack, op int) {
	// TODO: op is useless int this opcode.
	op2, _ := stack.Pop()
	op1, _ := stack.Pop()
	res := op1 + op2
	stack.Push(StackItem(res))
}
