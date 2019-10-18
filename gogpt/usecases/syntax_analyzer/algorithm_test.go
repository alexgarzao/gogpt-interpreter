package syntax

import (
	"testing"

	"github.com/alexgarzao/gogpt-interpreter/gogpt/entities/bytecode"
	"github.com/alexgarzao/gogpt-interpreter/gogpt/entities/constant_pool"

	"github.com/stretchr/testify/assert"

	lexer "github.com/alexgarzao/gogpt-interpreter/gogpt/entities/lexical_analyzer"
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
