package tree

import (
	"fmt"
	"testing"

	"github.com/tao-yi/data-structure-in-go/util"
)

type NodeData int

func NewNodeData(data int) NodeData {
	return NodeData(data)
}

func (n NodeData) Compare(i interface{}) int {
	return int(n) - int(i.(NodeData))
}

func InitTree() BinarySearchTree {
	tree := NewTree(NewNodeData(12))
	tree.AddNode(NewNodeData(5))
	tree.AddNode(NewNodeData(15))
	tree.AddNode(NewNodeData(3))
	tree.AddNode(NewNodeData(7))
	tree.AddNode(NewNodeData(13))
	tree.AddNode(NewNodeData(17))
	tree.AddNode(NewNodeData(1))
	tree.AddNode(NewNodeData(9))
	return tree
}

func TestAddNode(t *testing.T) {
	tree := NewTree(NewNodeData(10))
	tree.AddNode(NewNodeData(15))
	tree.AddNode(NewNodeData(1))
	tree.AddNode(NewNodeData(5))
	tree.AddNode(NewNodeData(2))
	tree.AddNode(NewNodeData(6))
	tree.AddNode(NewNodeData(35))
	tree.AddNode(NewNodeData(21))
	tree.AddNode(NewNodeData(26))

	util.MakeImg("binary-tree", &tree)

	fmt.Println(tree.Contains(NewNodeData(15)))
	fmt.Println(tree.Contains(NewNodeData(29)))
	tree.InOrder()
}

// func TestDeleteNode(t *testing.T) {
// 	tree := InitTree()
// 	fmt.Println(tree)
// 	fmt.Println(tree.Contains(NewNodeData(15)))
// 	fmt.Println(tree.Contains(NewNodeData(29)))
// 	tree.InOrder()

// 	tree.DeleteNode(NewNodeData(15))
// 	tree.InOrder()
// }

func TestBreadthFirstSearch(t *testing.T) {
	tree := InitTree()
	tree.BreadthFirstSearch()
}

func TestDepthFirstSearchRec(t *testing.T) {
	tree := InitTree()
	tree.DepthFirstSearchRec()
}

func TestDepthFirstSearchLoop(t *testing.T) {
	tree := InitTree()
	tree.DepthFirstSearchLoop()
}
