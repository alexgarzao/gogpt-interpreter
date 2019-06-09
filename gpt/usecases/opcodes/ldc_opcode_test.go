package opcodes

import (
	"github.com/alexgarzao/gpt-interpreter/gpt/entities/constant_pool"
	"github.com/alexgarzao/gpt-interpreter/gpt/entities/stack"
	"testing"

	"github.com/alexgarzao/gpt-interpreter/gpt/adapters"
	"github.com/stretchr/testify/assert"
)

func TestValidLdcInt123(t *testing.T) {
	// CP map:
	//		0: (INT) 123
	cp := constant_pool.NewCp()
	cpIndex := cp.Add(123)
	st := stack.NewStack()
	stdout := adapters.NewFakeStdout()
	cpv, _ := cp.Get(0)
	assert.Equal(t, cpv, constant_pool.CPItem(123))
	stv, _ := st.Top()
	assert.Equal(t, stv, stack.StackItem(0))

	// LDC 0
	j := NewLdcOpcode()
	j.CpIndex = cpIndex
	j.Execute(cp, st, stdout)
	cpv, _ = cp.Get(0)
	assert.Equal(t, cpv, constant_pool.CPItem(123))
	stv, _ = st.Top()
	assert.Equal(t, stv, stack.StackItem(123))
}

func TestValidLdcABC(t *testing.T) {
	// CP map:
	//		0: STR: "ABC"
	cp := constant_pool.NewCp()
	cpIndex := cp.Add("ABC")
	st := stack.NewStack()
	stdout := adapters.NewFakeStdout()
	cpv, _ := cp.Get(0)
	assert.Equal(t, cpv, constant_pool.CPItem("ABC"))
	stv, _ := st.Top()
	assert.Equal(t, stv, stack.StackItem(0))

	// LDC 0
	j := NewLdcOpcode()
	j.CpIndex = cpIndex
	j.Execute(cp, st, stdout)
	cpv, _ = cp.Get(0)
	assert.Equal(t, cpv, constant_pool.CPItem("ABC"))
	stv, _ = st.Top()
	assert.Equal(t, stv, stack.StackItem("ABC"))
}
