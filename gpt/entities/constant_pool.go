package opcodes

import (
	"errors"
)

type CPItem interface{}

type CP struct {
	items []CPItem
}

func NewCp() *CP {
	cp := &CP{}
	cp.items = make([]CPItem, 0)
	return cp
}

func (cp *CP) Add(item CPItem) int {
	if cp == nil {
		return -1
	}

	if n := cp.Find(item); n != -1 {
		return n
	}

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

func (cp *CP) Find(item CPItem) int {
	for i, v := range cp.items {
		if v == item {
			return i
		}
	}

	return -1
}
