package opcodes

import (
	"testing"

	"github.com/alexgarzao/gogpt-interpreter/gogpt/infrastructure"
	"github.com/alexgarzao/gogpt-interpreter/gogpt/model/constant_pool"
	"github.com/alexgarzao/gogpt-interpreter/gogpt/model/stack"
	"github.com/alexgarzao/gogpt-interpreter/gogpt/model/vars"
	"github.com/stretchr/testify/assert"
)

func TestValidLdcInt123(t *testing.T) {
	// CP map:
	//		0: (INT) 123
	cp := constant_pool.NewCP()
	cpIndex := cp.Add(123)
	vars := vars.NewVars()
	st := stack.NewStack()
	stdin := infrastructure.NewFakeStdin()
	stdout := infrastructure.NewFakeStdout()
	cpv, _ := cp.Get(0)
	assert.Equal(t, cpv, constant_pool.CPItem(123))
	stv, _ := st.Top()
	assert.Equal(t, stv, stack.StackItem(0))

	// LDC 0
	j := NewLdcOpcode()
	j.CpIndex = cpIndex
	j.Execute(cp, vars, st, stdin, stdout)
	cpv, _ = cp.Get(0)
	assert.Equal(t, cpv, constant_pool.CPItem(123))
	stv, _ = st.Top()
	assert.Equal(t, stv, stack.StackItem(123))
}

func TestValidLdcABC(t *testing.T) {
	// CP map:
	//		0: STR: "ABC"
	cp := constant_pool.NewCP()
	cpIndex := cp.Add("ABC")
	vars := vars.NewVars()
	st := stack.NewStack()
	stdin := infrastructure.NewFakeStdin()
	stdout := infrastructure.NewFakeStdout()
	cpv, _ := cp.Get(0)
	assert.Equal(t, cpv, constant_pool.CPItem("ABC"))
	stv, _ := st.Top()
	assert.Equal(t, stv, stack.StackItem(0))

	// LDC 0
	j := NewLdcOpcode()
	j.CpIndex = cpIndex
	j.Execute(cp, vars, st, stdin, stdout)
	cpv, _ = cp.Get(0)
	assert.Equal(t, cpv, constant_pool.CPItem("ABC"))
	stv, _ = st.Top()
	assert.Equal(t, stv, stack.StackItem("ABC"))
}
