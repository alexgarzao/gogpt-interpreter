package cp

import (
	"errors"
)

// CP has the items in a constant pool.
type CP struct {
	items []interface{}
}

// New creates a new constant pool.
func New() *CP {
	cp := &CP{}
	cp.items = make([]interface{}, 0)
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

	cp.items = append(cp.items, item)
	return len(cp.items) - 1
}

// Get gets an item from the constant pool.
func (cp *CP) Get(index int) (interface{}, error) {
	if index > len(cp.items)-1 {
		return 0, errors.New("Index not found")
	}

	res := cp.items[index]

	return res, nil
}

// Find finds if a specific item is at the constant pool.
func (cp *CP) Find(item interface{}) int {
	for i, v := range cp.items {
		if v == item {
			return i
		}
	}

	return -1
}
