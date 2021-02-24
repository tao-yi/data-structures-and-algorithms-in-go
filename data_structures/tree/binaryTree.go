package tree

import (
	"fmt"

	"github.com/tao-yi/data-structure-in-go/data_structures/doubly_linked_list"
)

// binary tree: each node can have at most 2 children
// strict binary tree: each node can have either 0 or 2 children
// complete binary tree: all levels except possibly the last are completey filled and all nodes are as left as possible

// Binary Search Tree: optimized for quick search/update
// for each node, value of all the left subtree is less than or equal to value of all the nodes in right subtree

type BinarySearchTree interface {
	Contains(Interface) bool
	AddNode(Interface)
	DeleteNode(Interface) *node
	InOrder()

	BreadthFirstSearch()
	DepthFirstSearchRec()
	DepthFirstSearchLoop()
}

type tree struct {
	root *node
}

func NewTree(data Interface) BinarySearchTree {
	return &tree{root: newNode(data)}
}

func (t *tree) AddNode(data Interface) {
	t.root.Insert(data)
}

// case1: delete a leaf node
// 		=> simply delete the node from tree
// case2: delete a node has a single child node
// 		=> copy the value of its child to the node and delete the child
// case3: delete a node has two chidren
//    => copy the value of the inorder successor to the node, delete the inorder successor
func (t *tree) DeleteNode(data Interface) *node {
	return deleteNode(t.root, data)
}

func (t *tree) Contains(data Interface) bool {
	return t.root.Search(data)
}

func deleteNode(currentNode *node, data Interface) *node {
	if currentNode == nil {
		return currentNode
	}
	if data.Compare(currentNode.data) < 0 {
		currentNode.left = deleteNode(currentNode, data)
	} else if data.Compare(currentNode.data) > 0 {
		currentNode.right = deleteNode(currentNode, data)
	} else {
		// found the node to be removed
		// if the node has only one child, or no child
		if currentNode.left == nil {
			return currentNode.right
		} else if currentNode.right == nil {
			return currentNode.left
		}
		// if the node has two children
		currentNode.data = getInOrderSuccessor(currentNode).data
	}
	return currentNode
}

func getInOrderSuccessor(node *node) *node {
	successor := node.left
	for successor != nil && successor.left != nil {
		successor = successor.left
	}
	return successor
}

func (t *tree) InOrder() {
	t.root.InOrder()
}

func (t *tree) BreadthFirstSearch() {
	if t.root == nil {
		return
	}

	q := doubly_linked_list.NewQueue()
	q.Enqueue(t.root)

	for !q.IsEmpty() {
		currentNode, ok := q.Dequeue()
		if !ok {
			return
		}
		fmt.Println(currentNode.(*node).data)
		if currentNode.(*node).left != nil {
			q.Enqueue(currentNode.(*node).left)
		}
		if currentNode.(*node).right != nil {
			q.Enqueue(currentNode.(*node).right)
		}
	}
}

func (t *tree) DepthFirstSearchRec() {
	var recVisit func(node *node)
	recVisit = func(node *node) {
		if t.root == nil {
			return
		}
		fmt.Println(node.data)
		if node.left != nil {
			recVisit(node.left)
		}
		if node.right != nil {
			recVisit(node.right)
		}
	}
	recVisit(t.root)
}

func (t *tree) DepthFirstSearchLoop() {
	if t.root == nil {
		return
	}

	s := doubly_linked_list.NewStack()
	s.Push(t.root)

	for !s.IsEmpty() {
		currentNode, ok := s.Pop()
		if !ok || currentNode == nil {
			return
		}
		fmt.Println(currentNode.(*node).data)
		if currentNode.(*node).right != nil {
			s.Push(currentNode.(*node).right)
		}
		if currentNode.(*node).left != nil {
			s.Push(currentNode.(*node).left)
		}
	}
}
