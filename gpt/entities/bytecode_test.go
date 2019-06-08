package opcodes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBytecodeAddingOneInstruction(t *testing.T) {
	bc := NewBytecode()
	bc.Add(Ldc, 111)

	v, _ := bc.Get(0)
	assert.Equal(t, v, Ldc)
	v, _ = bc.Get(1)
	assert.Equal(t, v, 111)

	assert.Equal(t, bc.Len(), 2)
}

func TestBytecodeAddingAndFetchingBytecodes(t *testing.T) {
	bc := NewBytecode()
	bc.Add(Ldc, 111)
	bc.Add(Ldc, 222)

	v, _ := bc.Next()
	assert.Equal(t, v, Ldc)
	v, _ = bc.Next()
	assert.Equal(t, v, 111)

	v, _ = bc.Next()
	assert.Equal(t, v, Ldc)
	v, _ = bc.Next()
	assert.Equal(t, v, 222)

	assert.Equal(t, bc.Len(), 4)
}

func TestBytecodeEofError(t *testing.T) {
	bc := NewBytecode()
	bc.Add(Ldc, 111)

	bc.Next()
	bc.Next()
	_, err := bc.Next()
	assert.EqualError(t, err, "Index not found")
}
