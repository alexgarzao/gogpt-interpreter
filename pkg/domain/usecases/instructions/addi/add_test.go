package addi

import (
	"testing"

	"github.com/alexgarzao/gogpt-interpreter/pkg/domain/entities/cp"
	"github.com/alexgarzao/gogpt-interpreter/pkg/domain/entities/stack"
	"github.com/alexgarzao/gogpt-interpreter/pkg/domain/entities/vars"
	"github.com/alexgarzao/gogpt-interpreter/pkg/domain/usecases/instructions/ldci"
	"github.com/alexgarzao/gogpt-interpreter/pkg/infrastructure"
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
	cp := cp.New()
	cpIndex2 := cp.Add(2)
	cpIndex3 := cp.Add(3)
	vars := vars.New()
	st := stack.New()
	stdin := infrastructure.NewFakeStdin()
	stdout := infrastructure.NewFakeStdout()

	// LDC 0
	ldc := ldci.New()
	ldc.CpIndex = cpIndex2
	ldc.Execute(cp, vars, st, stdin, stdout)
	ldc.CpIndex = cpIndex3
	ldc.Execute(cp, vars, st, stdin, stdout)

	add := New()
	add.Execute(cp, vars, st, stdin, stdout)

	stv, _ := st.Top()
	assert.Equal(t, 5, stv)
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

	cp := cp.New()
	cpIndex2 := cp.Add("Hello")
	cpIndex3 := cp.Add("World")
	vars := vars.New()
	st := stack.New()
	stdin := infrastructure.NewFakeStdin()
	stdout := infrastructure.NewFakeStdout()

	// LDC 0
	ldc := ldci.New()
	ldc.CpIndex = cpIndex2
	ldc.Execute(cp, vars, st, stdin, stdout)
	ldc.CpIndex = cpIndex3
	ldc.Execute(cp, vars, st, stdin, stdout)

	add := New()
	add.Execute(cp, vars, st, stdin, stdout)

	stv, _ := st.Top()
	assert.Equal(t, "HelloWorld", stv)
}
