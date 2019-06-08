package vm

import (
	"testing"

	adapters "github.com/alexgarzao/gpt-interpreter/gpt/adapters"

	opcodes "github.com/alexgarzao/gpt-interpreter/gpt/entities"
	"github.com/stretchr/testify/assert"
)

func TestBCERunningLdc222(t *testing.T) {
	// CP map:
	//		0: 222
	cp := opcodes.NewCp()
	cpIndex := cp.Add(222)
	st := opcodes.NewStack()
	stdout := adapters.NewFakeStdout()
	bc := opcodes.NewBytecode()
	bc.Add(opcodes.Ldc, cpIndex)
	bce := NewBytecodeExecutor()
	err := bce.Run(cp, st, stdout, bc)
	assert.Nil(t, err)
	cpv, _ := cp.Get(0)
	assert.Equal(t, cpv, opcodes.CPItem(222))
	stv, _ := st.Top()
	assert.Equal(t, stv, opcodes.StackItem(222))
}

func TestBCERunningNop(t *testing.T) {
	cp := opcodes.NewCp()
	st := opcodes.NewStack()
	stdout := adapters.NewFakeStdout()
	bc := opcodes.NewBytecode()
	bc.Add(opcodes.Nop, 0)
	bce := NewBytecodeExecutor()
	err := bce.Run(cp, st, stdout, bc)
	assert.Nil(t, err)
	_, err = cp.Get(0)
	assert.EqualError(t, err, "Index not found")
	_, err = st.Top()
	assert.EqualError(t, err, "Stack underflow")
}

func TestBCECompleteHelloWorld(t *testing.T) {
	// CP:
	//    0: STR "io.println"
	//    1: STR "Hello World!"

	// CODE:
	//    LDC 1 (Hello World!)
	//    CALL 0 (io.println)
	cp := opcodes.NewCp()
	printlnIndex := cp.Add("io.println")
	messageIndex := cp.Add("Hello World!")
	st := opcodes.NewStack()
	stdout := adapters.NewFakeStdout()
	bc := opcodes.NewBytecode()
	bc.Add(opcodes.Ldc, messageIndex)
	bc.Add(opcodes.Call, printlnIndex)
	bce := NewBytecodeExecutor()
	err := bce.Run(cp, st, stdout, bc)
	assert.Nil(t, err)
	assert.Equal(t, stdout.LastLine, "Hello World!\n")
}

func TestBCERunningInvalidOpcode(t *testing.T) {
	cp := opcodes.NewCp()
	st := opcodes.NewStack()
	stdout := adapters.NewFakeStdout()
	bc := opcodes.NewBytecode()
	bc.Add(123, 0)
	bce := NewBytecodeExecutor()
	err := bce.Run(cp, st, stdout, bc)
	assert.EqualError(t, err, "Invalid opcode 123")
}
