package opcodes

import (
	"errors"
)

type Item int

type Stack struct {
	items []Item
}

func NewStack() *Stack {
	s := &Stack{}
	s.items = make([]Item, 0)
	return s
}

func (s *Stack) Push(item Item) {
	s.items = append(s.items, item)
}

func (s *Stack) Pop() (Item, error) {
	l := len(s.items)
	if l == 0 {
		return 0, errors.New("Stack underflow")
	}

	res := s.items[l-1]
	s.items = s.items[:l-1]

	return res, nil
}

func (s *Stack) Top() (Item, error) {
	l := len(s.items)
	if l == 0 {
		return 0, errors.New("Stack underflow")
	}

	res := s.items[l-1]

	return res, nil
}
