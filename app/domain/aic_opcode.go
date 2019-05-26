package opcodes

type AicOpcode struct {
	Instruction
}

func NewAicOpcode() *AicOpcode {
	return &AicOpcode{Instruction{"AIC", Aic, 1}}
}

func (d *AicOpcode) Execute(cp *CP, stack *Stack, op int) {
	cp.Add(CPItem(op))
}
