package opcodes

type LicOpcode struct {
	Instruction
}

func NewLicOpcode(cp *Cp, stack *Stack) *LicOpcode {
	return &LicOpcode{Instruction{"LIC", Lic, 1, cp, stack}}
}

func (d *LicOpcode) Execute(op int) {
	cpv, _ := d.CP.Get(op)
	d.Stack.Push(StackItem(cpv))
}
