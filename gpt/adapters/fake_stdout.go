package interfaces

import "fmt"

type FakeStdout struct {
	LastLine string
}

func NewFakeStdout() *FakeStdout {
	return &FakeStdout{}
}

func (s *FakeStdout) Println(text interface{}) {
	s.LastLine = fmt.Sprintf("%v\n", text)
}
