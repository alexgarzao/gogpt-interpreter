package parser

// SymbolTable keeps info about the symbols that were found during the parser.
type SymbolTable struct {
	symbols map[string]int
}

// NewSymbolTable creates a new SymbolTable.
func NewSymbolTable() *SymbolTable {
	st := &SymbolTable{}
	st.symbols = make(map[string]int)

	return st
}

// Add adds a symbol into the table.
func (st *SymbolTable) Add(name string) int {
	_, exist := st.symbols[name]
	if exist {
		return -1
	}

	index := len(st.symbols)
	st.symbols[name] = index

	return index
}

// Index returns the index of a symbol. Returns -1 if not exists.
func (st *SymbolTable) Index(name string) int {
	index, exist := st.symbols[name]
	if !exist {
		return -1
	}

	return index
}
