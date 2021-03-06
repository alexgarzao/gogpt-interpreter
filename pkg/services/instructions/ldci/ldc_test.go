package ldci

import (
	"testing"

	"github.com/alexgarzao/gogpt-interpreter/pkg/infrastructure"
	"github.com/alexgarzao/gogpt-interpreter/pkg/model/cp"
	"github.com/alexgarzao/gogpt-interpreter/pkg/model/stack"
	"github.com/alexgarzao/gogpt-interpreter/pkg/model/vars"
	"github.com/stretchr/testify/assert"
)

func TestValidLdcInt123(t *testing.T) {
	// CP map:
	//		0: (INT) 123
	cp := cp.New()
	cpIndex := cp.Add(123)
	vars := vars.New()
	st := stack.New()
	stdin := infrastructure.NewFakeStdin()
	stdout := infrastructure.NewFakeStdout()
	cpv, _ := cp.Get(0)
	assert.Equal(t, 123, cpv)
	stv, _ := st.Top()
	assert.Equal(t, 0, stv)

	// LDC 0
	ldc := New()
	ldc.CpIndex = cpIndex
	ldc.Execute(cp, vars, st, stdin, stdout)
	cpv, _ = cp.Get(0)
	assert.Equal(t, 123, cpv)
	stv, _ = st.Top()
	assert.Equal(t, 123, stv)
}

func TestValidLdcABC(t *testing.T) {
	// CP map:
	//		0: STR: "ABC"
	cp := cp.New()
	cpIndex := cp.Add("ABC")
	vars := vars.New()
	st := stack.New()
	stdin := infrastructure.NewFakeStdin()
	stdout := infrastructure.NewFakeStdout()
	cpv, _ := cp.Get(0)
	assert.Equal(t, "ABC", cpv)
	stv, _ := st.Top()
	assert.Equal(t, 0, stv)

	// LDC 0
	ldc := New()
	ldc.CpIndex = cpIndex
	ldc.Execute(cp, vars, st, stdin, stdout)
	cpv, _ = cp.Get(0)
	assert.Equal(t, "ABC", cpv)
	stv, _ = st.Top()
	assert.Equal(t, "ABC", stv)
}
