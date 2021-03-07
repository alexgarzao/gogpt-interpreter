package infrastructure

import (
	"fmt"

	"github.com/alexgarzao/gogpt-interpreter/pkg/domain/usecases/instructions"
)

var _ instructions.StdoutInterface = &Stdout{}

// Stdout is responsible for output data to stdout.
type Stdout struct {
}

// NewStdout creates a new stdout.
func NewStdout() *Stdout {
	return &Stdout{}
}

// Println writes a line to stdout.
func (s *Stdout) Println(text interface{}) {
	fmt.Printf("%v\n", text)
}
