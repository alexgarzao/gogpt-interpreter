package opcodes

import (
	"errors"
)

type CPItem int

type CP struct {
	items []CPItem
}

func NewCp() *CP {
	cp := &CP{}
	cp.items = make([]CPItem, 0)
	return cp
}

func (cp *CP) Add(item CPItem) int {
	cp.items = append(cp.items, item)
	return len(cp.items) - 1
}

func (cp *CP) Get(index int) (CPItem, error) {
	if index > len(cp.items)-1 {
		return 0, errors.New("Index not found")
	}

	res := cp.items[index]

	return res, nil
}
