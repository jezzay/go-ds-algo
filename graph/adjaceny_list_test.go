package graph

import (
	"testing"
	"fmt"
)

func TestNewAdjacencyListNode(t *testing.T) {
	root := NewAdjacencyListNode("root", createNodeList("root"))

	if len(root.children) != 3 {
		t.Error("Expected number of children to equal 3")
	}
	if root.children[0].name != "root-child-1" {
		t.Error("Expected first child to be named root-child-1")
	}

}

func TestDepthFirstSearch(t *testing.T) {
	rootChildren := createNodeList("root")
	rootChildren[0].children = createNodeList("child1")
	rootChildren[1].children = createNodeList("child2")
	rootChildren[2].children = createNodeList("child3")
	root := NewAdjacencyListNode("root", rootChildren)

	visitedNodes := make(map[string]bool)

	DepthFirstSearch(&root, func(node *Node) {
		fmt.Printf("Visting %v \n", node.name)
		_, visited := visitedNodes[node.name]
		if visited {
			t.Errorf("Exptect to only visit node %v once, but visited more than once", node)
		}
		visitedNodes[node.name] = true
	})

}

func createNodeList(namePrefix string) []Node {
	nodes := make([]Node, 0)
	nodes = append(nodes, Node{namePrefix + "-child-1", false, nil})
	nodes = append(nodes, Node{namePrefix + "-child-2", false, nil})
	nodes = append(nodes, Node{namePrefix + "-child-3", false, nil})
	return nodes
}
