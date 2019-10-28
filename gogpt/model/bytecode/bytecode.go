package bytecode

import (
	"errors"
)

// Bytecode is responsible for keep the opcodes.
type Bytecode struct {
	opcodes []int
}

// NewBytecode creates a new bytecode.
func NewBytecode() *Bytecode {
	return &Bytecode{
		opcodes: make([]int, 0),
	}
}

// Add adds a new opcode to the bytecode.
func (bc *Bytecode) Add(item int, op int) {
	if bc == nil {
		return
	}
	bc.opcodes = append(bc.opcodes, item)
	bc.opcodes = append(bc.opcodes, op)
}

// Get returns the opcode at a specific index.
func (bc *Bytecode) Get(index int) (int, error) {
	if bc == nil || index > len(bc.opcodes)-1 {
		return 0, errors.New("Index not found")
	}

	res := bc.opcodes[index]

	return res, nil
}

// Len returns the bytecode length.
func (bc *Bytecode) Len() int {
	if bc == nil {
		return 0
	}

	return len(bc.opcodes)
}
