package vm

import (
	"testing"

	"github.com/alexgarzao/gpt-interpreter/app/domain"
	"github.com/stretchr/testify/assert"
)

func TestBCERunningLc1(t *testing.T) {
	cp := opcodes.NewCp()
	st := opcodes.NewStack()
	bc := opcodes.NewBytecode()
	bc.Add(opcodes.Aic, 1)
	bce := NewBytecodeExecutor()
	bce.Run(cp, st, bc)
	cpv, _ := cp.Get(0)
	assert.Equal(t, cpv, opcodes.CPItem(1))
}
