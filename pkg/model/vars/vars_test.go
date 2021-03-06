package vars

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVarsOneValue(t *testing.T) {
	vars := New()
	vars.Set(0, 111)
	v, _ := vars.Get(0)
	assert.Equal(t, 111, v)
}

func TestVarsTwoValues(t *testing.T) {
	vars := New()
	vars.Set(0, 111)
	vars.Set(1, 222)
	value, _ := vars.Get(0)
	assert.Equal(t, 111, value)
	value, _ = vars.Get(1)
	assert.Equal(t, 222, value)
}

func TestVarsTwoTypeValues(t *testing.T) {
	vars := New()
	vars.Set(0, 111)
	vars.Set(1, "222")
	value, _ := vars.Get(0)
	assert.Equal(t, 111, value)
	value, _ = vars.Get(1)
	assert.Equal(t, "222", value)
}

func TestVarsWitInvalidIndex(t *testing.T) {
	vars := New()
	vars.Set(0, 111)
	value, err := vars.Get(1)
	assert.EqualError(t, err, "Variable index undefined")
	assert.Equal(t, 0, value)
}

func TestVarsTwoTypeValuesWithAdd(t *testing.T) {
	vars := New()
	v1index := vars.Add()
	vars.Set(v1index, 111)
	v2index := vars.Add()
	vars.Set(v2index, "222")
	value, _ := vars.Get(v1index)
	assert.Equal(t, 111, value)
	value, _ = vars.Get(v2index)
	assert.Equal(t, "222", value)
}
