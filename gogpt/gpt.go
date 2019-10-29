package main

import (
	"errors"
	"io/ioutil"
	"log"
	"os"

	"github.com/alexgarzao/gogpt-interpreter/gogpt/infrastructure"
	"github.com/alexgarzao/gogpt-interpreter/gogpt/model/stack"
	"github.com/alexgarzao/gogpt-interpreter/gogpt/model/vars"
	"github.com/alexgarzao/gogpt-interpreter/gogpt/services/bce"
	"github.com/alexgarzao/gogpt-interpreter/gogpt/services/lexer"
	"github.com/alexgarzao/gogpt-interpreter/gogpt/services/parser"
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

	pr := p.Parser()
	if pr.Parsed == false {
		log.Fatalf("Error in parsing: %v", pr.Err)
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
