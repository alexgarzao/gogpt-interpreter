package opcodes

type NopOpcode struct {
	Instruction
}

func NewNopOpcode() *NopOpcode {
	return &NopOpcode{Instruction{"NOP", Nop, 1}}
}

func (d *NopOpcode) FetchOperands(bc *Bytecode) error {
	return nil
}

func (d *NopOpcode) Execute(cp *CP, stack *Stack, stdout StdoutInterface) error {
	return nil
}
