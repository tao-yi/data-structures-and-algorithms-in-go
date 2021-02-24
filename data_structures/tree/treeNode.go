package tree

import "fmt"

type Interface interface {
	Compare(interface{}) int
}

type node struct {
	data  Interface
	left  *node
	right *node
}

func newNode(data Interface) *node {
	return &node{data: data}
}

func (n *node) Insert(data Interface) {
	if data.Compare(n.data) < 0 {
		if n.left == nil {
			n.left = newNode(data)
		} else {
			n.left.Insert(data)
		}
	} else {
		if n.right == nil {
			n.right = newNode(data)
		} else {
			n.right.Insert(data)
		}
	}
}

func (n *node) Search(data Interface) bool {
	if data == n.data {
		return true
	}

	if data.Compare(n.data) < 0 {
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
