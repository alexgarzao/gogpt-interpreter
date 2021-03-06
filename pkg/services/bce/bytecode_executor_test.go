package bce

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/alexgarzao/gogpt-interpreter/pkg/infrastructure"
	"github.com/alexgarzao/gogpt-interpreter/pkg/model/bytecode"
	"github.com/alexgarzao/gogpt-interpreter/pkg/model/cp"
	"github.com/alexgarzao/gogpt-interpreter/pkg/model/stack"
	"github.com/alexgarzao/gogpt-interpreter/pkg/model/vars"
	"github.com/alexgarzao/gogpt-interpreter/pkg/services/instructions"
	"github.com/alexgarzao/gogpt-interpreter/pkg/services/lexer"
	"github.com/alexgarzao/gogpt-interpreter/pkg/services/parser"
)

func TestBCEAddingAndFetchingBytecodes(t *testing.T) {
	bc := bytecode.New()
	bc.Add(instructions.LDC, 111)
	bc.Add(instructions.LDC, 222)

	assert.Equal(t, 4, bc.Len())

	bce := New(bc)

	v, _ := bce.Next()
	assert.Equal(t, instructions.LDC, v)
	v, _ = bce.Next()
	assert.Equal(t, 111, v)

	v, _ = bce.Next()
	assert.Equal(t, instructions.LDC, v)
	v, _ = bce.Next()
	assert.Equal(t, 222, v)
}

func TestBCERunningLDC222(t *testing.T) {
	// CP map:
	//		0: 222
	cp := cp.New()
	cpIndex := cp.Add(222)
	vars := vars.New()
	st := stack.New()
	stdin := infrastructure.NewFakeStdin()
	stdout := infrastructure.NewFakeStdout()
	bc := bytecode.New()
	bc.Add(instructions.LDC, cpIndex)
	bce := New(bc)
	err := bce.Run(cp, vars, st, stdin, stdout)
	assert.Nil(t, err)
	cpv, _ := cp.Get(0)
	assert.Equal(t, 222, cpv)
	stv, _ := st.Top()
	assert.Equal(t, 222, stv)
}

func TestBCERunningNOP(t *testing.T) {
	cp := cp.New()
	vars := vars.New()
	st := stack.New()
	stdin := infrastructure.NewFakeStdin()
	stdout := infrastructure.NewFakeStdout()
	bc := bytecode.New()
	bc.Add(instructions.NOP, 0)
	bce := New(bc)
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
	cp := cp.New()
	printlnIndex := cp.Add("io.println")
	messageIndex := cp.Add("Hello World!")
	argsCountIndex := cp.Add(1)

	vars := vars.New()
	st := stack.New()
	stdin := infrastructure.NewFakeStdin()
	stdout := infrastructure.NewFakeStdout()
	bc := bytecode.New()
	bc.Add(instructions.LDC, messageIndex)
	bc.Add(instructions.LDC, argsCountIndex)
	bc.Add(instructions.CALL, printlnIndex)
	bce := New(bc)
	err := bce.Run(cp, vars, st, stdin, stdout)
	assert.Nil(t, err)
	assert.Equal(t, "Hello World!", stdout.LastLine)
}

func TestBCERunningInvalidOpcode(t *testing.T) {
	cp := cp.New()
	vars := vars.New()
	st := stack.New()
	stdin := infrastructure.NewFakeStdin()
	stdout := infrastructure.NewFakeStdout()
	bc := bytecode.New()
	bc.Add(123, 0)
	bce := New(bc)
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

	cp := cp.New()
	printlnIndex := cp.Add("io.println")
	messageIndex1 := cp.Add("Qual o seu nome?")
	readlnIndex := cp.Add("io.readln")
	messageIndex2 := cp.Add("Olá ")
	argsCountIndex := cp.Add(1)

	bc := bytecode.New()
	bc.Add(instructions.LDC, messageIndex1)
	bc.Add(instructions.LDC, argsCountIndex)
	bc.Add(instructions.CALL, printlnIndex)
	bc.Add(instructions.CALL, readlnIndex)
	bc.Add(instructions.STV, 0)
	bc.Add(instructions.LDC, messageIndex2)
	bc.Add(instructions.LDC, argsCountIndex)
	bc.Add(instructions.CALL, printlnIndex)
	bc.Add(instructions.LDV, 0)
	bc.Add(instructions.LDC, argsCountIndex)
	bc.Add(instructions.CALL, printlnIndex)

	vars := vars.New()
	st := stack.New()
	stdin := infrastructure.NewFakeStdin()
	stdout := infrastructure.NewFakeStdout()
	stdin.NextLineToRead("aaa123")

	bce := New(bc)
	err := bce.Run(cp, vars, st, stdin, stdout)

	assert.Nil(t, err)
	assert.Equal(t, "aaa123", stdout.LastLine)
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

	l := lexer.New(a)
	p := parser.New(l)

	pr := p.Parser()
	assert.True(t, pr.Parsed)
	bce := New(p.GetBC())
	stdin := infrastructure.NewFakeStdin()
	stdout := infrastructure.NewFakeStdout()
	st := stack.New()
	vars := vars.New()

	err := bce.Run(p.GetCP(), vars, st, stdin, stdout)
	assert.Nil(t, err)

	assert.Equal(t, "99", stdout.LastLine)
}
