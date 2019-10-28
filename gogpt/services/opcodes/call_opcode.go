package opcodes

import (
	"strconv"

	"github.com/alexgarzao/gogpt-interpreter/gogpt/model/constant_pool"
	"github.com/alexgarzao/gogpt-interpreter/gogpt/model/stack"
	"github.com/alexgarzao/gogpt-interpreter/gogpt/model/vars"
)

// CallOpcode is responsible for get the operand and call the lib function.
type CallOpcode struct {
	Instruction
	CpIndex int
}

// NewCallOpcode creates a new CallOpcode.
func NewCallOpcode() *CallOpcode {
	return &CallOpcode{Instruction{"CALL", Call, 1}, 0}
}

// GetOperandCount gets the numbers os opcode operands.
func (i *CallOpcode) GetOperandCount() int {
	return i.OperandCount
}

// FetchOperands gets the opcode operands.
func (i *CallOpcode) FetchOperands(op int) error {
	i.CpIndex = op
	return nil
}

// Execute receives the context and runs the opcode.
func (i *CallOpcode) Execute(cp *constant_pool.CP, vars *vars.Vars, st *stack.Stack, stdin StdinInterface, stdout StdoutInterface) error {
	cpv, err := cp.Get(i.CpIndex)
	if err != nil {
		return err
	}

	if cpv == "io.println" {
		stv, err := st.Pop()
		if err != nil {
			return err
		}

		argsCount := stv.(int)

		text := ""

		for argsCount > 0 {
			stv, err := st.Pop()
			if err != nil {
				return err
			}

			res := ""
			switch stv.(type) {
			case int:
				res = strconv.Itoa(stv.(int))
			case string:
				res = stv.(string)
			}

			text = res + text

			argsCount--
		}

		stdout.Println(text)
	} else if cpv == "io.readln" {
		text := stdin.Readln()
		st.Push(text)
	}

	return nil
}
