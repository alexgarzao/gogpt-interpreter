package cp

import (
	"testing"

	"github.com/alexgarzao/gogpt-interpreter/pkg/domain"
	"github.com/stretchr/testify/assert"
)

func TestCpAddIntConstants(t *testing.T) {
	cp := New()
	assert.Equal(t, 0, cp.Add(123))
	assert.Equal(t, 1, cp.Add(456))
}

func TestCpGetIntConstants(t *testing.T) {
	cp := New()
	assert.Equal(t, 0, cp.Add(123))
	assert.Equal(t, 1, cp.Add(456))
	v, _ := cp.Get(0)
	assert.Equal(t, 123, v)
	v, _ = cp.Get(1)
	assert.Equal(t, 456, v)
}

func TestCpGetIntConstantsError(t *testing.T) {
	cp := New()
	assert.Equal(t, 0, cp.Add(123))
	assert.Equal(t, 1, cp.Add(456))
	v, err := cp.Get(0)
	assert.Equal(t, 123, v)
	assert.NoError(t, err)
	v, err = cp.Get(1)
	assert.Equal(t, 456, v)
	assert.NoError(t, err)
	v, err = cp.Get(2)
	assert.Equal(t, err, domain.ErrIndexNotFound)
}

func TestCpAddStrConstants(t *testing.T) {
	cp := New()
	assert.Equal(t, 0, cp.Add("ABC"))
	assert.Equal(t, 1, cp.Add("DEF"))
}

func TestCpGetStrConstants(t *testing.T) {
	cp := New()
	assert.Equal(t, 0, cp.Add("ABC"))
	assert.Equal(t, 1, cp.Add("DEF"))
	v, _ := cp.Get(0)
	assert.Equal(t, "ABC", v)
	v, _ = cp.Get(1)
	assert.Equal(t, "DEF", v)
}

func TestCpGetStrConstantsError(t *testing.T) {
	cp := New()
	assert.Equal(t, 0, cp.Add("ABC"))
	assert.Equal(t, 1, cp.Add("DEF"))
	v, err := cp.Get(0)
	assert.Equal(t, "ABC", v)
	assert.NoError(t, err)
	v, err = cp.Get(1)
	assert.Equal(t, "DEF", v)
	assert.NoError(t, err)
	v, err = cp.Get(2)
	assert.Equal(t, err, domain.ErrIndexNotFound)
}

func TestCpAddingIntAndStrConstants(t *testing.T) {
	cp := New()
	assert.Equal(t, 0, cp.Add(123))
	assert.Equal(t, 1, cp.Add("456"))
	v, _ := cp.Get(0)
	assert.Equal(t, 123, v)
	v, _ = cp.Get(1)
	assert.Equal(t, "456", v)
}

func TestCpAddingDuplicatedValues(t *testing.T) {
	cp := New()
	assert.Equal(t, 0, cp.Add(123))
	assert.Equal(t, 0, cp.Add(123))
	assert.Equal(t, 1, cp.Add(456))
	assert.Equal(t, 0, cp.Add(123))
}

func TestCPAddWithoutInstance(t *testing.T) {
	var cp *CP
	cp.Add(111)

	assert.Equal(t, -1, cp.Find(111))
}
