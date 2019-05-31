package opcodes

type CallOpcode struct {
	Instruction
	CpIndex int
}

func NewCallOpcode() *CallOpcode {
	return &CallOpcode{Instruction{"CALL", Call, 1}, 0}
}

func (i *CallOpcode) FetchOperands(bc *Bytecode) error {
	var err error
	i.CpIndex, err = bc.Next()
	return err
}

func (i *CallOpcode) Execute(cp *CP, stack *Stack, stdout StdoutInterface) error {
	cpv, err := cp.Get(i.CpIndex)
	if err != nil {
		return err
	}

	if cpv == "io.println" {
		stv, err := stack.Pop()
		if err != nil {
			return err
		}

		stdout.Println(stv)
	}

	return nil
}
