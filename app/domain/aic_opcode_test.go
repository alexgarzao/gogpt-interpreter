package opcodes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidAic(t *testing.T) {
	// AIC 1
	cp := NewCp()
	stack := NewStack()
	i := NewAicOpcode(cp, stack)
	i.Execute(1)
	v, _ := cp.Get(0)
	assert.Equal(t, v, CpItem(1))
}
