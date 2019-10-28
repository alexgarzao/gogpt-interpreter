package opcodes

import (
	"testing"

	"github.com/alexgarzao/gogpt-interpreter/gogpt/infrastructure"
	"github.com/alexgarzao/gogpt-interpreter/gogpt/model/constant_pool"
	"github.com/alexgarzao/gogpt-interpreter/gogpt/model/stack"
	"github.com/alexgarzao/gogpt-interpreter/gogpt/model/vars"
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
	cp := constant_pool.NewCP()
	cpIndex2 := cp.Add(2)
	cpIndex3 := cp.Add(3)
	vars := vars.NewVars()
	st := stack.NewStack()
	stdin := infrastructure.NewFakeStdin()
	stdout := infrastructure.NewFakeStdout()

	// LDC 0
	ldc := NewLDCOpcode()
	ldc.CpIndex = cpIndex2
	ldc.Execute(cp, vars, st, stdin, stdout)
	ldc.CpIndex = cpIndex3
	ldc.Execute(cp, vars, st, stdin, stdout)

	add := NewADDOpcode()
	add.Execute(cp, st)

	stv, _ := st.Top()
	assert.Equal(t, stv, stack.StackItem(5))
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

	cp := constant_pool.NewCP()
	cpIndex2 := cp.Add("Hello")
	cpIndex3 := cp.Add("World")
	vars := vars.NewVars()
	st := stack.NewStack()
	stdin := infrastructure.NewFakeStdin()
	stdout := infrastructure.NewFakeStdout()

	// LDC 0
	ldc := NewLDCOpcode()
	ldc.CpIndex = cpIndex2
	ldc.Execute(cp, vars, st, stdin, stdout)
	ldc.CpIndex = cpIndex3
	ldc.Execute(cp, vars, st, stdin, stdout)

	add := NewADDOpcode()
	add.Execute(cp, st)

	stv, _ := st.Top()
	assert.Equal(t, stv, stack.StackItem("HelloWorld"))
}
