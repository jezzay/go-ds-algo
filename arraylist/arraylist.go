package arraylist

type ArrayList struct {
	list         []string
	size         int
	growthFactor int
}

func New() *ArrayList {
	return &ArrayList{growthFactor: 2}
}

func (a *ArrayList) Insert(val string) {
	a.growBy(1)
	a.list[a.size] = val
	a.size++
}

func (a *ArrayList) Remove(index int) {
	if index >= a.size {
		return
	}
	if index == a.size-1 {
		a.list = a.list[:index]
		a.size--
		return
	}
	before := a.list[:index]
	after := a.list[index+1:]
	a.list = append(before, after...)
	a.size--
}

func (a *ArrayList) growBy(size int) {
	currentCap := cap(a.list)
	if a.size+size >= currentCap {
		newSize := a.growthFactor * (len(a.list) + size)
		newList := make([]string, newSize, newSize)
		copy(newList, a.list)
		a.list = newList
	}
}
