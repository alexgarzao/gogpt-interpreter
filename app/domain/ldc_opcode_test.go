package opcodes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidLdcInt123(t *testing.T) {
	// CP map:
	//		0: (INT) 123
	cp := NewCp()
	cp_index := cp.Add(123)
	stack := NewStack()
	cpv, _ := cp.Get(0)
	assert.Equal(t, cpv, CPItem(123))
	stv, _ := stack.Top()
	assert.Equal(t, stv, StackItem(0))

	// LDC 0
	j := NewLdcOpcode()
	j.CpIndex = BytecodeItem(cp_index)
	j.Execute(cp, stack)
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
	cpv, _ := cp.Get(0)
	assert.Equal(t, cpv, CPItem("ABC"))
	stv, _ := stack.Top()
	assert.Equal(t, stv, StackItem(0))

	// LDC 0
	j := NewLdcOpcode()
	j.CpIndex = BytecodeItem(cpIndex)
	j.Execute(cp, stack)
	cpv, _ = cp.Get(0)
	assert.Equal(t, cpv, CPItem("ABC"))
	stv, _ = stack.Top()
	assert.Equal(t, stv, StackItem("ABC"))
}
