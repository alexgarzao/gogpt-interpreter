package bytecode

import (
	"testing"

	"github.com/alexgarzao/gogpt-interpreter/pkg/domain"
	"github.com/stretchr/testify/assert"
)

const MyFakeOpcode = 1

func TestBytecodeAddingOneInstruction(t *testing.T) {
	bc := New()
	bc.Add(MyFakeOpcode, 111)

	v, _ := bc.Get(0)
	assert.Equal(t, MyFakeOpcode, v)
	v, _ = bc.Get(1)
	assert.Equal(t, 111, v)

	assert.Equal(t, 2, bc.Len())
}

func TestBytecodeAddingAndFetchingBytecodes(t *testing.T) {
	bc := New()
	bc.Add(MyFakeOpcode, 111)
	bc.Add(MyFakeOpcode, 222)

	v, _ := bc.Get(0)
	assert.Equal(t, MyFakeOpcode, v)
	v, _ = bc.Get(1)
	assert.Equal(t, 111, v)

	v, _ = bc.Get(2)
	assert.Equal(t, MyFakeOpcode, v)
	v, _ = bc.Get(3)
	assert.Equal(t, 222, v)

	assert.Equal(t, 4, bc.Len())
}

func TestBytecodeEofError(t *testing.T) {
	bc := New()
	bc.Add(MyFakeOpcode, 111)

	_, err := bc.Get(2)
	assert.Equal(t, err, domain.ErrIndexNotFound)
}

func TestBytecodeAddWithoutInstance(t *testing.T) {
	var bc *Bytecode
	bc.Add(MyFakeOpcode, 111)

	assert.Equal(t, 0, bc.Len())
}
