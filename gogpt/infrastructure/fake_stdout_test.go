package infrastructure

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOutputStringMessage(t *testing.T) {
	fs := NewFakeStdout()
	fs.Println("ABC")
	assert.Equal(t, "ABC", fs.LastLine)
}
