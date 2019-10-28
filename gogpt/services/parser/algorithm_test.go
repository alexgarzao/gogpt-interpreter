package parser

import (
	"testing"

	"github.com/alexgarzao/gogpt-interpreter/gogpt/model/bytecode"
	"github.com/alexgarzao/gogpt-interpreter/gogpt/model/constant_pool"
	"github.com/alexgarzao/gogpt-interpreter/gogpt/services/lexer"
	"github.com/alexgarzao/gogpt-interpreter/gogpt/services/opcodes"
	"github.com/stretchr/testify/assert"
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
	//    1: INT 0
	// CODE:
	//    LDC  1 (0)
	//    CALL 0 (io.println)

	expectedCp := constant_pool.NewCP()
	printlnIndex := expectedCp.Add("io.println")
	argsCountIndex := expectedCp.Add(0)

	l := lexer.NewLexer(c)
	p := NewAlgorithm(l)
	expectedBc := bytecode.NewBytecode()
	expectedBc.Add(opcodes.Ldc, argsCountIndex)
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
	//    2: INT 1
	// CODE:
	//    LDC 1 (Olá mundo!)
	//    LDC 2 (1)
	//    CALL 0 (io.println)

	expectedCp := constant_pool.NewCP()
	printlnIndex := expectedCp.Add("io.println")
	messageIndex := expectedCp.Add("Olá mundo!")
	argsCountIndex := expectedCp.Add(1)

	l := lexer.NewLexer(c)
	p := NewAlgorithm(l)
	expectedBc := bytecode.NewBytecode()
	expectedBc.Add(opcodes.Ldc, messageIndex)
	expectedBc.Add(opcodes.Ldc, argsCountIndex)
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
	//    2: INT 1
	//    3: STR "mundo!"
	// CODE:
	//    LDC 1 (Olá...)
	//    LDC 2 (1)
	//    CALL 0 (io.println)
	//    LDC 3 (mundo!)
	//    LDC 2 (1)
	//    CALL 0 (io.println)

	expectedCp := constant_pool.NewCP()
	printlnIndex := expectedCp.Add("io.println")
	messageIndex1 := expectedCp.Add("Olá...")
	argsCountIndex := expectedCp.Add(1)
	messageIndex2 := expectedCp.Add("mundo!")

	l := lexer.NewLexer(c)
	p := NewAlgorithm(l)
	expectedBc := bytecode.NewBytecode()
	expectedBc.Add(opcodes.Ldc, messageIndex1)
	expectedBc.Add(opcodes.Ldc, argsCountIndex)
	expectedBc.Add(opcodes.Call, printlnIndex)
	expectedBc.Add(opcodes.Ldc, messageIndex2)
	expectedBc.Add(opcodes.Ldc, argsCountIndex)
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
	//    2: INT 1
	//    3: STR "io.readln"
	//    4: STR "Olá "
	// VAR:
	//    0: STR "nome"
	// CODE:
	//    LDC 1 (Qual o seu nome?)
	//    LDC 2 (1)
	//    CALL 0 (io.println)
	//    CALL 3 (io.readln)
	//    STV 0 (nome)
	//    LDC 4 (Olá )
	//    LDC 2 (1)
	//    CALL 0 (io.println)
	//    LDV 0 (nome)
	//    LDC 3 (1)
	//    CALL 0 (io.println)

	expectedCp := constant_pool.NewCP()
	printlnIndex := expectedCp.Add("io.println")
	messageIndex1 := expectedCp.Add("Qual o seu nome?")
	argsCountIndex := expectedCp.Add(1)
	readlnIndex := expectedCp.Add("io.readln")
	messageIndex2 := expectedCp.Add("Olá ")

	l := lexer.NewLexer(c)
	p := NewAlgorithm(l)
	expectedBc := bytecode.NewBytecode()
	expectedBc.Add(opcodes.Ldc, messageIndex1)
	expectedBc.Add(opcodes.Ldc, argsCountIndex)
	expectedBc.Add(opcodes.Call, printlnIndex)
	expectedBc.Add(opcodes.Call, readlnIndex)
	expectedBc.Add(opcodes.Stv, 0)
	expectedBc.Add(opcodes.Ldc, messageIndex2)
	expectedBc.Add(opcodes.Ldc, argsCountIndex)
	expectedBc.Add(opcodes.Call, printlnIndex)
	expectedBc.Add(opcodes.Ldv, 0)
	expectedBc.Add(opcodes.Ldc, argsCountIndex)
	expectedBc.Add(opcodes.Call, printlnIndex)

	pr := p.Parser()
	assert.Equal(t, true, pr.Parsed)
	assert.Equal(t, expectedCp, p.GetCP())
	assert.Equal(t, expectedBc, p.GetBC())
}

func TestBytecodeHelloWorldWithTwoArgs(t *testing.T) {
	c :=
		`algoritmo olá_mundo;
início
	imprima("Olá...", "mundo!");
fim`
	// CP:
	//    0: STR "io.println"
	//    1: STR "Olá..."
	//    2: STR "mundo!"
	//    3: INT 2
	// CODE:
	//    LDC 1 (Olá...)
	//    LDC 2 (mundo!)
	//    LDC 3 (2)
	//    CALL 0 (io.println)

	expectedCp := constant_pool.NewCP()
	printlnIndex := expectedCp.Add("io.println")
	messageIndex1 := expectedCp.Add("Olá...")
	messageIndex2 := expectedCp.Add("mundo!")
	argsCountIndex := expectedCp.Add(2)

	l := lexer.NewLexer(c)
	p := NewAlgorithm(l)
	expectedBc := bytecode.NewBytecode()
	expectedBc.Add(opcodes.Ldc, messageIndex1)
	expectedBc.Add(opcodes.Ldc, messageIndex2)
	expectedBc.Add(opcodes.Ldc, argsCountIndex)
	expectedBc.Add(opcodes.Call, printlnIndex)

	pr := p.Parser()
	assert.Equal(t, true, pr.Parsed)
	assert.Equal(t, expectedCp, p.GetCP())
	assert.Equal(t, expectedBc, p.GetBC())
}
