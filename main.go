package main

import (
	"github.com/jezzay/go-ds-algo/hashtable"
	"fmt"
	"github.com/jezzay/go-ds-algo/trees"
)

func main() {
	h := hashtable.New()
	testHashTable(h)

	bt := trees.NewBinaryTree(5, nil, nil)
	fmt.Printf("%v+", bt)
}

func testHashTable(h *hashtable.HashTable) {
	h.Put("key", "value")
	h.Put("key2", "value2")
	value, ok := h.Get("key")
	if ok {
		fmt.Println(value)
	}
}
