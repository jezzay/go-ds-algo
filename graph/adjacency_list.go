package graph

type AdjacencyList struct {
	nodes []Node
}

type Node struct {
	name     string
	visited  bool
	children []Node
}

func NewAdjacencyListNode(name string, children []Node) Node {
	return Node{name, false, children}
}

func DepthFirstSearch(root *Node, visit func(node *Node)) {
	if root == nil {
		return
	}
	visit(root)
	root.visited = true
	for _, n := range root.children {
		if !n.visited {
			DepthFirstSearch(&n, visit)
		}
	}
}
