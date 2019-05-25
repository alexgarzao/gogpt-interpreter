package opcodes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidAic(t *testing.T) {
	// AIC 1
	cp := 0
	stack := 0
	i := NewAicOpcode(&cp, &stack)
	i.Execute(1)
	assert.True(t, cp == 1)
}
