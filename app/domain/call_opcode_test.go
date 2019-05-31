package opcodes

import (
	"testing"

	"github.com/stretchr/testify/assert"

	interfaces "github.com/alexgarzao/gpt-interpreter/app/interface"
)

func TestCallStringHello(t *testing.T) {
	// CP:
	//    0: STR "io.println"
	//    1: STR "Hello World!"

	// CODE:
	//    LDC 1 (Hello World!)
	//    CALL 0 (io.println)
	cp := NewCp()
	printlnIndex := cp.Add("io.println")
	messageIndex := cp.Add("Hello World!")
	stack := NewStack()
	stdout := interfaces.NewFakeStdout()
	ldc := NewLdcOpcode()
	ldc.CpIndex = messageIndex
	ldc.Execute(cp, stack, stdout)
	call := NewCallOpcode()
	call.CpIndex = printlnIndex
	call.Execute(cp, stack, stdout)
	assert.Equal(t, stdout.LastLine, "Hello World!\n")
}
