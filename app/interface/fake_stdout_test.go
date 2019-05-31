package interfaces

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringMessage(t *testing.T) {
	fs := NewFakeStdout()
	fs.Println("ABC")
	assert.Equal(t, fs.LastLine, "ABC\n")
}
