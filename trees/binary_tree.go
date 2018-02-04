package trees

type BinaryTreeNode struct {
	Value int
	Left  *BinaryTreeNode
	Right *BinaryTreeNode
}

func NewBinaryTree(value int, left *BinaryTreeNode, right *BinaryTreeNode) *BinaryTreeNode {
	return &BinaryTreeNode{value, left, right}
}

// visit each node in the left branch, then the current node, then all nodes on the right branch
func inOrderTraversal(t *BinaryTreeNode, visit func(node *BinaryTreeNode)) {
	if t != nil {
		inOrderTraversal(t.Left, visit)
		visit(t)
		inOrderTraversal(t.Right, visit)
	}
}

// visits the current node before its child nodes
func perOrderTraversal(t *BinaryTreeNode, visit func(node *BinaryTreeNode)) {
	if t != nil {
		visit(t)
		perOrderTraversal(t.Left, visit)
		perOrderTraversal(t.Right, visit)
	}
}

// visits the children of the node, visiting the root node last.
func postOrderTraversal(t *BinaryTreeNode, visit func(node *BinaryTreeNode)) {
	if t != nil {
		postOrderTraversal(t.Left, visit)
		postOrderTraversal(t.Right, visit)
		visit(t)
	}
}
