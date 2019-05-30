package opcodes

import (
	"errors"
)

type BytecodeItem byte

type Bytecode struct {
	IP    int
	items []BytecodeItem
}

func NewBytecode() *Bytecode {
	return &Bytecode{
		IP:    0,
		items: make([]BytecodeItem, 0),
	}
}

func (bc *Bytecode) Add(item Opcode, op BytecodeItem) {
	bc.items = append(bc.items, BytecodeItem(item))
	bc.items = append(bc.items, op)
}

func (bc *Bytecode) Get(index int) (BytecodeItem, error) {
	if index > len(bc.items)-1 {
		return 0, errors.New("Index not found")
	}

	res := bc.items[index]

	return res, nil
}

func (bc *Bytecode) Len() int {
	return len(bc.items)
}

func (bc *Bytecode) Next() (code BytecodeItem, err error) {
	code, err = bc.Get(bc.IP)
	bc.IP += 1
	return
}
