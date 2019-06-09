package main

import (
	"errors"
	"github.com/alexgarzao/gpt-interpreter/gpt/entities/stack"
	"io/ioutil"
	"log"
	"os"

	"github.com/alexgarzao/gpt-interpreter/gpt/adapters"
	"github.com/alexgarzao/gpt-interpreter/gpt/entities/lexical_analyzer"
	"github.com/alexgarzao/gpt-interpreter/gpt/usecases/bytecode_executor"
	"github.com/alexgarzao/gpt-interpreter/gpt/usecases/syntax_analyzer"
)

func main() {
	filename, err := getFilenameFromArgs()
	if err != nil {
		log.Fatalln(err)
	}

	algorithm, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalln(err)
	}

	l := lexer.NewLexer(string(algorithm))
	p := syntax.NewProgram()

	if p.TryToParse(l) == false {
		log.Fatalln("Error in parsing")
		return
	}

	bce := bce.NewBytecodeExecutor()
	stdout := adapters.NewStdout()
	st := stack.NewStack()

	err = bce.Run(p.GetCP(), st, stdout, p.GetBC())
	if err != nil {
		log.Fatalf("Error %s\n", err)
	}
}

func getFilenameFromArgs() (string, error) {
	args := os.Args
	if len(args) < 2 {
		return "", errors.New("Usage: gpt <file>")
	}
	return os.Args[1], nil
}
