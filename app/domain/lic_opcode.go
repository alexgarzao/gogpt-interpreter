package opcodes

type LicOpcode struct {
	Instruction
}

func NewLicOpcode() *LicOpcode {
	return &LicOpcode{Instruction{"LIC", Lic, 1}}
}

func (d *LicOpcode) Execute(cp *CP, stack *Stack, op int) {
	cpv, _ := cp.Get(op)
	stack.Push(StackItem(cpv))
}
