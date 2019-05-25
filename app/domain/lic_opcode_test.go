package opcodes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidLic(t *testing.T) {
	cp := NewCp()
	stack := NewStack()
	// AIC 1
	i := NewAicOpcode(cp, stack)
	i.Execute(1)
	cpv, _ := cp.Get(0)
	assert.Equal(t, cpv, CpItem(1))
	stv, _ := stack.Top()
	assert.Equal(t, stv, StackItem(0))

	// LIC 0
	j := NewLicOpcode(cp, stack)
	j.Execute(0)
	cpv, _ = cp.Get(0)
	assert.Equal(t, cpv, CpItem(1))
	stv, _ = stack.Top()
	assert.Equal(t, stv, StackItem(1))
}
