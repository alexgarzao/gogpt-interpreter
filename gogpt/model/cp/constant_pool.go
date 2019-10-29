package cp

import (
	"errors"
)

// CPItem has the type of a Constant Pool item.
type CPItem interface{}

// CP has the items in a constant pool.
type CP struct {
	items []CPItem
}

// NewCP creates a new constant pool.
func NewCP() *CP {
	cp := &CP{}
	cp.items = make([]CPItem, 0)
	return cp
}

// Add adds a new item to the constant pool.
func (cp *CP) Add(item CPItem) int {
	if cp == nil {
		return -1
	}

	if n := cp.Find(item); n != -1 {
		return n
	}

	cp.items = append(cp.items, item)
	return len(cp.items) - 1
}

// Get gets an item from the constant pool.
func (cp *CP) Get(index int) (CPItem, error) {
	if index > len(cp.items)-1 {
		return 0, errors.New("Index not found")
	}

	res := cp.items[index]

	return res, nil
}

// Find finds if a specific item is at the constant pool.
func (cp *CP) Find(item CPItem) int {
	for i, v := range cp.items {
		if v == item {
			return i
		}
	}

	return -1
}
