package opcodes

type CallOpcode struct {
	Instruction
	CpIndex int
}

func NewCallOpcode() *CallOpcode {
	return &CallOpcode{Instruction{"CALL", Call, 1}, 0}
}

func (i *CallOpcode) FetchOperands(bc *Bytecode) {
	i.CpIndex, _ = bc.Next()
}

func (i *CallOpcode) Execute(cp *CP, stack *Stack, stdout StdoutInterface) {
	cpv, _ := cp.Get(i.CpIndex)
	if cpv == "io.println" {
		stv, _ := stack.Pop()
		stdout.Println(stv)
	}
}
