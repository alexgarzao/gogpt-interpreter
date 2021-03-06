package stack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStackPushOneValue(t *testing.T) {
	s := New()
	s.Push(111)
	v, _ := s.Top()
	assert.Equal(t, 111, v)
}

func TestStackPushTwoValues(t *testing.T) {
	s := New()
	s.Push(111)
	s.Push(222)
	v, _ := s.Top()
	assert.Equal(t, 222, v)
}

func TestStackPushAndPopOneValue(t *testing.T) {
	s := New()
	s.Push(111)
	v, _ := s.Pop()
	assert.Equal(t, 111, v)
}

func TestStackPushAndPopTwoValues(t *testing.T) {
	s := New()
	s.Push(111)
	s.Push(222)
	v, _ := s.Pop()
	assert.Equal(t, 222, v)
	v, _ = s.Pop()
	assert.Equal(t, 111, v)
}

func TestStackTop(t *testing.T) {
	s := New()
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
	s := New()
	_, err := s.Pop()
	assert.EqualError(t, err, "Stack underflow")
	s.Push(111)
	_, err = s.Pop()
	assert.NoError(t, err)
	_, err = s.Pop()
	assert.EqualError(t, err, "Stack underflow")
}

func TestStackPushAndPopTwoValuesCheckingSize(t *testing.T) {
	s := New()
	assert.Equal(t, 0, s.Size())

	s.Push(111)
	assert.Equal(t, 1, s.Size())

	s.Push(222)
	assert.Equal(t, 2, s.Size())

	v, _ := s.Pop()
	assert.Equal(t, 222, v)
	assert.Equal(t, 1, s.Size())

	v, _ = s.Pop()
	assert.Equal(t, 111, v)
	assert.Equal(t, 0, s.Size())
}
