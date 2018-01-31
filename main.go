package main

import (
	"github.com/jezzay/go-ds-algo/hashtable"
	"fmt"
)

func main() {
	h := hashtable.New()
	testHashTable(h)
}

func testHashTable(h *hashtable.HashTable) {
	h.Put("key", "value")
	h.Put("key2", "value2")
	value, ok := h.Get("key")
	if ok {
		fmt.Println(value)
	}
}
