package opcodes

import (
	"testing"

	"github.com/alexgarzao/gogpt-interpreter/gogpt/infrastructure"
	"github.com/alexgarzao/gogpt-interpreter/gogpt/model/cp"
	"github.com/alexgarzao/gogpt-interpreter/gogpt/model/stack"
	"github.com/alexgarzao/gogpt-interpreter/gogpt/model/vars"
	"github.com/stretchr/testify/assert"
)

func TestValidNop(t *testing.T) {
	cp := cp.NewCP()
	vars := vars.NewVars()
	st := stack.NewStack()
	stdin := infrastructure.NewFakeStdin()
	stdout := infrastructure.NewFakeStdout()

	// NOP
	i := NewNOPOpcode()
	i.Execute(cp, vars, st, stdin, stdout)
	_, err := cp.Get(0)
	assert.EqualError(t, err, "Index not found")
	_, err = st.Top()
	assert.EqualError(t, err, "Stack underflow")
}
