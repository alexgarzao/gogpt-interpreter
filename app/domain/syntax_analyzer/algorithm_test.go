package analyzer

import (
	"testing"

	"github.com/stretchr/testify/assert"

	opcodes "github.com/alexgarzao/gpt-interpreter/app/domain"
	lexer "github.com/alexgarzao/gpt-interpreter/app/domain/lexical_analyzer"
)

func TestValidEmptyProgram(t *testing.T) {
	c :=
		`algoritmo olá_mundo;
início
fim`
	l := lexer.NewLexer(c)
	p := NewProgram()
	assert.Equal(t, true, p.TryToParse(l))
}

func TestValidHelloWorldAlgorithm(t *testing.T) {
	c :=
		`algoritmo olá_mundo;
início
	imprima("Olá mundo!");
fim`
	l := lexer.NewLexer(c)
	p := NewProgram()
	assert.Equal(t, true, p.TryToParse(l))
}

func TestValidHelloWorldWithTwoSentences(t *testing.T) {
	c :=
		`algoritmo olá_mundo;
início
	imprima("Olá...");
	imprima("Mundo!");
fim`
	l := lexer.NewLexer(c)
	p := NewProgram()
	assert.Equal(t, true, p.TryToParse(l))
}

func TestBytecodeEmptyProgram(t *testing.T) {
	c :=
		`algoritmo olá_mundo;
início
fim`
	l := lexer.NewLexer(c)
	p := NewProgram()
	bc := opcodes.NewBytecode()
	assert.Equal(t, true, p.TryToParse(l))
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

	expectedCp := opcodes.NewCp()
	printlnIndex := expectedCp.Add("io.println")

	l := lexer.NewLexer(c)
	p := NewProgram()
	expectedBc := opcodes.NewBytecode()
	expectedBc.Add(opcodes.Call, printlnIndex)

	assert.Equal(t, true, p.TryToParse(l))
	assert.Equal(t, expectedCp, p.GetCP())
	assert.Equal(t, expectedBc, p.GetBC())
}

func TestBytecodeHelloWorldProgram(t *testing.T) {
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

	expectedCp := opcodes.NewCp()
	printlnIndex := expectedCp.Add("io.println")
	messageIndex := expectedCp.Add("Olá mundo!")

	l := lexer.NewLexer(c)
	p := NewProgram()
	expectedBc := opcodes.NewBytecode()
	expectedBc.Add(opcodes.Ldc, messageIndex)
	expectedBc.Add(opcodes.Call, printlnIndex)

	assert.Equal(t, true, p.TryToParse(l))
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

	expectedCp := opcodes.NewCp()
	printlnIndex := expectedCp.Add("io.println")
	messageIndex1 := expectedCp.Add("Olá...")
	messageIndex2 := expectedCp.Add("mundo!")

	l := lexer.NewLexer(c)
	p := NewProgram()
	expectedBc := opcodes.NewBytecode()
	expectedBc.Add(opcodes.Ldc, messageIndex1)
	expectedBc.Add(opcodes.Call, printlnIndex)
	expectedBc.Add(opcodes.Ldc, messageIndex2)
	expectedBc.Add(opcodes.Call, printlnIndex)

	assert.Equal(t, true, p.TryToParse(l))
	assert.Equal(t, expectedCp, p.GetCP())
	assert.Equal(t, expectedBc, p.GetBC())
}
