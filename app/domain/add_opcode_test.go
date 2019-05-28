package opcodes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidAdd2And3(t *testing.T) {
	// 2 + 3 == 5
	// CP:
	// 		0: (INT) 2
	// 		1: (INT) 3
	// CODE:
	// 		LDC 0
	// 		LDC 1
	// 		ADD
	cp := NewCp()
	cpIndex2 := cp.Add(2)
	cpIndex3 := cp.Add(3)
	stack := NewStack()

	// LDC 0
	ldc := NewLdcOpcode()
	ldc.Execute(cp, stack, cpIndex2)
	ldc.Execute(cp, stack, cpIndex3)

	add := NewAddOpcode()
	add.Execute(cp, stack, 0)

	stv, _ := stack.Top()
	assert.Equal(t, stv, StackItem(5))
}
