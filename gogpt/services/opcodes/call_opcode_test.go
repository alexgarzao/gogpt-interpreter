package opcodes

import (
	"testing"

	"github.com/alexgarzao/gogpt-interpreter/gogpt/infrastructure"
	"github.com/alexgarzao/gogpt-interpreter/gogpt/model/constant_pool"
	"github.com/alexgarzao/gogpt-interpreter/gogpt/model/stack"
	"github.com/alexgarzao/gogpt-interpreter/gogpt/model/vars"
	"github.com/stretchr/testify/assert"
)

func TestCallStringHello(t *testing.T) {
	// CP:
	//    0: STR "io.println"
	//    1: STR "Hello World!"
	//    2: INT 1

	// CODE:
	//    LDC 1 (Hello World!)
	//    LDC 2 (1)
	//    CALL 0 (io.println)
	cp := constant_pool.NewCP()
	printlnIndex := cp.Add("io.println")
	messageIndex := cp.Add("Hello World!")
	argsCountIndex := cp.Add(1)
	vars := vars.NewVars()
	st := stack.NewStack()
	stdin := infrastructure.NewFakeStdin()
	stdout := infrastructure.NewFakeStdout()

	ldc := NewLDCOpcode()
	ldc.CpIndex = messageIndex
	ldc.Execute(cp, vars, st, stdin, stdout)
	ldc = NewLDCOpcode()
	ldc.CpIndex = argsCountIndex
	ldc.Execute(cp, vars, st, stdin, stdout)
	call := NewCALLOpcode()
	call.CpIndex = printlnIndex
	call.Execute(cp, vars, st, stdin, stdout)
	assert.Equal(t, stdout.LastLine, "Hello World!")
}
