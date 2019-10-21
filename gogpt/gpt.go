package main

import (
	"errors"
	"io/ioutil"
	"log"
	"os"

	"github.com/alexgarzao/gogpt-interpreter/gogpt/adapters"
	lexer "github.com/alexgarzao/gogpt-interpreter/gogpt/entities/lexical_analyzer"
	"github.com/alexgarzao/gogpt-interpreter/gogpt/entities/stack"
	"github.com/alexgarzao/gogpt-interpreter/gogpt/entities/vars"
	bce "github.com/alexgarzao/gogpt-interpreter/gogpt/usecases/bytecode_executor"
	syntax "github.com/alexgarzao/gogpt-interpreter/gogpt/usecases/syntax_analyzer"
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
	p := syntax.NewAlgorithm(l)

	pr := p.Parser()
	if pr.Parsed == false {
		log.Fatalf("Error in parsing: %v", pr.Err)
		return
	}

	bce := bce.NewBytecodeExecutor(p.GetBC())
	stdout := adapters.NewStdout()
	st := stack.NewStack()
	vars := vars.NewVars()

	err = bce.Run(p.GetCP(), vars, st, stdout)
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
