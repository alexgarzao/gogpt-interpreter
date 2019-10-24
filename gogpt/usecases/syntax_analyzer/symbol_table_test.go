package syntax

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddingSymbolsToSymbolTable(t *testing.T) {
	st := NewSymbolTable()

	assert.Equal(t, 0, st.Add("Var1"))
	assert.Equal(t, 1, st.Add("Var2"))
	assert.Equal(t, -1, st.Add("Var2"))

	assert.Equal(t, 0, st.Index("Var1"))
	assert.Equal(t, 1, st.Index("Var2"))
	assert.Equal(t, -1, st.Index("Var3"))
}
