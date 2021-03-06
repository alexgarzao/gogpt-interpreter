package cp

import (
	"github.com/alexgarzao/gogpt-interpreter/pkg/domain"
)

// CP has the items in a constant pool.
type CP struct {
	constants []interface{}
}

// New creates a new constant pool.
func New() *CP {
	cp := &CP{}
	cp.constants = make([]interface{}, 0)
	return cp
}

// Add adds a new item to the constant pool.
func (cp *CP) Add(item interface{}) int {
	if cp == nil {
		return -1
	}

	if n := cp.Find(item); n != -1 {
		return n
	}

	cp.constants = append(cp.constants, item)

	return len(cp.constants) - 1
}

// Get gets an item from the constant pool.
func (cp *CP) Get(index int) (interface{}, error) {
	if cp == nil || index > len(cp.constants)-1 {
		return 0, domain.ErrIndexNotFound
	}

	res := cp.constants[index]

	return res, nil
}

// Find finds if a specific item is at the constant pool.
func (cp *CP) Find(item interface{}) int {
	if cp == nil {
		return -1
	}

	for i, v := range cp.constants {
		if v == item {
			return i
		}
	}

	return -1
}
