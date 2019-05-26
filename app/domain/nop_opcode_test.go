package opcodes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidNop(t *testing.T) {
	cp := NewCp()
	stack := NewStack()
	// NOP
	i := NewNopOpcode()
	i.Execute(cp, stack, 1)
	_, err := cp.Get(0)
	assert.EqualError(t, err, "Index not found")
	_, err = stack.Top()
	assert.EqualError(t, err, "Stack underflow")
}
