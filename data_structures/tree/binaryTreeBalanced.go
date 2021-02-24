package tree

type BalancedBinarySearchTree interface {
	BinarySearchTree
	// make sure to call leftRotate or rightRotate after each insert
	// if violate balance
	leftRotate()
	rightRotate()
}

type balancedBST struct {
	root *node
}

func NewBalancedBST(data Interface) BalancedBinarySearchTree {
	return &balancedBST{root: newNode(data)}
}

func (t *balancedBST) Contains(Interface) bool {
	return false
}

func (t *balancedBST) AddNode(Interface) {}

func (t *balancedBST) DeleteNode(Interface) *node {
	return nil
}

func (t *balancedBST) InOrder()              {}
func (t *balancedBST) BreadthFirstSearch()   {}
func (t *balancedBST) DepthFirstSearchRec()  {}
func (t *balancedBST) DepthFirstSearchLoop() {}

func (t *balancedBST) leftRotate()  {}
func (t *balancedBST) rightRotate() {}
