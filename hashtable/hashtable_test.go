package hashtable

import "testing"

func TestNew(t *testing.T) {
	m := New()
	if m == nil {
		t.Errorf("Expected New() to produce a non nil hashtable")
	}
}

func TestMap_Put(t *testing.T) {
	m := &HashTable{table: make([]*ListNode, 5)}
	m.Put("key", "value")
	node := m.table[4]
	if node == nil {
		t.Fatalf(`Exepected m.Put("key","value") to add to index %v, recieved nil`, 4)
	}
	if node.key != "key" {
		t.Errorf(`Exepected m.Put("key","value") to add node with key: "%v", recieved: "%v"`, "key", node.key)
	}
	if node.value != "value" {
		t.Errorf(`Exepected m.Put("key","value") to add node with value: "%v", recieved: "%v"`, "value", node.value)
	}
}

func TestMap_Get(t *testing.T) {
	m := &HashTable{table: make([]*ListNode, 5)}
	m.Put("key", "value")
	node := m.table[4]
	v, ok := m.Get("key")
	if !ok {
		t.Errorf(`Exepected m.Get("key") to return ok=true: recieved ok=%v`, ok)
	}
	if v != "value" {
		t.Errorf(`Exepected m.Get("key") to return:  %v recieved %v`, "value", v)
	}
	if node.value != v {
		t.Errorf(`Exepected m.Get("key") to return value %v from index %v: recieved %v`, node.value, 4, v)
	}
}

func TestMap_Contains(t *testing.T) {
	m := &HashTable{table: make([]*ListNode, 5)}
	m.Put("key", "value")
	ok := m.Contains("key")
	if !ok {
		t.Errorf(`Exepected m.Get("key") to return ok=true: recieved ok=%v`, ok)
	}
}

func TestHashKey(t *testing.T) {
	m := &HashTable{table: make([]*ListNode, 5)}
	hashKey := m.hashKey("key")
	if hashKey != 4 {
		t.Errorf("Exepected hashKey() to equal %v, recieved %v", 4, hashKey)
	}
	hashKey = m.hashKey("!abc123")
	if hashKey != 2 {
		t.Errorf("Exepected hashKey() to equal %v, recieved %v", 2, hashKey)
	}
}

func TestHashCode(t *testing.T) {
	code := hashCode("key")
	if code != 106079 {
		t.Errorf("Expected hashCode() to equal %v, recieved %v", 106079, code)
	}
	code = hashCode("KEY")
	if code != 74303 {
		t.Errorf("Expected hashCode() to equal %v, recieved %v", 74303, code)
	}
}
