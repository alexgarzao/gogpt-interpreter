package syntax

import (
	"testing"

	"github.com/alexgarzao/gogpt-interpreter/gogpt/adapters"
	"github.com/alexgarzao/gogpt-interpreter/gogpt/entities/bytecode"
	"github.com/alexgarzao/gogpt-interpreter/gogpt/entities/constant_pool"
	"github.com/alexgarzao/gogpt-interpreter/gogpt/entities/stack"
	"github.com/alexgarzao/gogpt-interpreter/gogpt/entities/vars"

	"github.com/stretchr/testify/assert"

	lexer "github.com/alexgarzao/gogpt-interpreter/gogpt/entities/lexical_analyzer"
	bce "github.com/alexgarzao/gogpt-interpreter/gogpt/usecases/bytecode_executor"
	"github.com/alexgarzao/gogpt-interpreter/gogpt/usecases/opcodes"
)

func TestValidEmptyAlgorithm(t *testing.T) {
	c :=
		`algoritmo olá_mundo;
início
fim`
	l := lexer.NewLexer(c)
	p := NewAlgorithm(l)
	pr := p.Parser()
	assert.Equal(t, true, pr.Parsed)
}

func TestValidHelloWorldAlgorithm(t *testing.T) {
	c :=
		`algoritmo olá_mundo;
início
	imprima("Olá mundo!");
fim`
	l := lexer.NewLexer(c)
	p := NewAlgorithm(l)
	pr := p.Parser()
	assert.Equal(t, true, pr.Parsed)
}

func TestValidHelloWorldWithTwoSentences(t *testing.T) {
	c :=
		`algoritmo olá_mundo;
início
	imprima("Olá...");
	imprima("Mundo!");
fim`
	l := lexer.NewLexer(c)
	p := NewAlgorithm(l)
	pr := p.Parser()
	assert.Equal(t, true, pr.Parsed)
}

func TestBytecodeEmptyAlgorithm(t *testing.T) {
	c :=
		`algoritmo olá_mundo;
início
fim`
	l := lexer.NewLexer(c)
	p := NewAlgorithm(l)
	bc := bytecode.NewBytecode()
	pr := p.Parser()
	assert.Equal(t, true, pr.Parsed)
	assert.Equal(t, bc, p.GetBC())
}

func TestBytecodeFunctionCallWithoutArguments(t *testing.T) {
	c :=
		`algoritmo olá_mundo;
início
	imprima();
fim`
	// CP:
	//    0: STR "io.println"
	// CODE:
	//    CALL 0 (io.println)

	expectedCp := constant_pool.NewCp()
	printlnIndex := expectedCp.Add("io.println")

	l := lexer.NewLexer(c)
	p := NewAlgorithm(l)
	expectedBc := bytecode.NewBytecode()
	expectedBc.Add(opcodes.Call, printlnIndex)

	pr := p.Parser()
	assert.Equal(t, true, pr.Parsed)
	assert.Equal(t, expectedCp, p.GetCP())
	assert.Equal(t, expectedBc, p.GetBC())
}

func TestBytecodeHelloWorldAlgorithm(t *testing.T) {
	c :=
		`algoritmo olá_mundo;
início
	imprima("Olá mundo!");
fim`
	// CP:
	//    0: STR "io.println"
	//    1: STR "Olá mundo!"
	// CODE:
	//    LDC 1 (Olá mundo!)
	//    CALL 0 (io.println)

	expectedCp := constant_pool.NewCp()
	printlnIndex := expectedCp.Add("io.println")
	messageIndex := expectedCp.Add("Olá mundo!")

	l := lexer.NewLexer(c)
	p := NewAlgorithm(l)
	expectedBc := bytecode.NewBytecode()
	expectedBc.Add(opcodes.Ldc, messageIndex)
	expectedBc.Add(opcodes.Call, printlnIndex)

	pr := p.Parser()
	assert.Equal(t, true, pr.Parsed)
	assert.Equal(t, expectedCp, p.GetCP())
	assert.Equal(t, expectedBc, p.GetBC())
}

