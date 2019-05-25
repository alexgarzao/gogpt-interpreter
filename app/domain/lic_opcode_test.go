package opcodes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidLic(t *testing.T) {
	cp := 0
	stack := 0
	// AIC 1
	i := NewAicOpcode(&cp, &stack)
	i.Execute(1)
	assert.True(t, cp == 1)
	assert.True(t, stack == 0)

	// LIC 0
	j := NewLicOpcode(&cp, &stack)
	j.Execute(0)
	assert.True(t, cp == 1)
	assert.True(t, stack == 1)
}
