package main

import (
	"errors"
	"io/ioutil"
	"log"
	"os"

	adapters "github.com/alexgarzao/gpt-interpreter/gpt/adapters"
	opcodes "github.com/alexgarzao/gpt-interpreter/gpt/entities"
	lexer "github.com/alexgarzao/gpt-interpreter/gpt/entities/lexical_analyzer"
	syntax "github.com/alexgarzao/gpt-interpreter/gpt/entities/syntax_analyzer"
	usecases "github.com/alexgarzao/gpt-interpreter/gpt/usecases"
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

	bce := usecases.NewBytecodeExecutor()
	stdout := adapters.NewStdout()
	st := opcodes.NewStack()

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
