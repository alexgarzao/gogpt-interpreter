package opcodes

import (
	"errors"
)

type CpItem int

type Cp struct {
	items []CpItem
}

func NewCp() *Cp {
	cp := &Cp{}
	cp.items = make([]CpItem, 0)
	return cp
}

func (cp *Cp) Add(item CpItem) int {
	cp.items = append(cp.items, item)
	return len(cp.items) - 1
}

func (cp *Cp) Get(index int) (CpItem, error) {
	if index > len(cp.items)-1 {
		return 0, errors.New("Index not found")
	}

	res := cp.items[index]

	return res, nil
}
