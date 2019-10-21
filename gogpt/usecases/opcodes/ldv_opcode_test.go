package opcodes

import (
	"testing"

	"github.com/alexgarzao/gogpt-interpreter/gogpt/entities/constant_pool"
	"github.com/alexgarzao/gogpt-interpreter/gogpt/entities/stack"
	"github.com/alexgarzao/gogpt-interpreter/gogpt/entities/vars"

	"github.com/alexgarzao/gogpt-interpreter/gogpt/adapters"
	"github.com/stretchr/testify/assert"
)

func TestValidLdvInt(t *testing.T) {
	// CP map:
	//		0: (INT) 123
	// VAR map:
	//		0: (INT) Value
	cp := constant_pool.NewCp()
	cpIndex := cp.Add(123)
	st := stack.NewStack()
	stdout := adapters.NewFakeStdout()
	vars := vars.NewVars()
	varIndex := 0

	// LDC 0
	j := NewLdcOpcode()
	j.CpIndex = cpIndex
	j.Execute(cp, vars, st, stdout)

	// STV 0
	stvOpcode := NewStvOpcode()
	stvOpcode.VarIndex = varIndex
	stvOpcode.Execute(cp, vars, st, stdout)
	vv, _ := vars.Get(varIndex)
	assert.Equal(t, vv, constant_pool.CPItem(123))
	assert.Equal(t, 0, st.Size())

	// LDV 0
	ldvOpcode := NewLdvOpcode()
	ldvOpcode.VarIndex = varIndex
	ldvOpcode.Execute(cp, vars, st, stdout)

	stv, _ := st.Top()
	assert.Equal(t, stv, stack.StackItem(123))
}

func TestValidLdvStr(t *testing.T) {
	// CP map:
	//		0: (STR) ABC
	// VAR map:
	//		0: (STR) Value
	cp := constant_pool.NewCp()
	cpIndex := cp.Add("ABC")
	st := stack.NewStack()
	stdout := adapters.NewFakeStdout()
	vars := vars.NewVars()
	varIndex := 0

	// LDC 0
	j := NewLdcOpcode()
	j.CpIndex = cpIndex
	j.Execute(cp, vars, st, stdout)

	// STV 0
	stvOpcode := NewStvOpcode()
	stvOpcode.VarIndex = varIndex
	stvOpcode.Execute(cp, vars, st, stdout)
	vv, _ := vars.Get(varIndex)
	assert.Equal(t, vv, constant_pool.CPItem("ABC"))
	assert.Equal(t, 0, st.Size())

	// LDV 0
	ldvOpcode := NewLdvOpcode()
	ldvOpcode.VarIndex = varIndex
	ldvOpcode.Execute(cp, vars, st, stdout)

	stv, _ := st.Top()
	assert.Equal(t, stv, stack.StackItem("ABC"))
}
