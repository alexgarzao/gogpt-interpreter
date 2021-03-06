package nopi

import (
	"testing"

	"github.com/alexgarzao/gogpt-interpreter/pkg/domain"
	"github.com/alexgarzao/gogpt-interpreter/pkg/domain/entities/cp"
	"github.com/alexgarzao/gogpt-interpreter/pkg/domain/entities/stack"
	"github.com/alexgarzao/gogpt-interpreter/pkg/domain/entities/vars"
	"github.com/alexgarzao/gogpt-interpreter/pkg/infrastructure"
	"github.com/stretchr/testify/assert"
)

func TestValidNop(t *testing.T) {
	cp := cp.New()
	vars := vars.New()
	st := stack.New()
	stdin := infrastructure.NewFakeStdin()
	stdout := infrastructure.NewFakeStdout()

	// NOP
	nop := New()
	nop.Execute(cp, vars, st, stdin, stdout)
	_, err := cp.Get(0)
	assert.Equal(t, err, domain.ErrIndexNotFound)
	_, err = st.Top()
	assert.Equal(t, err, domain.ErrStackUnderflow)
}
