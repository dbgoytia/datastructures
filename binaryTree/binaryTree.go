package binaryTree

import "fmt"

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

	if val > root.Val {
		if root.Right == nil {
			n := &BinaryTreeNode{
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
			root.Left = &BinaryTreeNode{val, nil, nil}
		} else {
			bst.InsertNode(root.Left, val)
		}
	}
	return root
}

// Calculates the height of the tree
func (bst *BinaryTree) Height(root *BinaryTreeNode) int {
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
func (bst *BinaryTree) InOrderTraverse(root *BinaryTreeNode) {
	if root == nil {
		return
	}

	bst.InOrderTraverse(root.Left)
	fmt.Printf("%d ", root.Val)
	bst.InOrderTraverse(root.Right)
}

// Returns the In Order Transversal as an slice of values
func (bst *BinaryTree) InOrderTraverseArray(root *BinaryTreeNode) []int {
	vals := make([]int, 0)
	bst.inOrderTraverseArrayHelper(root, &vals)
	return vals
}

func (bst *BinaryTree) inOrderTraverseArrayHelper(root *BinaryTreeNode, vals *[]int) {
	if root == nil {
		return
	}
	bst.inOrderTraverseArrayHelper(root.Left, vals)
	*vals = append(*vals, root.Val)
	bst.inOrderTraverseArrayHelper(root.Right, vals)
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