func TestBytecodeHelloWorldWithTwoWrites(t *testing.T) {
	c :=
		`algoritmo olá_mundo;
início
	imprima("Olá...");
	imprima("mundo!");
fim`
	// CP:
	//    0: STR "io.println"
	//    1: STR "Olá..."
	//    2: STR "mundo!"
	// CODE:
	//    LDC 1 (Olá...)
	//    CALL 0 (io.println)
	//    LDC 2 (mundo!)
	//    CALL 0 (io.println)

	expectedCp := constant_pool.NewCp()
	printlnIndex := expectedCp.Add("io.println")
	messageIndex1 := expectedCp.Add("Olá...")
	messageIndex2 := expectedCp.Add("mundo!")

	l := lexer.NewLexer(c)
	p := NewAlgorithm(l)
	expectedBc := bytecode.NewBytecode()
	expectedBc.Add(opcodes.Ldc, messageIndex1)
	expectedBc.Add(opcodes.Call, printlnIndex)
	expectedBc.Add(opcodes.Ldc, messageIndex2)
	expectedBc.Add(opcodes.Call, printlnIndex)

	pr := p.Parser()
	assert.Equal(t, true, pr.Parsed)
	assert.Equal(t, expectedCp, p.GetCP())
	assert.Equal(t, expectedBc, p.GetBC())
}

func TestInvalidCompleteAlgorithmDeclarationWithoutId(t *testing.T) {
	c :=
		`algoritmo ;
início
	imprima("Olá...");
fim`
	l := lexer.NewLexer(c)
	p := NewAlgorithm(l)

	pr := p.Parser()
	assert.Equal(t, false, pr.Parsed)
	assert.EqualError(t, pr.Err, "Expected IDENT")
}

func TestInvalidCompleteAlgorithmDeclarationWithoutSemicolon(t *testing.T) {
	c :=
		`algoritmo ola
início
	imprima("Olá...");
fim`
	l := lexer.NewLexer(c)
	p := NewAlgorithm(l)

	pr := p.Parser()
	assert.Equal(t, false, pr.Parsed)
	assert.EqualError(t, pr.Err, "Expected SEMICOLON")
}

func TestBytecodeHelloWorldWithInput(t *testing.T) {
	c :=
		`algoritmo qual_o_seu_nome;

		variáveis
			nome: literal;
		fim-variáveis
		
		início
			imprima("Qual o seu nome?");
			nome := leia();
			imprima("Olá ");
			imprima(nome);
		fim
		`
	// CP:
	//    0: STR "io.println"
	//    1: STR "Qual o seu nome?"
	//    2: STR "io.readln"
	//    3: STR "Olá "
	// VAR:
	//    0: STR "nome"
	// CODE:
	//    LDC 1 (Qual o seu nome?)
	//    CALL 0 (io.println)
	//    CALL 2 (io.readln)
	//    STV 0 (nome)
	//    LDC 3 (Olá )
	//    CALL 0 (io.println)
	//    LDV 0 (nome)
	//    CALL 0 (io.println)

	expectedCp := constant_pool.NewCp()
	printlnIndex := expectedCp.Add("io.println")
	messageIndex1 := expectedCp.Add("Qual o seu nome?")
	readlnIndex := expectedCp.Add("io.readln")
	messageIndex2 := expectedCp.Add("Olá ")

	l := lexer.NewLexer(c)
	p := NewAlgorithm(l)
	expectedBc := bytecode.NewBytecode()
	expectedBc.Add(opcodes.Ldc, messageIndex1)
	expectedBc.Add(opcodes.Call, printlnIndex)
	expectedBc.Add(opcodes.Call, readlnIndex)
	expectedBc.Add(opcodes.Stv, 0)
	expectedBc.Add(opcodes.Ldc, messageIndex2)
	expectedBc.Add(opcodes.Call, printlnIndex)
	expectedBc.Add(opcodes.Ldv, 0)
	expectedBc.Add(opcodes.Call, printlnIndex)

	pr := p.Parser()
	assert.Equal(t, true, pr.Parsed)
	assert.Equal(t, expectedCp, p.GetCP())
	assert.Equal(t, expectedBc, p.GetBC())
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
	p := NewAlgorithm(l)

	pr := p.Parser()
	assert.True(t, pr.Parsed)
	bce := bce.NewBytecodeExecutor(p.GetBC())
	stdout := adapters.NewFakeStdout()
	st := stack.NewStack()
	vars := vars.NewVars()

	err := bce.Run(p.GetCP(), vars, st, stdout)
	assert.Nil(t, err)

	assert.Equal(t, "99\n", stdout.LastLine)
}
