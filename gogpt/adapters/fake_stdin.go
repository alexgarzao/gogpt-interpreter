package adapters

type FakeStdin struct {
	nextLineToRead string
}

func NewFakeStdin() *FakeStdin {
	return &FakeStdin{}
}

func (s *FakeStdin) NextLineToRead(line string) {
	s.nextLineToRead = line
}

func (s *FakeStdin) Readln() string {
	return s.nextLineToRead
}
