package tree

import (
	"fmt"
	"testing"
)

func TestAddNode(t *testing.T) {
	tree := NewTree(10)
	tree.AddNode(15)
	tree.AddNode(1)
	tree.AddNode(5)
	tree.AddNode(2)
	tree.AddNode(6)
	tree.AddNode(35)
	tree.AddNode(21)
	tree.AddNode(26)

	fmt.Println(tree)

	fmt.Println(tree.Contains(15))
	fmt.Println(tree.Contains(29))
	tree.InOrder()
}