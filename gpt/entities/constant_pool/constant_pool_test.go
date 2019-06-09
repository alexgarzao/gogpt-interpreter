package constant_pool

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCpAddIntConstants(t *testing.T) {
	cp := NewCp()
	assert.Equal(t, cp.Add(123), 0)
	assert.Equal(t, cp.Add(456), 1)
}

func TestCpGetIntConstants(t *testing.T) {
	cp := NewCp()
	assert.Equal(t, cp.Add(123), 0)
	assert.Equal(t, cp.Add(456), 1)
	v, _ := cp.Get(0)
	assert.Equal(t, v, CPItem(123))
	v, _ = cp.Get(1)
	assert.Equal(t, v, CPItem(456))
}

func TestCpGetIntConstantsError(t *testing.T) {
	cp := NewCp()
	assert.Equal(t, cp.Add(123), 0)
	assert.Equal(t, cp.Add(456), 1)
	v, err := cp.Get(0)
	assert.Equal(t, v, CPItem(123))
	assert.NoError(t, err)
	v, err = cp.Get(1)
	assert.Equal(t, v, CPItem(456))
	assert.NoError(t, err)
	v, err = cp.Get(2)
	assert.EqualError(t, err, "Index not found")
}

func TestCpAddStrConstants(t *testing.T) {
	cp := NewCp()
	assert.Equal(t, cp.Add("ABC"), 0)
	assert.Equal(t, cp.Add("DEF"), 1)
}

func TestCpGetStrConstants(t *testing.T) {
	cp := NewCp()
	assert.Equal(t, cp.Add("ABC"), 0)
	assert.Equal(t, cp.Add("DEF"), 1)
	v, _ := cp.Get(0)
	assert.Equal(t, v, CPItem("ABC"))
	v, _ = cp.Get(1)
	assert.Equal(t, v, CPItem("DEF"))
}

func TestCpGetStrConstantsError(t *testing.T) {
	cp := NewCp()
	assert.Equal(t, cp.Add("ABC"), 0)
	assert.Equal(t, cp.Add("DEF"), 1)
	v, err := cp.Get(0)
	assert.Equal(t, v, CPItem("ABC"))
	assert.NoError(t, err)
	v, err = cp.Get(1)
	assert.Equal(t, v, CPItem("DEF"))
	assert.NoError(t, err)
	v, err = cp.Get(2)
	assert.EqualError(t, err, "Index not found")
}

func TestCpAddingIntAndStrConstants(t *testing.T) {
	cp := NewCp()
	assert.Equal(t, cp.Add(123), 0)
	assert.Equal(t, cp.Add("456"), 1)
	v, _ := cp.Get(0)
	assert.Equal(t, v, CPItem(123))
	v, _ = cp.Get(1)
	assert.Equal(t, v, CPItem("456"))
}

func TestCpAddingDuplicatedValues(t *testing.T) {
	cp := NewCp()
	assert.Equal(t, 0, cp.Add(123))
	assert.Equal(t, 0, cp.Add(123))
	assert.Equal(t, 1, cp.Add(456))
	assert.Equal(t, 0, cp.Add(123))
}
