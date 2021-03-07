package infrastructure

import (
	"fmt"

	"github.com/alexgarzao/gogpt-interpreter/pkg/domain/usecases/instructions"
)

var _ instructions.StdoutInterface = &FakeStdout{}

// FakeStdout is responsible for mock the stdout.
type FakeStdout struct {
	LastLine string
}

// NewFakeStdout creates a new mocked stdout.
func NewFakeStdout() *FakeStdout {
	return &FakeStdout{}
}

// Println prints the next mocked info.
func (fs *FakeStdout) Println(text interface{}) {
	fs.LastLine = fmt.Sprintf("%v", text)
}
