package adapters

import (
	"bufio"
	"fmt"
	"os"
)

type Stdout struct {
}

func NewStdout() *Stdout {
	return &Stdout{}
}

func (s *Stdout) Println(text interface{}) {
	fmt.Printf("%v\n", text)
}

func (s *Stdout) Readln() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	text := scanner.Text()

	return text
}
