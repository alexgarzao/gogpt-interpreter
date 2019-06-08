package opcodes

import (
	"testing"

	adapters "github.com/alexgarzao/gpt-interpreter/gpt/adapters"
	"github.com/stretchr/testify/assert"
)

func TestValidNop(t *testing.T) {
	cp := NewCp()
	stack := NewStack()
	stdout := adapters.NewFakeStdout()

	// NOP
	i := NewNopOpcode()
	i.Execute(cp, stack, stdout)
	_, err := cp.Get(0)
	assert.EqualError(t, err, "Index not found")
	_, err = stack.Top()
	assert.EqualError(t, err, "Stack underflow")
}
