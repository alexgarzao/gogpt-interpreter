package parser

type SymbolTable struct {
	symbols map[string]int
}

func NewSymbolTable() *SymbolTable {
	st := &SymbolTable{}
	st.symbols = make(map[string]int)

	return st
}

func (st *SymbolTable) Add(name string) int {
	_, exist := st.symbols[name]
	if exist {
		return -1
	}

	index := len(st.symbols)
	st.symbols[name] = index

	return index
}

func (st *SymbolTable) Index(name string) int {
	index, exist := st.symbols[name]
	if !exist {
		return -1
	}

	return index
}
