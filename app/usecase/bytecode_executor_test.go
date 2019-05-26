package vm

import (
	"testing"

	"github.com/alexgarzao/gpt-interpreter/app/domain"
	"github.com/stretchr/testify/assert"
)

func TestBCERunningAic1(t *testing.T) {
	cp := opcodes.NewCp()
	st := opcodes.NewStack()
	bc := opcodes.NewBytecode()
	bc.Add(opcodes.Aic, 1)
	bce := NewBytecodeExecutor()
	bce.Run(cp, st, bc)
	cpv, _ := cp.Get(0)
	assert.Equal(t, cpv, opcodes.CPItem(1))
}

func TestBCERunningNop(t *testing.T) {
	cp := opcodes.NewCp()
	st := opcodes.NewStack()
	bc := opcodes.NewBytecode()
	bc.Add(opcodes.Nop, 0)
	bce := NewBytecodeExecutor()
	bce.Run(cp, st, bc)
	_, err := cp.Get(0)
	assert.EqualError(t, err, "Index not found")
	_, err = st.Top()
	assert.EqualError(t, err, "Stack underflow")
}
