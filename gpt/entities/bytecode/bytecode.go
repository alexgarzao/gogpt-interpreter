package bytecode

import (
	"errors"
)

type Bytecode struct {
	IP    int
	items []int
}

func NewBytecode() *Bytecode {
	return &Bytecode{
		IP:    0,
		items: make([]int, 0),
	}
}

func (bc *Bytecode) Add(item int, op int) {
	if bc == nil {
		return
	}
	bc.items = append(bc.items, item)
	bc.items = append(bc.items, op)
}

func (bc *Bytecode) Get(index int) (int, error) {
	if index > len(bc.items)-1 {
		return 0, errors.New("Index not found")
	}

	res := bc.items[index]

	return res, nil
}

func (bc *Bytecode) Len() int {
	return len(bc.items)
}

func (bc *Bytecode) Next() (code int, err error) {
	code, err = bc.Get(bc.IP)
	bc.IP += 1
	return
}
