package binarySearchTree

import "fmt"

type BinarySearchTree struct {
	Root *BinarySearchTreeNode
}

type BinarySearchTreeNode struct {
	Val   int
	Left  *BinarySearchTreeNode
	Right *BinarySearchTreeNode
}

// Creates a new node in the binary tree
func NewNode(val int) BinarySearchTreeNode {
	return BinarySearchTreeNode{
		Val:   val,
		Left:  nil,
		Right: nil,
	}
}

// Inserts a new node on the binary tree
// or returns an error otherwise
func (bst *BinarySearchTree) InsertNode(root *BinarySearchTreeNode, val int) *BinarySearchTreeNode {

	if root == nil {
		return &BinarySearchTreeNode{
			Val:   val,
			Left:  nil,
			Right: nil,
		}
	}

	if val > root.Val {
		if root.Right == nil {
			n := &BinarySearchTreeNode{
				Val:   val,
				Left:  nil,
				Right: nil,
			}
			root.Right = n
		} else {
			bst.InsertNode(root.Right, val)
		}

	} else if val < root.Val {
		if root.Left == nil {
			root.Left = &BinarySearchTreeNode{val, nil, nil}
		} else {
			bst.InsertNode(root.Left, val)
		}
	}
	return root
}

// Calculates the height of the tree
func (bst *BinarySearchTree) Height(root *BinarySearchTreeNode) int {
	if root == nil {
		return 0
	}

	left := bst.Height(root.Left)
	right := bst.Height(root.Right)

	if left > right {
		return left + 1
	} else {
		return right + 1
	}
}

// Returns the In Order Transversal as an slice of values
func (bst *BinarySearchTree) InOrderTraverse(root *BinarySearchTreeNode) {
	if root == nil {
		return
	}

	bst.InOrderTraverse(root.Left)
	fmt.Printf("%d ", root.Val)
	bst.InOrderTraverse(root.Right)
}

// Returns the In Order Transversal as an slice of values
func (bst *BinarySearchTree) InOrderTraverseArray(root *BinarySearchTreeNode) []int {
	vals := make([]int, 0)
	bst.inOrderTraverseArrayHelper(root, &vals)
	return vals
}

func (bst *BinarySearchTree) inOrderTraverseArrayHelper(root *BinarySearchTreeNode, vals *[]int) {
	if root == nil {
		return
	}
	bst.inOrderTraverseArrayHelper(root.Left, vals)
	*vals = append(*vals, root.Val)
	bst.inOrderTraverseArrayHelper(root.Right, vals)
}

// Returns the In Order Transversal as an slice of values
func (bst *BinarySearchTree) PreOrderTraverseArray(root *BinarySearchTreeNode) []int {
	vals := make([]int, 0)
	bst.preOrderTraverseArrayHelper(root, &vals)
	return vals
}

func (bst *BinarySearchTree) preOrderTraverseArrayHelper(root *BinarySearchTreeNode, vals *[]int) {
	if root == nil {
		return
	}
	*vals = append(*vals, root.Val)
	bst.preOrderTraverseArrayHelper(root.Left, vals)
	bst.preOrderTraverseArrayHelper(root.Right, vals)
}

// Prints a CLI readable rendering of the tree
func (n *BinarySearchTreeNode) String() {
}

// Returns true if Item exists in tree
func (n *BinarySearchTreeNode) Search(value int) bool {
	return true
}

// Visits all nodes in pre-order traversing
func (n *BinarySearchTreeNode) PreOrderTraverse() {

}

// Visits all nodes with post-order traversing
func (n *BinarySearchTreeNode) PostOrderTraverse() {

}

// Returns item with the min value stored in tree
func (n *BinarySearchTreeNode) Min() int {
	return 0
}

// Returns item with the max value stored in tree
func (n *BinarySearchTreeNode) Max() int {
	return 0
}

// Removes the item from the Binary tree
func (n *BinarySearchTreeNode) Remove(t BinarySearchTreeNode) {
}
