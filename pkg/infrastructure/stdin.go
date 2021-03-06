package infrastructure

import (
	"bufio"
	"os"
)

// Stdin is responsible for input data from stdin.
type Stdin struct {
}

// NewStdin creates a new stdin.
func NewStdin() *Stdin {
	return &Stdin{}
}

// Readln reads a line from the stdin.
func (s *Stdin) Readln() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	text := scanner.Text()

	return text
}
