package calli

import (
	"strconv"

	"github.com/alexgarzao/gogpt-interpreter/gogpt/model/cp"
	"github.com/alexgarzao/gogpt-interpreter/gogpt/model/stack"
	"github.com/alexgarzao/gogpt-interpreter/gogpt/model/vars"
	"github.com/alexgarzao/gogpt-interpreter/gogpt/services/instructions"
)

// CALLInst is responsible for get the operand and call the lib function.
type CALLInst struct {
	instructions.Instruction
	CpIndex int
}

// New creates a new CALLInst.
func New() *CALLInst {
	return &CALLInst{
		instructions.Instruction{
			Name:   "CALL",
			Opcode: instructions.CALL,
		},
		0,
	}
}

// FetchOperands gets the opcode operands.
func (i *CALLInst) FetchOperands(fetch instructions.FetchOperandsImplementation) error {
	var err error
	i.CpIndex, err = fetch.Next()

	return err
}

// Execute receives the context and runs the opcode.
func (i *CALLInst) Execute(cp *cp.CP, vars *vars.Vars, st *stack.Stack, stdin instructions.StdinInterface, stdout instructions.StdoutInterface) error {
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