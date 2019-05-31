package opcodes

import (
	"testing"

	interfaces "github.com/alexgarzao/gpt-interpreter/app/interface"
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
	stdout := interfaces.NewFakeStdout()

	// LDC 0
	ldc := NewLdcOpcode()
	ldc.CpIndex = cpIndex2
	ldc.Execute(cp, stack, stdout)
	ldc.CpIndex = cpIndex3
	ldc.Execute(cp, stack, stdout)

	add := NewAddOpcode()
	add.Execute(cp, stack)

	stv, _ := stack.Top()
	assert.Equal(t, stv, StackItem(5))
}

func TestValidAddHelloAndWorld(t *testing.T) {
	// "Hello" + "World" == "HelloWorld"
	// CP:
	// 		0: (STR) "Hello"
	// 		1: (STR) "World"
	// CODE:
	// 		LDC 0
	// 		LDC 1
	// 		ADD

	cp := NewCp()
	cpIndex2 := cp.Add("Hello")
	cpIndex3 := cp.Add("World")
	stack := NewStack()
	stdout := interfaces.NewFakeStdout()

	// LDC 0
	ldc := NewLdcOpcode()
	ldc.CpIndex = cpIndex2
	ldc.Execute(cp, stack, stdout)
	ldc.CpIndex = cpIndex3
	ldc.Execute(cp, stack, stdout)

	add := NewAddOpcode()
	add.Execute(cp, stack)

	stv, _ := stack.Top()
	assert.Equal(t, stv, StackItem("HelloWorld"))
}
