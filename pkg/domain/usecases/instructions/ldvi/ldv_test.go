package ldvi

import (
	"testing"

	"github.com/alexgarzao/gogpt-interpreter/pkg/domain/entities/cp"
	"github.com/alexgarzao/gogpt-interpreter/pkg/domain/entities/stack"
	"github.com/alexgarzao/gogpt-interpreter/pkg/domain/entities/vars"
	"github.com/alexgarzao/gogpt-interpreter/pkg/domain/usecases/instructions/ldci"
	"github.com/alexgarzao/gogpt-interpreter/pkg/domain/usecases/instructions/stvi"
	"github.com/alexgarzao/gogpt-interpreter/pkg/infrastructure"
	"github.com/stretchr/testify/assert"
)

func TestValidLdvInt(t *testing.T) {
	// CP map:
	//		0: (INT) 123
	// VAR map:
	//		0: (INT) Value
	cp := cp.New()
	cpIndex := cp.Add(123)
	st := stack.New()
	stdin := infrastructure.NewFakeStdin()
	stdout := infrastructure.NewFakeStdout()
	vars := vars.New()
	varIndex := 0

	// LDC 0
	ldc := ldci.New()
	ldc.CpIndex = cpIndex
	ldc.Execute(cp, vars, st, stdin, stdout)

	// STV 0
	stv := stvi.New()
	stv.VarIndex = varIndex
	stv.Execute(cp, vars, st, stdin, stdout)
	vv, _ := vars.Get(varIndex)
	assert.Equal(t, 123, vv)
	assert.Equal(t, 0, st.Size())

	// LDV 0
	ldvInst := New()
	ldvInst.VarIndex = varIndex
	ldvInst.Execute(cp, vars, st, stdin, stdout)

	stackValue, _ := st.Top()
	assert.Equal(t, 123, stackValue)
}

func TestValidLdvStr(t *testing.T) {
	// CP map:
	//		0: (STR) ABC
	// VAR map:
	//		0: (STR) Value
	cp := cp.New()
	cpIndex := cp.Add("ABC")
	st := stack.New()
	stdin := infrastructure.NewFakeStdin()
	stdout := infrastructure.NewFakeStdout()
	vars := vars.New()
	varIndex := 0

	// LDC 0
	ldc := ldci.New()
	ldc.CpIndex = cpIndex
	ldc.Execute(cp, vars, st, stdin, stdout)

	// STV 0
	stv := stvi.New()
	stv.VarIndex = varIndex
	stv.Execute(cp, vars, st, stdin, stdout)
	vv, _ := vars.Get(varIndex)
	assert.Equal(t, "ABC", vv)
	assert.Equal(t, 0, st.Size())

	// LDV 0
	ldv := New()
	ldv.VarIndex = varIndex
	ldv.Execute(cp, vars, st, stdin, stdout)

	stackValue, _ := st.Top()
	assert.Equal(t, "ABC", stackValue)
}
