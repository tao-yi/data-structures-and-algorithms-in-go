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

type node struct {
	data  int
	left  *node
	right *node
}

func NewNode(data int) *node {
	return &node{data: data}
}

func (n *node) Insert(data int) {
	if data < n.data {
		if n.left == nil {
			n.left = NewNode(data)
		} else {
			n.left.Insert(data)
		}
	} else {
		if n.right == nil {
			n.right = NewNode(data)
		} else {
			n.right.Insert(data)
		}
	}
}

func (n *node) Search(data int) bool {
	if data == n.data {
		return true
	}

	if data < n.data {
		if n.left == nil {
			return false
		}
		return n.left.Search(data)
	} else {
		if n.right == nil {
			return false
		}
		return n.right.Search(data)
	}
}

func (n *node) InOrder() {
	if n.left != nil {
		n.left.InOrder()
	}
	// process data
	fmt.Println(n.data)
	if n.right != nil {
		n.right.InOrder()
	}
}

type BinarySearchTree interface {
	Contains(int) bool
	AddNode(int)
	DeleteNode(int) bool
	InOrder()

	BreadthFirstSearch()
	DepthFirstSearchRec()
	DepthFirstSearchLoop()
}

type tree struct {
	root *node
}

func NewTree(data int) BinarySearchTree {
	return &tree{root: NewNode(data)}
}

func (t *tree) AddNode(data int) {
	t.root.Insert(data)
}

func (t *tree) DeleteNode(data int) bool {
	return false
}

func (t *tree) Contains(data int) bool {
	return t.root.Search(data)
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
