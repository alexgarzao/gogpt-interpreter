package stack

import (
	"github.com/alexgarzao/gogpt-interpreter/pkg/domain"
)

// Stack has the items of a stack.
type Stack struct {
	items []interface{}
}

// New creates a new stack.
func New() *Stack {
	s := &Stack{}
	s.items = make([]interface{}, 0)
	return s
}

// Push pushs a new item onto the stack.
func (s *Stack) Push(item interface{}) {
	s.items = append(s.items, item)
}

// Pop pops a item from the stack.
func (s *Stack) Pop() (interface{}, error) {
	l := len(s.items)
	if l == 0 {
		return 0, domain.ErrStackUnderflow
	}

	res := s.items[l-1]
	s.items = s.items[:l-1]

	return res, nil
}

// Top gets the top item on the stack.
func (s *Stack) Top() (interface{}, error) {
	l := len(s.items)
	if l == 0 {
		return 0, domain.ErrStackUnderflow
	}

	res := s.items[l-1]

	return res, nil
}

// Size returns the stack size.
func (s *Stack) Size() int {
	return len(s.items)
}
