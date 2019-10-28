package infrastructure

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNextStringToRead(t *testing.T) {
	fs := NewFakeStdin()
	fs.NextLineToRead("aaa123")
	assert.Equal(t, fs.Readln(), "aaa123")
}
