package bytecode

import (
	"errors"
)

// Bytecode is responsible for keep the instructions.
type Bytecode struct {
	instructions []int
}

// New creates a new bytecode.
func New() *Bytecode {
	return &Bytecode{
		instructions: make([]int, 0),
	}
}

// Add adds a new opcode to the bytecode.
func (bc *Bytecode) Add(item int, op int) {
	if bc == nil {
		return
	}
	bc.instructions = append(bc.instructions, item)
	bc.instructions = append(bc.instructions, op)
}

// Get returns the opcode at a specific index.
func (bc *Bytecode) Get(index int) (int, error) {
	if bc == nil || index > len(bc.instructions)-1 {
		return 0, errors.New("Index not found")
	}

	res := bc.instructions[index]

	return res, nil
}

// Len returns the bytecode length.
func (bc *Bytecode) Len() int {
	if bc == nil {
		return 0
	}

	return len(bc.instructions)
}
