package opcodes

import (
	"testing"

	adapters "github.com/alexgarzao/gpt-interpreter/gpt/adapters"
	"github.com/stretchr/testify/assert"
)

func TestValidLdcInt123(t *testing.T) {
	// CP map:
	//		0: (INT) 123
	cp := NewCp()
	cpIndex := cp.Add(123)
	stack := NewStack()
	stdout := adapters.NewFakeStdout()
	cpv, _ := cp.Get(0)
	assert.Equal(t, cpv, CPItem(123))
	stv, _ := stack.Top()
	assert.Equal(t, stv, StackItem(0))

	// LDC 0
	j := NewLdcOpcode()
	j.CpIndex = cpIndex
	j.Execute(cp, stack, stdout)
	cpv, _ = cp.Get(0)
	assert.Equal(t, cpv, CPItem(123))
	stv, _ = stack.Top()
	assert.Equal(t, stv, StackItem(123))
}

func TestValidLdcABC(t *testing.T) {
	// CP map:
	//		0: STR: "ABC"
	cp := NewCp()
	cpIndex := cp.Add("ABC")
	stack := NewStack()
	stdout := adapters.NewFakeStdout()
	cpv, _ := cp.Get(0)
	assert.Equal(t, cpv, CPItem("ABC"))
	stv, _ := stack.Top()
	assert.Equal(t, stv, StackItem(0))

	// LDC 0
	j := NewLdcOpcode()
	j.CpIndex = cpIndex
	j.Execute(cp, stack, stdout)
	cpv, _ = cp.Get(0)
	assert.Equal(t, cpv, CPItem("ABC"))
	stv, _ = stack.Top()
	assert.Equal(t, stv, StackItem("ABC"))
}
