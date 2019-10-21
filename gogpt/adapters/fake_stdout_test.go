package adapters

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOutputStringMessage(t *testing.T) {
	fs := NewFakeStdout()
	fs.Println("ABC")
	assert.Equal(t, fs.LastLine, "ABC\n")
}

func TestNextStringToRead(t *testing.T) {
	fs := NewFakeStdout()
	fs.NextLineToRead("aaa123")
	assert.Equal(t, fs.Readln(), "aaa123")
}
