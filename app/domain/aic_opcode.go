package opcodes

type AicOpcode struct {
	Instruction
}

func NewAicOpcode(cp *int, stack *int) *AicOpcode {
	return &AicOpcode{Instruction{"AIC", Aic, 1, cp, stack}}
}

func (d *AicOpcode) Execute(op int) {
	*d.CP = op
}
