package bce

import (
	"github.com/stretchr/testify/assert"
	"testing"

	"github.com/alexgarzao/gogpt-interpreter/gogpt/adapters"
	"github.com/alexgarzao/gogpt-interpreter/gogpt/entities/bytecode"
	"github.com/alexgarzao/gogpt-interpreter/gogpt/entities/constant_pool"
	"github.com/alexgarzao/gogpt-interpreter/gogpt/entities/stack"
	"github.com/alexgarzao/gogpt-interpreter/gogpt/usecases/opcodes"
)

func TestBCEAddingAndFetchingBytecodes(t *testing.T) {
	bc := bytecode.NewBytecode()
	bc.Add(opcodes.Ldc, 111)
	bc.Add(opcodes.Ldc, 222)

	assert.Equal(t, bc.Len(), 4)

	bce := NewBytecodeExecutor(bc)

	v, _ := bce.Next()
	assert.Equal(t, v, opcodes.Ldc)
	v, _ = bce.Next()
	assert.Equal(t, v, 111)

	v, _ = bce.Next()
	assert.Equal(t, v, opcodes.Ldc)
	v, _ = bce.Next()
	assert.Equal(t, v, 222)
}

func TestBCERunningLdc222(t *testing.T) {
	// CP map:
	//		0: 222
	cp := constant_pool.NewCp()
	cpIndex := cp.Add(222)
	st := stack.NewStack()
	stdout := adapters.NewFakeStdout()
	bc := bytecode.NewBytecode()
	bc.Add(opcodes.Ldc, cpIndex)
	bce := NewBytecodeExecutor(bc)
	err := bce.Run(cp, st, stdout)
	assert.Nil(t, err)
	cpv, _ := cp.Get(0)
	assert.Equal(t, cpv, constant_pool.CPItem(222))
	stv, _ := st.Top()
	assert.Equal(t, stv, stack.StackItem(222))
}

func TestBCERunningNop(t *testing.T) {
	cp := constant_pool.NewCp()
	st := stack.NewStack()
	stdout := adapters.NewFakeStdout()
	bc := bytecode.NewBytecode()
	bc.Add(opcodes.Nop, 0)
	bce := NewBytecodeExecutor(bc)
	err := bce.Run(cp, st, stdout)
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
	cp := constant_pool.NewCp()
	printlnIndex := cp.Add("io.println")
	messageIndex := cp.Add("Hello World!")
	st := stack.NewStack()
	stdout := adapters.NewFakeStdout()
	bc := bytecode.NewBytecode()
	bc.Add(opcodes.Ldc, messageIndex)
	bc.Add(opcodes.Call, printlnIndex)
	bce := NewBytecodeExecutor(bc)
	err := bce.Run(cp, st, stdout)
	assert.Nil(t, err)
	assert.Equal(t, stdout.LastLine, "Hello World!\n")
}

func TestBCERunningInvalidOpcode(t *testing.T) {
	cp := constant_pool.NewCp()
	st := stack.NewStack()
	stdout := adapters.NewFakeStdout()
	bc := bytecode.NewBytecode()
	bc.Add(123, 0)
	bce := NewBytecodeExecutor(bc)
	err := bce.Run(cp, st, stdout)
	assert.EqualError(t, err, "Invalid opcode 123")
}
