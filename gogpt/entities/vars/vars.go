package vars

import (
	"errors"
)

type VarsItem interface{}

type Vars struct {
	items map[int]VarsItem
}

func NewVars() *Vars {
	s := &Vars{}
	s.items = make(map[int]VarsItem)

	return s
}

func (s *Vars) Add() int {
	index := len(s.items)
	s.items[index] = nil

	return index
}

func (s *Vars) Set(index int, item VarsItem) {
	s.items[index] = item
}

func (s *Vars) Get(index int) (VarsItem, error) {
	value, ok := s.items[index]
	if !ok {
		return 0, errors.New("Variable index undefined")
	}

	return value, nil
}
