package opcodes

type LdcOpcode struct {
	Instruction
}

func NewLdcOpcode() *LdcOpcode {
	return &LdcOpcode{Instruction{"LDC", Ldc, 1}}
}

func (d *LdcOpcode) Execute(cp *CP, stack *Stack, op int) {
	cpv, _ := cp.Get(op)
	stack.Push(StackItem(cpv))
}
