package binaryTree

type BinaryTree struct {
	Root *BinaryTreeNode
}

type BinaryTreeNode struct {
	Val   int
	Left  *BinaryTreeNode
	Right *BinaryTreeNode
}

// Creates a new node in the binary tree
func NewNode(val int) BinaryTreeNode {
	return BinaryTreeNode{
		Val:   val,
		Left:  nil,
		Right: nil,
	}
}

// Inserts a new node on the binary tree
// or returns an error otherwise
func (bst *BinaryTree) InsertNode(root *BinaryTreeNode, val int) *BinaryTreeNode {
	if root == nil {
		return &BinaryTreeNode{
			Val:   val,
			Left:  nil,
			Right: nil,
		}
	}
	return root
}

// Visits all nodes in-order traversing
func (bst *BinaryTree) InOrderTraverse() []int {
	return []int{}
}

// Prints a CLI readable rendering of the tree
func (n *BinaryTreeNode) String() {
}

// Returns true if Item exists in tree
func (n *BinaryTreeNode) Search(value int) bool {
	return true
}

// Visits all nodes in pre-order traversing
func (n *BinaryTreeNode) PreOrderTraverse() {

}

// Visits all nodes with post-order traversing
func (n *BinaryTreeNode) PostOrderTraverse() {

}

// Returns item with the min value stored in tree
func (n *BinaryTreeNode) Min() int {
	return 0
}

// Returns item with the max value stored in tree
func (n *BinaryTreeNode) Max() int {
	return 0
}

// Removes the item from the Binary tree
func (n *BinaryTreeNode) Remove(t BinaryTreeNode) {
}
