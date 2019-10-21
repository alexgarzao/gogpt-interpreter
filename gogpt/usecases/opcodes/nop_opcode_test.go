package opcodes

import (
	"testing"

	"github.com/alexgarzao/gogpt-interpreter/gogpt/entities/constant_pool"
	"github.com/alexgarzao/gogpt-interpreter/gogpt/entities/stack"
	"github.com/alexgarzao/gogpt-interpreter/gogpt/entities/vars"

	"github.com/alexgarzao/gogpt-interpreter/gogpt/adapters"
	"github.com/stretchr/testify/assert"
)

func TestValidNop(t *testing.T) {
	cp := constant_pool.NewCp()
	vars := vars.NewVars()
	st := stack.NewStack()
	stdout := adapters.NewFakeStdout()

	// NOP
	i := NewNopOpcode()
	i.Execute(cp, vars, st, stdout)
	_, err := cp.Get(0)
	assert.EqualError(t, err, "Index not found")
	_, err = st.Top()
	assert.EqualError(t, err, "Stack underflow")
}
