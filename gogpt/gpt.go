package main

import (
	"errors"
	"github.com/alexgarzao/gogpt-interpreter/gogpt/entities/stack"
	"io/ioutil"
	"log"
	"os"

	"github.com/alexgarzao/gogpt-interpreter/gogpt/adapters"
	"github.com/alexgarzao/gogpt-interpreter/gogpt/entities/lexical_analyzer"
	"github.com/alexgarzao/gogpt-interpreter/gogpt/usecases/bytecode_executor"
	"github.com/alexgarzao/gogpt-interpreter/gogpt/usecases/syntax_analyzer"
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

	bce := bce.NewBytecodeExecutor(p.GetBC())
	stdout := adapters.NewStdout()
	st := stack.NewStack()

	err = bce.Run(p.GetCP(), st, stdout)
	if err != nil {
		log.Fatalf("Error %s\n", err)
	}
}

func getFilenameFromArgs() (string, error) {
	args := os.Args
	if len(args) < 2 {
		return "", errors.New("Usage: gogpt <file>")
	}
	return os.Args[1], nil
}
