package main

import (
	"log"

	opcodes "github.com/alexgarzao/gpt-interpreter/app/domain"
	lexer "github.com/alexgarzao/gpt-interpreter/app/domain/lexical_analyzer"
	syntax "github.com/alexgarzao/gpt-interpreter/app/domain/syntax_analyzer"
	interfaces "github.com/alexgarzao/gpt-interpreter/app/interface"
	usecase "github.com/alexgarzao/gpt-interpreter/app/usecase"
)

func main() {
	c :=
		`algoritmo olá_mundo;
início
	imprima("Olá mundo!");
fim`
	l := lexer.NewLexer(c)
	p := syntax.NewProgram()

	if p.TryToParse(l) == false {
		log.Fatalln("Error in parsing")
		return
	}

	bce := usecase.NewBytecodeExecutor()
	stdout := interfaces.NewStdout()
	st := opcodes.NewStack()

	err := bce.Run(p.GetCP(), st, stdout, p.GetBC())
	if err != nil {
		log.Fatalf("Error %s\n", err)
	}
}
