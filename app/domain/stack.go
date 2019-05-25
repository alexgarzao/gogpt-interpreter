package opcodes

import (
	"errors"
)

type StackItem int

type Stack struct {
	items []StackItem
}

func NewStack() *Stack {
	s := &Stack{}
	s.items = make([]StackItem, 0)
	return s
}

func (s *Stack) Push(item StackItem) {
	s.items = append(s.items, item)
}

func (s *Stack) Pop() (StackItem, error) {
	l := len(s.items)
	if l == 0 {
		return 0, errors.New("Stack underflow")
	}

	res := s.items[l-1]
	s.items = s.items[:l-1]

	return res, nil
}

func (s *Stack) Top() (StackItem, error) {
	l := len(s.items)
	if l == 0 {
		return 0, errors.New("Stack underflow")
	}

	res := s.items[l-1]

	return res, nil
}
