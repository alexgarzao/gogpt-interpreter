package adapters

import (
	"bufio"
	"os"
)

type Stdin struct {
}

func NewStdin() *Stdin {
	return &Stdin{}
}

func (s *Stdin) Readln() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	text := scanner.Text()

	return text
}
