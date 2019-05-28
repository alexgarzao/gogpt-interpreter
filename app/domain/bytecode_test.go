package opcodes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBytecodeAddingOneInstruction(t *testing.T) {
	bc := NewBytecode()
	bc.Add(Ldc, 111)

	v, _ := bc.Get(0)
	assert.Equal(t, v, BytecodeItem(Ldc))
	v, _ = bc.Get(1)
	assert.Equal(t, v, BytecodeItem(111))

	assert.Equal(t, bc.Len(), 2)
}
