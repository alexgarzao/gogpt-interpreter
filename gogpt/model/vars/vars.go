package vars

import (
	"errors"
)

// Vars has the variables of an algorithm.
type Vars struct {
	vars map[int]interface{}
}

// New creates a new vars.
func New() *Vars {
	s := &Vars{}
	s.vars = make(map[int]interface{})

	return s
}

// Add adds a new var.
func (s *Vars) Add() int {
	index := len(s.vars)
	s.vars[index] = nil

	return index
}

// Set defines a var value.
func (s *Vars) Set(index int, value interface{}) {
	s.vars[index] = value
}

// Get gets a var value.
func (s *Vars) Get(index int) (interface{}, error) {
	value, ok := s.vars[index]
	if !ok {
		return 0, errors.New("Variable index undefined")
	}

	return value, nil
}
