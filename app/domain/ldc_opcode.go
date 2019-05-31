package opcodes

type LdcOpcode struct {
	Instruction
	CpIndex int
}

func NewLdcOpcode() *LdcOpcode {
	return &LdcOpcode{Instruction{"LDC", Ldc, 1}, 0}
}

func (d *LdcOpcode) FetchOperands(bc *Bytecode) error {
	var err error
	d.CpIndex, err = bc.Next()
	return err
}

func (d *LdcOpcode) Execute(cp *CP, stack *Stack, stdout StdoutInterface) error {
	cpv, err := cp.Get(d.CpIndex)
	if err != nil {
		return err
	}

	stack.Push(StackItem(cpv))

	return nil
}
