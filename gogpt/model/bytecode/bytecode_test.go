package bytecode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const MyFakeOpcode = 1

func TestBytecodeAddingOneInstruction(t *testing.T) {
	bc := New()
	bc.Add(MyFakeOpcode, 111)

	v, _ := bc.Get(0)
	assert.Equal(t, v, MyFakeOpcode)
	v, _ = bc.Get(1)
	assert.Equal(t, v, 111)

	assert.Equal(t, bc.Len(), 2)
}

func TestBytecodeAddingAndFetchingBytecodes(t *testing.T) {
	bc := New()
	bc.Add(MyFakeOpcode, 111)
	bc.Add(MyFakeOpcode, 222)

	v, _ := bc.Get(0)
	assert.Equal(t, v, MyFakeOpcode)
	v, _ = bc.Get(1)
	assert.Equal(t, v, 111)

	v, _ = bc.Get(2)
	assert.Equal(t, v, MyFakeOpcode)
	v, _ = bc.Get(3)
	assert.Equal(t, v, 222)

	assert.Equal(t, bc.Len(), 4)
}

func TestBytecodeEofError(t *testing.T) {
	bc := New()
	bc.Add(MyFakeOpcode, 111)

	_, err := bc.Get(2)
	assert.EqualError(t, err, "Index not found")
}

func TestBytecodeAddWithoutInstance(t *testing.T) {
	var bc *Bytecode
	bc.Add(MyFakeOpcode, 111)

	assert.Equal(t, 0, bc.Len())
}
