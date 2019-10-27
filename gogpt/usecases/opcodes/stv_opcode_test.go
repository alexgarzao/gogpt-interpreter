package opcodes

import (
	"testing"

	"github.com/alexgarzao/gogpt-interpreter/gogpt/entities/constant_pool"
	"github.com/alexgarzao/gogpt-interpreter/gogpt/entities/stack"
	"github.com/alexgarzao/gogpt-interpreter/gogpt/entities/vars"

	"github.com/alexgarzao/gogpt-interpreter/gogpt/adapters"
	"github.com/stretchr/testify/assert"
)

func TestValidStvInt(t *testing.T) {
	// CP map:
	//		0: (INT) 123
	// VAR map:
	//		0: (INT) Value
	cp := constant_pool.NewCp()
	cpIndex := cp.Add(123)
	st := stack.NewStack()
	stdin := adapters.NewFakeStdin()
	stdout := adapters.NewFakeStdout()
	vars := vars.NewVars()
	varIndex := 0

	// LDC 0
	j := NewLdcOpcode()
	j.CpIndex = cpIndex
	j.Execute(cp, vars, st, stdin, stdout)
	stv, _ := st.Top()
	assert.Equal(t, stv, stack.StackItem(123))

	// STV 0
	stvOpcode := NewStvOpcode()
	stvOpcode.VarIndex = varIndex
	stvOpcode.Execute(cp, vars, st, stdin, stdout)
	vv, _ := vars.Get(varIndex)
	assert.Equal(t, vv, constant_pool.CPItem(123))
	assert.Equal(t, 0, st.Size())
}

func TestValidStvStr(t *testing.T) {
	// CP map:
	//		0: (STR) ABC
	// VAR map:
	//		0: (STR) Value
	cp := constant_pool.NewCp()
	cpIndex := cp.Add("ABC")
	st := stack.NewStack()
	stdin := adapters.NewFakeStdin()
	stdout := adapters.NewFakeStdout()
	vars := vars.NewVars()
	varIndex := vars.Add()

	// LDC 0
	j := NewLdcOpcode()
	j.CpIndex = cpIndex
	j.Execute(cp, vars, st, stdin, stdout)
	stv, _ := st.Top()
	assert.Equal(t, stv, stack.StackItem("ABC"))

	// STV 0
	stvOpcode := NewStvOpcode()
	stvOpcode.VarIndex = varIndex
	stvOpcode.Execute(cp, vars, st, stdin, stdout)
	vv, _ := vars.Get(varIndex)
	assert.Equal(t, vv, constant_pool.CPItem("ABC"))
	assert.Equal(t, 0, st.Size())
}
