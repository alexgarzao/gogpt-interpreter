package bce

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/alexgarzao/gogpt-interpreter/gogpt/infrastructure"
	"github.com/alexgarzao/gogpt-interpreter/gogpt/model/bytecode"
	"github.com/alexgarzao/gogpt-interpreter/gogpt/model/constant_pool"
	"github.com/alexgarzao/gogpt-interpreter/gogpt/model/stack"
	"github.com/alexgarzao/gogpt-interpreter/gogpt/model/vars"
	"github.com/alexgarzao/gogpt-interpreter/gogpt/services/lexer"
	"github.com/alexgarzao/gogpt-interpreter/gogpt/services/opcodes"
	"github.com/alexgarzao/gogpt-interpreter/gogpt/services/parser"
)

func TestBCEAddingAndFetchingBytecodes(t *testing.T) {
	bc := bytecode.NewBytecode()
	bc.Add(opcodes.Ldc, 111)
	bc.Add(opcodes.Ldc, 222)

	assert.Equal(t, bc.Len(), 4)

	bce := NewBytecodeExecutor(bc)

	v, _ := bce.next()
	assert.Equal(t, v, opcodes.Ldc)
	v, _ = bce.next()
	assert.Equal(t, v, 111)

	v, _ = bce.next()
	assert.Equal(t, v, opcodes.Ldc)
	v, _ = bce.next()
	assert.Equal(t, v, 222)
}

func TestBCERunningLdc222(t *testing.T) {
	// CP map:
	//		0: 222
	cp := constant_pool.NewCP()
	cpIndex := cp.Add(222)
	vars := vars.NewVars()
	st := stack.NewStack()
	stdin := infrastructure.NewFakeStdin()
	stdout := infrastructure.NewFakeStdout()
	bc := bytecode.NewBytecode()
	bc.Add(opcodes.Ldc, cpIndex)
	bce := NewBytecodeExecutor(bc)
	err := bce.Run(cp, vars, st, stdin, stdout)
	assert.Nil(t, err)
	cpv, _ := cp.Get(0)
	assert.Equal(t, cpv, constant_pool.CPItem(222))
	stv, _ := st.Top()
	assert.Equal(t, stv, stack.StackItem(222))
}

func TestBCERunningNop(t *testing.T) {
	cp := constant_pool.NewCP()
	vars := vars.NewVars()
	st := stack.NewStack()
	stdin := infrastructure.NewFakeStdin()
	stdout := infrastructure.NewFakeStdout()
	bc := bytecode.NewBytecode()
	bc.Add(opcodes.Nop, 0)
	bce := NewBytecodeExecutor(bc)
	err := bce.Run(cp, vars, st, stdin, stdout)
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
	bc := bytecode.NewBytecode()
	bc.Add(opcodes.Ldc, messageIndex)
	bc.Add(opcodes.Ldc, argsCountIndex)
	bc.Add(opcodes.Call, printlnIndex)
	bce := NewBytecodeExecutor(bc)
	err := bce.Run(cp, vars, st, stdin, stdout)
	assert.Nil(t, err)
	assert.Equal(t, stdout.LastLine, "Hello World!\n")
}

func TestBCERunningInvalidOpcode(t *testing.T) {
	cp := constant_pool.NewCP()
	vars := vars.NewVars()
	st := stack.NewStack()
	stdin := infrastructure.NewFakeStdin()
	stdout := infrastructure.NewFakeStdout()
	bc := bytecode.NewBytecode()
	bc.Add(123, 0)
	bce := NewBytecodeExecutor(bc)
	err := bce.Run(cp, vars, st, stdin, stdout)
	assert.EqualError(t, err, "Invalid opcode 123")
}

func TestBCEHelloWorldWithInput(t *testing.T) {
	// CP:
	//    0: STR "io.println"
	//    1: STR "Qual o seu nome?"
	//    2: STR "io.readln"
	//    3: STR "Olá "
	//    4: INT 1
	// VAR:
	//    0: STR "nome"
	// CODE:
	//    LDC 1 (Qual o seu nome?)
	//    LDC 4 (1)
	//    CALL 0 (io.println)
	//    CALL 2 (io.readln)
	//    STV 0 (nome)
	//    LDC 3 (Olá )
	//    LDC 4 (1)
	//    CALL 0 (io.println)
	//    LDV 0 (nome)
	//    LDC 4 (1)
	//    CALL 0 (io.println)

	cp := constant_pool.NewCP()
	printlnIndex := cp.Add("io.println")
	messageIndex1 := cp.Add("Qual o seu nome?")
	readlnIndex := cp.Add("io.readln")
	messageIndex2 := cp.Add("Olá ")
	argsCountIndex := cp.Add(1)

	bc := bytecode.NewBytecode()
	bc.Add(opcodes.Ldc, messageIndex1)
	bc.Add(opcodes.Ldc, argsCountIndex)
	bc.Add(opcodes.Call, printlnIndex)
	bc.Add(opcodes.Call, readlnIndex)
	bc.Add(opcodes.Stv, 0)
	bc.Add(opcodes.Ldc, messageIndex2)
	bc.Add(opcodes.Ldc, argsCountIndex)
	bc.Add(opcodes.Call, printlnIndex)
	bc.Add(opcodes.Ldv, 0)
	bc.Add(opcodes.Ldc, argsCountIndex)
	bc.Add(opcodes.Call, printlnIndex)

	vars := vars.NewVars()
	st := stack.NewStack()
	stdin := infrastructure.NewFakeStdin()
	stdout := infrastructure.NewFakeStdout()
	stdin.NextLineToRead("aaa123")

	bce := NewBytecodeExecutor(bc)
	err := bce.Run(cp, vars, st, stdin, stdout)

	assert.Nil(t, err)
	assert.Equal(t, "aaa123\n", stdout.LastLine)
}

func TestRunningWithTwoVars(t *testing.T) {
	a :=
		`algoritmo two_vars;

		variáveis
			nome: literal;
			idade: literal;
		fim-variáveis
		
		início
			nome := "name";
			idade := "99";
			imprima("Olá ");
			imprima(nome);
			imprima("Você tem a seguinte idade: ");
			imprima(idade);
		fim
		`

	l := lexer.NewLexer(a)
	p := parser.NewAlgorithm(l)

	pr := p.Parser()
	assert.True(t, pr.Parsed)
	bce := NewBytecodeExecutor(p.GetBC())
	stdin := infrastructure.NewFakeStdin()
	stdout := infrastructure.NewFakeStdout()
	st := stack.NewStack()
	vars := vars.NewVars()

	err := bce.Run(p.GetCP(), vars, st, stdin, stdout)
	assert.Nil(t, err)

	assert.Equal(t, "99\n", stdout.LastLine)
}
