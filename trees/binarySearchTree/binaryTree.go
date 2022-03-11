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
func (bst *BinarySearchTree) InsertNode(val int) {
	if bst.Root == nil {
		bst.Root = &BinarySearchTreeNode{
			Val:   val,
			Left:  nil,
			Right: nil,
		}
	}
	bst.insertNodeHelper(bst.Root, val)
}

func (bst *BinarySearchTree) insertNodeHelper(root *BinarySearchTreeNode, val int) {
	if root.Val == val {
		return
	}
	if val > root.Val {
		if root.Right == nil {
			root.Right = &BinarySearchTreeNode{val, nil, nil}
		} else {
			bst.insertNodeHelper(root.Right, val)
		}
	} else {
		if root.Left == nil {
			root.Left = &BinarySearchTreeNode{val, nil, nil}
		} else {
			bst.insertNodeHelper(root.Left, val)
		}
	}
}

// Removes the item from the Binary tree
func (n *BinarySearchTreeNode) Remove(val int) {

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

// Returns the Post order transversal as an slice of values
func (bst *BinarySearchTree) PostOrderTraverseArray(root *BinarySearchTreeNode) []int {
	vals := make([]int, 0)
	bst.postOrderTraversalArrayHelper(root, &vals)
	return vals
}

func (bst *BinarySearchTree) postOrderTraversalArrayHelper(root *BinarySearchTreeNode, vals *[]int) {
	if root == nil {
		return
	}
	bst.postOrderTraversalArrayHelper(root.Left, vals)
	bst.postOrderTraversalArrayHelper(root.Right, vals)
	*vals = append(*vals, root.Val)
}

// Returns the elements of a tree in breadth-first arrays (level order)
func (bst *BinarySearchTree) BreadFirstArray(root *BinarySearchTreeNode) []int {
	vals := make([]int, 0)
	height := bst.Height(root)
	for i := 0; i <= height; i++ {
		bst.breadFirstArrayHelper(root, &vals, i)
	}
	return vals
}

func (bst *BinarySearchTree) breadFirstArrayHelper(root *BinarySearchTreeNode, vals *[]int, level int) {

	if root == nil {
		return
	}

	if level == 1 {
		*vals = append(*vals, root.Val)
	}

	if level > 1 {
		bst.breadFirstArrayHelper(root.Left, vals, level-1)
		bst.breadFirstArrayHelper(root.Right, vals, level-1)
	}

}

// Returns item with the min value stored in tree
func (bst *BinarySearchTree) Min(root *BinarySearchTreeNode) int {
	n := root
	for {
		if n.Left == nil {
			return n.Val
		}
		n = n.Left
	}
}

// Returns item with the max value stored in tree
func (bst *BinarySearchTree) Max(root *BinarySearchTreeNode) int {
	n := root
	for {
		if n.Right == nil {
			return n.Val
		}
		n = n.Right
	}
}

// Returns true if Item exists in tree
func (bst *BinarySearchTree) BinarySearch(root *BinarySearchTreeNode, value int) bool {

	if root == nil {
		return false
	}

	if root.Val == value {
		return true
	}

	if root.Val <= value {
		return bst.BinarySearch(root.Right, value)
	} else {
		return bst.BinarySearch(root.Left, value)
	}
}
