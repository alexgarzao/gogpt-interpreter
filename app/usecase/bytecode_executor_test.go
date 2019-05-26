package vm

import (
	"testing"

	"github.com/alexgarzao/gpt-interpreter/app/domain"
	"github.com/stretchr/testify/assert"
)

func TestBCERunningLic222(t *testing.T) {
	// CP map:
	//		0: 222
	cp := opcodes.NewCp()
	cp_index := cp.Add(222)
	st := opcodes.NewStack()
	bc := opcodes.NewBytecode()
	bc.Add(opcodes.Lic, opcodes.BytecodeItem(cp_index))
	bce := NewBytecodeExecutor()
	bce.Run(cp, st, bc)
	cpv, _ := cp.Get(0)
	assert.Equal(t, cpv, opcodes.CPItem(222))
	stv, _ := st.Top()
	assert.Equal(t, stv, opcodes.StackItem(222))
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
