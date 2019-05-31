package opcodes

import (
	"testing"
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
	st := NewStack()
	ldc := NewLdcOpcode()
	ldc.CpIndex = messageIndex
	ldc.Execute(cp, st)
	call := NewCallOpcode()
	call.CpIndex = printlnIndex
	call.Execute(cp, st)
	// TODO: check println output
}
