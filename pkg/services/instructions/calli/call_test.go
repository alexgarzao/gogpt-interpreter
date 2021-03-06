package calli

import (
	"testing"

	"github.com/alexgarzao/gogpt-interpreter/pkg/infrastructure"
	"github.com/alexgarzao/gogpt-interpreter/pkg/model/cp"
	"github.com/alexgarzao/gogpt-interpreter/pkg/model/stack"
	"github.com/alexgarzao/gogpt-interpreter/pkg/model/vars"
	"github.com/alexgarzao/gogpt-interpreter/pkg/services/instructions/ldci"
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
	cp := cp.New()
	printlnIndex := cp.Add("io.println")
	messageIndex := cp.Add("Hello World!")
	argsCountIndex := cp.Add(1)
	vars := vars.New()
	st := stack.New()
	stdin := infrastructure.NewFakeStdin()
	stdout := infrastructure.NewFakeStdout()

	ldc := ldci.New()
	ldc.CpIndex = messageIndex
	ldc.Execute(cp, vars, st, stdin, stdout)
	ldc = ldci.New()
	ldc.CpIndex = argsCountIndex
	ldc.Execute(cp, vars, st, stdin, stdout)
	call := New()
	call.CpIndex = printlnIndex
	call.Execute(cp, vars, st, stdin, stdout)
	assert.Equal(t, "Hello World!", stdout.LastLine)
}
