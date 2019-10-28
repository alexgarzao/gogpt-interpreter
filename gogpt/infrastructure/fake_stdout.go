package infrastructure

import "fmt"

// FakeStdout is responsible for mock the stdout.
type FakeStdout struct {
	LastLine string
}

// NewFakeStdout creates a new mocked stdout.
func NewFakeStdout() *FakeStdout {
	return &FakeStdout{}
}

// Println prints the next mocked info.
func (s *FakeStdout) Println(text interface{}) {
	s.LastLine = fmt.Sprintf("%v\n", text)
}
