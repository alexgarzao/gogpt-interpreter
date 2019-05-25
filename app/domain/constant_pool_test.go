package opcodes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCpAddConstant(t *testing.T) {
	cp := NewCp()
	assert.Equal(t, cp.Add(123), 0)
	assert.Equal(t, cp.Add(456), 1)
}

func TestCpGetConstant(t *testing.T) {
	cp := NewCp()
	assert.Equal(t, cp.Add(123), 0)
	assert.Equal(t, cp.Add(456), 1)
	v, _ := cp.Get(0)
	assert.Equal(t, v, CpItem(123))
	v, _ = cp.Get(1)
	assert.Equal(t, v, CpItem(456))
}

func TestCpGetConstantError(t *testing.T) {
	cp := NewCp()
	assert.Equal(t, cp.Add(123), 0)
	assert.Equal(t, cp.Add(456), 1)
	v, err := cp.Get(0)
	assert.Equal(t, v, CpItem(123))
	assert.NoError(t, err)
	v, err = cp.Get(1)
	assert.Equal(t, v, CpItem(456))
	assert.NoError(t, err)
	v, err = cp.Get(2)
	assert.EqualError(t, err, "Index not found")
}
