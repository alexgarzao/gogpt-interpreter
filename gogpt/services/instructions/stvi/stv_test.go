package stvi

import (
	"testing"

	"github.com/alexgarzao/gogpt-interpreter/gogpt/infrastructure"
	"github.com/alexgarzao/gogpt-interpreter/gogpt/model/cp"
	"github.com/alexgarzao/gogpt-interpreter/gogpt/model/stack"
	"github.com/alexgarzao/gogpt-interpreter/gogpt/model/vars"
	"github.com/alexgarzao/gogpt-interpreter/gogpt/services/instructions/ldci"
	"github.com/stretchr/testify/assert"
)

func TestValidStvInt(t *testing.T) {
	// CP map:
	//		0: (INT) 123
	// VAR map:
	//		0: (INT) Value
	cp := cp.New()
	cpIndex := cp.Add(123)
	st := stack.New()
	stdin := infrastructure.NewFakeStdin()
	stdout := infrastructure.NewFakeStdout()
	vars := vars.New()
	varIndex := 0

	// LDC 0
	ldc := ldci.New()
	ldc.CpIndex = cpIndex
	ldc.Execute(cp, vars, st, stdin, stdout)
	stackValue, _ := st.Top()
	assert.Equal(t, stackValue, 123)

	// STV 0
	stv := New()
	stv.VarIndex = varIndex
	stv.Execute(cp, vars, st, stdin, stdout)
	vv, _ := vars.Get(varIndex)
	assert.Equal(t, vv, 123)
	assert.Equal(t, 0, st.Size())
}

func TestValidStvStr(t *testing.T) {
	// CP map:
	//		0: (STR) ABC
	// VAR map:
	//		0: (STR) Value
	cp := cp.New()
	cpIndex := cp.Add("ABC")
	st := stack.New()
	stdin := infrastructure.NewFakeStdin()
	stdout := infrastructure.NewFakeStdout()
	vars := vars.New()
	varIndex := vars.Add()

	// LDC 0
	ldc := ldci.New()
	ldc.CpIndex = cpIndex
	ldc.Execute(cp, vars, st, stdin, stdout)
	stackValue, _ := st.Top()
	assert.Equal(t, stackValue, "ABC")

	// STV 0
	stv := New()
	stv.VarIndex = varIndex
	stv.Execute(cp, vars, st, stdin, stdout)
	vv, _ := vars.Get(varIndex)
	assert.Equal(t, vv, "ABC")
	assert.Equal(t, 0, st.Size())
}
