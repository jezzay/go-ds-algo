package main

import (
	"github.com/jezzay/go-ds-algo/hashtable"
	"fmt"
)

func main() {
	h := hashtable.New()
	h.Put("key", "value")
	h.Put("key2", "value2")
	value, ok := h.Get("key")
	if ok {
		fmt.Println(value)
	}

}
