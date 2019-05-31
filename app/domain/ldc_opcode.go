package opcodes

type LdcOpcode struct {
	Instruction
	CpIndex int
}

func NewLdcOpcode() *LdcOpcode {
	return &LdcOpcode{Instruction{"LDC", Ldc, 1}, 0}
}

func (d *LdcOpcode) FetchOperands(bc *Bytecode) {
	d.CpIndex, _ = bc.Next()
}

func (d *LdcOpcode) Execute(cp *CP, stack *Stack, stdout StdoutInterface) {
	cpv, _ := cp.Get(d.CpIndex)
	stack.Push(StackItem(cpv))
}
