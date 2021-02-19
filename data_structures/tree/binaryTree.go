package tree

import "fmt"

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
	InOrder()
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

func (t *tree) Contains(data int) bool {
	return t.root.Search(data)
}

func (t *tree) InOrder() {
	t.root.InOrder()
}
