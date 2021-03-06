package vars

import (
	"github.com/alexgarzao/gogpt-interpreter/pkg/domain"
)

// Vars has the variables of an algorithm.
type Vars struct {
	vars map[int]interface{}
}

// New creates a new vars.
func New() *Vars {
	v := &Vars{}
	v.vars = make(map[int]interface{})

	return v
}

// Add adds a new var.
func (v *Vars) Add() int {
	index := len(v.vars)
	v.vars[index] = nil

	return index
}

// Set defines a var value.
func (v *Vars) Set(index int, value interface{}) {
	v.vars[index] = value
}

// Get gets a var value.
func (v *Vars) Get(index int) (interface{}, error) {
	value, ok := v.vars[index]
	if !ok {
		return 0, domain.ErrUndefinedVarIndex
	}

	return value, nil
}
