package opcodes

type LicOpcode struct {
	Instruction
}

func NewLicOpcode(cp *int, stack *int) *LicOpcode {
	return &LicOpcode{Instruction{"LIC", Lic, 1, cp, stack}}
}

func (d *LicOpcode) Execute(op int) {
	*d.Stack = *d.CP
}
