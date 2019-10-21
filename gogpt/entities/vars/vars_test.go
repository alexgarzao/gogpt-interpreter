package vars

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVarsOneValue(t *testing.T) {
	s := NewVars()
	s.Set(0, 111)
	v, _ := s.Get(0)
	assert.Equal(t, v, VarsItem(111))
}

func TestVarsTwoValues(t *testing.T) {
	s := NewVars()
	s.Set(0, 111)
	s.Set(1, 222)
	v, _ := s.Get(0)
	assert.Equal(t, v, VarsItem(111))
	v, _ = s.Get(1)
	assert.Equal(t, v, VarsItem(222))
}

func TestVarsTwoTypeValues(t *testing.T) {
	s := NewVars()
	s.Set(0, 111)
	s.Set(1, "222")
	v, _ := s.Get(0)
	assert.Equal(t, v, VarsItem(111))
	v, _ = s.Get(1)
	assert.Equal(t, v, VarsItem("222"))
}

func TestVarsWitInvalidIndex(t *testing.T) {
	s := NewVars()
	s.Set(0, 111)
	v, err := s.Get(1)
	assert.EqualError(t, err, "Variable index undefined")
	assert.Equal(t, VarsItem(0), v)
}

func TestVarsTwoTypeValuesWithAdd(t *testing.T) {
	s := NewVars()
	v1index := s.Add()
	s.Set(v1index, 111)
	v2index := s.Add()
	s.Set(v2index, "222")
	v, _ := s.Get(v1index)
	assert.Equal(t, v, VarsItem(111))
	v, _ = s.Get(v2index)
	assert.Equal(t, v, VarsItem("222"))
}
