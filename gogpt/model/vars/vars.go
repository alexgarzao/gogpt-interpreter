package vars

import (
	"errors"
)

// VarsItem has the type of a var item.
type VarsItem interface{}

// Vars has the variables of an algorithm.
type Vars struct {
	vars map[int]VarsItem
}

// NewVars creates a new vars.
func NewVars() *Vars {
	s := &Vars{}
	s.vars = make(map[int]VarsItem)

	return s
}

// Add adds a new var.
func (s *Vars) Add() int {
	index := len(s.vars)
	s.vars[index] = nil

	return index
}

// Set defines a var value.
func (s *Vars) Set(index int, value VarsItem) {
	s.vars[index] = value
}

// Get gets a var value.
func (s *Vars) Get(index int) (VarsItem, error) {
	value, ok := s.vars[index]
	if !ok {
		return 0, errors.New("Variable index undefined")
	}

	return value, nil
}
