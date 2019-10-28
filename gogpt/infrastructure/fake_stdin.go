package infrastructure

// FakeStdin is responsible for mock the stdin.
type FakeStdin struct {
	nextLineToRead string
}

// NewFakeStdin creates a new mocked stdin.
func NewFakeStdin() *FakeStdin {
	return &FakeStdin{}
}

// NextLineToRead defines the next line to be read in a mocked test.
func (s *FakeStdin) NextLineToRead(line string) {
	s.nextLineToRead = line
}

// Readln reads the next mocked line.
func (s *FakeStdin) Readln() string {
	return s.nextLineToRead
}
