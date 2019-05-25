package opcodes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStackPushOneValue(t *testing.T) {
	s := NewStack()
	s.Push(111)
	v, _ := s.Top()
	assert.Equal(t, v, Item(111))
}

func TestStackPushTwoValues(t *testing.T) {
	s := NewStack()
	s.Push(111)
	s.Push(222)
	v, _ := s.Top()
	assert.Equal(t, v, Item(222))
}

func TestStackPushAndPopOneValue(t *testing.T) {
	s := NewStack()
	s.Push(111)
	v, _ := s.Pop()
	assert.Equal(t, v, Item(111))
}

func TestStackPushAndPopTwoValues(t *testing.T) {
	s := NewStack()
	s.Push(111)
	s.Push(222)
	v, _ := s.Pop()
	assert.Equal(t, v, Item(222))
	v, _ = s.Pop()
	assert.Equal(t, v, Item(111))
}

func TestStackTop(t *testing.T) {
	s := NewStack()
	_, err := s.Top()
	assert.EqualError(t, err, "Stack underflow")
	s.Push(111)
	_, err = s.Top()
	assert.NoError(t, err)
	s.Push(222)
	_, err = s.Top()
	assert.NoError(t, err)
	_, _ = s.Pop()
	_, err = s.Top()
	assert.NoError(t, err)
	_, _ = s.Pop()
	_, err = s.Top()
	assert.EqualError(t, err, "Stack underflow")
}

func TestStackPopError(t *testing.T) {
	s := NewStack()
	_, err := s.Pop()
	assert.EqualError(t, err, "Stack underflow")
	s.Push(111)
	_, err = s.Pop()
	assert.NoError(t, err)
	_, err = s.Pop()
	assert.EqualError(t, err, "Stack underflow")
}
