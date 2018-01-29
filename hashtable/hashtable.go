package hashtable

import (
	"math"
)

type ListNode struct {
	key   string
	value string
	next  *ListNode
}

type HashTable interface {
	Put(k string, v string)
}

type Map struct {
	table []*ListNode
}

func New() *Map {
	return &Map{table: make([]*ListNode, 64)}
}

func (m *Map) Put(k string, v string) {
	hash := m.hashKey(k)
	n := m.table[hash]
	for n != nil {
		if n.key == k {
			break
		}
		n = n.next
	}
	if n != nil {
		n.value = v
	} else {
		node := &ListNode{k, v, nil}
		node.next = m.table[hash]
		m.table[hash] = node
	}
}

func (m *Map) Get(k string) (string, bool) {
	hash := m.hashKey(k)
	n := m.table[hash]
	for n != nil {
		if n.key == k {
			return n.value, true
		}
		n = n.next
	}
	return "", false
}

func (m *Map) Contains(k string) (bool) {
	_, ok := m.Get(k)
	return ok
}

func (m *Map) hashKey(k string) int {
	hash := hashCode(k) % len(m.table)
	abs := math.Abs(float64(hash))
	return int(abs)
}

func hashCode(s string) int {
	h := 0
	for i := 0; i < len(s); i++ {
		h = 31*h + int(s[i])
	}
	return h
}
