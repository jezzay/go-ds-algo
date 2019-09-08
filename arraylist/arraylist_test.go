package arraylist

import "testing"

func TestArrayList_Insert(t *testing.T) {
	list := New()
	list.Insert("hello")
	list.Insert("world")
	list.Insert("!!!!")
}

func TestArrayList_Remove(t *testing.T) {
	list := New()
	list.Insert("hello")
	list.Insert("world")
	list.Insert("!!!!")

	list.Remove(1)
	list.Insert("Good")
	list.Insert("Day")
	list.Remove(3)
	list.Remove(3)
}
