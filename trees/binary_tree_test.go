package trees

import (
	"testing"
	"fmt"
)

func TestNew(t *testing.T) {
	tree := NewBinaryTree(5, nil, nil)
	if tree.Value != 5 {
		t.Errorf("Expected value of tree to equal %v , recieved %v", 5, tree.Value)
	}

}

func Test_inOrderTraversal(t *testing.T) {
	visit := func(tree *BinaryTreeNode) {
		fmt.Printf("Tree = %v \n", tree)
	}
	tree := createBinaryTree()
	inOrderTraversal(tree, visit)
}

func Test_perOrderTraversal(t *testing.T)  {
	visit := func(tree *BinaryTreeNode) {
		fmt.Printf("Tree = %v \n", tree)
	}
	tree := createBinaryTree()
	perOrderTraversal(tree, visit)
}

func Test_postOrderTraversal(t *testing.T)  {
	visit := func(tree *BinaryTreeNode) {
		fmt.Printf("Tree = %v \n", tree)
	}
	tree := createBinaryTree()
	postOrderTraversal(tree, visit)
}

func createBinaryTree() *BinaryTreeNode {
	leftLeaf := NewBinaryTree(2, nil, nil)
	leftNext := NewBinaryTree(3, leftLeaf, nil)
	leftRoot := NewBinaryTree(4, leftNext, nil)

	rightLeaf := NewBinaryTree(10, nil, nil)
	rightNext := NewBinaryTree(8, nil, rightLeaf)
	rightRoot := NewBinaryTree(7, nil, rightNext)

	return NewBinaryTree(5, leftRoot, rightRoot)
}
