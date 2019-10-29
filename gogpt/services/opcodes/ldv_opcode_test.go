package opcodes

import (
	"testing"

	"github.com/alexgarzao/gogpt-interpreter/gogpt/infrastructure"
	"github.com/alexgarzao/gogpt-interpreter/gogpt/model/cp"
	"github.com/alexgarzao/gogpt-interpreter/gogpt/model/stack"
	"github.com/alexgarzao/gogpt-interpreter/gogpt/model/vars"
	"github.com/stretchr/testify/assert"
)

func TestValidLdvInt(t *testing.T) {
	// CP map:
	//		0: (INT) 123
	// VAR map:
	//		0: (INT) Value
	cp := cp.NewCP()
	cpIndex := cp.Add(123)
	st := stack.NewStack()
	stdin := infrastructure.NewFakeStdin()
	stdout := infrastructure.NewFakeStdout()
	vars := vars.NewVars()
	varIndex := 0

	// LDC 0
	j := NewLDCOpcode()
	j.CpIndex = cpIndex
	j.Execute(cp, vars, st, stdin, stdout)

	// STV 0
	stvOpcode := NewSTVOpcode()
	stvOpcode.VarIndex = varIndex
	stvOpcode.Execute(cp, vars, st, stdin, stdout)
	vv, _ := vars.Get(varIndex)
	assert.Equal(t, vv, 123)
	assert.Equal(t, 0, st.Size())

	// LDV 0
	ldvOpcode := NewLDVOpcode()
	ldvOpcode.VarIndex = varIndex
	ldvOpcode.Execute(cp, vars, st, stdin, stdout)

	stv, _ := st.Top()
	assert.Equal(t, stv, 123)
}

func TestValidLdvStr(t *testing.T) {
	// CP map:
	//		0: (STR) ABC
	// VAR map:
	//		0: (STR) Value
	cp := cp.NewCP()
	cpIndex := cp.Add("ABC")
	st := stack.NewStack()
	stdin := infrastructure.NewFakeStdin()
	stdout := infrastructure.NewFakeStdout()
	vars := vars.NewVars()
	varIndex := 0

	// LDC 0
	j := NewLDCOpcode()
	j.CpIndex = cpIndex
	j.Execute(cp, vars, st, stdin, stdout)

	// STV 0
	stvOpcode := NewSTVOpcode()
	stvOpcode.VarIndex = varIndex
	stvOpcode.Execute(cp, vars, st, stdin, stdout)
	vv, _ := vars.Get(varIndex)
	assert.Equal(t, vv, "ABC")
	assert.Equal(t, 0, st.Size())

	// LDV 0
	ldvOpcode := NewLDVOpcode()
	ldvOpcode.VarIndex = varIndex
	ldvOpcode.Execute(cp, vars, st, stdin, stdout)

	stv, _ := st.Top()
	assert.Equal(t, stv, "ABC")
}
