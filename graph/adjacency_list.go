package graph

type AdjacencyList struct {
	nodes []Node
}

type Node struct {
	name     string
	visited  bool
	children []Node
}

func depthFirstSearch(root *Node, visit func(node *Node)) {
	if root == nil {
		return
	}
	visit(root)
	root.visited = true
	for _, n := range root.children {
		if !n.visited {
			depthFirstSearch(&n, visit)
		}
	}
}
