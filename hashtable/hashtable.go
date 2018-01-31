package hashtable

import (
	"math"
)

type HashTable interface {
	Put(k string, v string)
	Get(k string) (v string, ok bool)
	Contains(k string) (bool)
}

type ListNode struct {
	key   string
	value string
	next  *ListNode
}

type HashMap struct {
	table []*ListNode
}

func New() *HashMap {
	return &HashMap{table: make([]*ListNode, 64)}
}

func (m *HashMap) Put(k string, v string) {
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

func (m *HashMap) Get(k string) (string, bool) {
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

func (m *HashMap) Contains(k string) (bool) {
	_, ok := m.Get(k)
	return ok
}

func (m *HashMap) hashKey(k string) int {
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
