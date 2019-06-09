package opcodes

import (
	"github.com/alexgarzao/gogpt-interpreter/gogpt/entities/constant_pool"
	"github.com/alexgarzao/gogpt-interpreter/gogpt/entities/stack"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/alexgarzao/gogpt-interpreter/gogpt/adapters"
)

func TestCallStringHello(t *testing.T) {
	// CP:
	//    0: STR "io.println"
	//    1: STR "Hello World!"

	// CODE:
	//    LDC 1 (Hello World!)
	//    CALL 0 (io.println)
	cp := constant_pool.NewCp()
	printlnIndex := cp.Add("io.println")
	messageIndex := cp.Add("Hello World!")
	st := stack.NewStack()
	stdout := adapters.NewFakeStdout()
	ldc := NewLdcOpcode()
	ldc.CpIndex = messageIndex
	ldc.Execute(cp, st, stdout)
	call := NewCallOpcode()
	call.CpIndex = printlnIndex
	call.Execute(cp, st, stdout)
	assert.Equal(t, stdout.LastLine, "Hello World!\n")
}
