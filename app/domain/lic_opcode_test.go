package opcodes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidLic123(t *testing.T) {
	// CP map:
	//		0: 123
	cp := NewCp()
	cp_index := cp.Add(123)
	stack := NewStack()
	cpv, _ := cp.Get(0)
	assert.Equal(t, cpv, CPItem(123))
	stv, _ := stack.Top()
	assert.Equal(t, stv, StackItem(0))

	// LIC 0
	j := NewLicOpcode()
	j.Execute(cp, stack, cp_index)
	cpv, _ = cp.Get(0)
	assert.Equal(t, cpv, CPItem(123))
	stv, _ = stack.Top()
	assert.Equal(t, stv, StackItem(123))
}
