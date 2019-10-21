package adapters

import "fmt"

type FakeStdout struct {
	LastLine       string
	nextLineToRead string
}

func NewFakeStdout() *FakeStdout {
	return &FakeStdout{}
}

func (s *FakeStdout) Println(text interface{}) {
	s.LastLine = fmt.Sprintf("%v\n", text)
}

func (s *FakeStdout) NextLineToRead(line string) {
	s.nextLineToRead = line
}

func (s *FakeStdout) Readln() string {
	return s.nextLineToRead
}
