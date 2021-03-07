package main

import (
	"errors"
	"io/ioutil"
	"log"
	"os"

	"github.com/alexgarzao/gogpt-interpreter/pkg/domain/entities/stack"
	"github.com/alexgarzao/gogpt-interpreter/pkg/domain/entities/vars"
	"github.com/alexgarzao/gogpt-interpreter/pkg/domain/usecases/bce"
	"github.com/alexgarzao/gogpt-interpreter/pkg/domain/usecases/lexer"
	"github.com/alexgarzao/gogpt-interpreter/pkg/domain/usecases/parser"
	"github.com/alexgarzao/gogpt-interpreter/pkg/infrastructure"
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

	l := lexer.New(string(algorithm))
	p := parser.New(l)

	err = p.Parser()
	if err != nil {
		log.Fatalf("Error in parsing: %v", err)
		return
	}

	bce := bce.New(p.GetBC())
	stdin := infrastructure.NewStdin()
	stdout := infrastructure.NewStdout()
	st := stack.New()
	vars := vars.New()

	err = bce.Run(p.GetCP(), vars, st, stdin, stdout)
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
