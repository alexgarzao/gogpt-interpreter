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
func (fs *FakeStdin) NextLineToRead(line string) {
	fs.nextLineToRead = line
}

// Readln reads the next mocked line.
func (fs *FakeStdin) Readln() string {
	return fs.nextLineToRead
}
