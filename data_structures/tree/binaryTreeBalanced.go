package tree

type BalancedBinarySearchTree interface {
	BinarySearchTree
	// make sure to call leftRotate or rightRotate after each insert
	// if violate balance
	leftRotate(*node) *node
	rightRotate(*node) *node
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

/*
		 4
	3		  6
		  5		8
becomes:
	   6
  4		  8
3   5
leftRotate takes grandparent node as parameter
*/
func (t *balancedBST) leftRotate(node *node) *node {
	tmp := node.right
	node.right = tmp.left
	tmp.left = node
	return tmp
}

/*
			 8
		4		 10
	2		5
1
becomes:
		  4
		2		5
	1		8		10
*/
func (t *balancedBST) rightRotate(node *node) *node {
	tmp := node.left
	node.left = tmp.right
	tmp.right = node
	return tmp
}

/*
4
	8
6
first do a right rotate of parent
4
	6
		8
then left rotate of grandparent
	6
4		8
*/
func (t *balancedBST) rightLeftRotate(node *node) *node {
	// 4.right becomes 6
	node.right = t.rightRotate(node.right)
	return t.leftRotate(node)
}

/*
	8
4
	6
first do a left rotate of parent
		8
	6
4
then right rotate of grandparent
	6
4		8
*/
func (t *balancedBST) leftRightRotate(node *node) *node {
	node.left = t.leftRotate(node.left)
	return t.rightRotate(node)
}
