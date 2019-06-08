package main

import (
	"errors"
	"io/ioutil"
	"log"
	"os"

	opcodes "github.com/alexgarzao/gpt-interpreter/app/domain"
	lexer "github.com/alexgarzao/gpt-interpreter/app/domain/lexical_analyzer"
	syntax "github.com/alexgarzao/gpt-interpreter/app/domain/syntax_analyzer"
	interfaces "github.com/alexgarzao/gpt-interpreter/app/interface"
	usecase "github.com/alexgarzao/gpt-interpreter/app/usecase"
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

	bce := usecase.NewBytecodeExecutor()
	stdout := interfaces.NewStdout()
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
