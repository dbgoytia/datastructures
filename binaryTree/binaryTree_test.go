package binaryTree_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/dbgoytia/datastructures/binaryTree"
)

var _ = Describe("BinaryTree", func() {

	It("Inserting a value when root is nil becomes the root", func() {
		// Expected tree:
		//		  61
		// Basic tree
		tree := binaryTree.BinaryTree{}
		// A node
		node := binaryTree.BinaryTreeNode{
			61,
			nil,
			nil,
		}
		tree.InsertNode(&node, 61)
		Expect(tree.Root == &node)
	})

	It("Inserts value in the correct position when received in disorder", func() {
		// Expected tree:
		//		  4
		//       / \
		//      2   7
		//     / \  / \
		//    1  3

		// Expect(1).To(Equal(2))
	})

	// It("Inserts value in the correct position when received in disorder", func() {
	// 	// Expected tree:
	// 	//		  4
	// 	//       / \
	// 	//      2   7
	// 	//     / \  / \
	// 	//    1  3
	// 	tree := binaryTree.BinaryTree{}
	// 	// nodes
	// 	values := []int{4, 7, 2, 3, 1}
	// 	for _, value := range values {
	// 		tree.InsertNode(tree.Root, value)
	// 	}
	// 	Expect(tree.InOrderTraverse()).To(Equal([]int{4, 2, 7, 1, 3}))
	// })

	// It("In order transversal", func() {
	// 	// Given tree:
	// 	//		   1
	// 	//       /  \
	// 	//      2    3
	// 	//     / \  / \
	// 	//    4   5
	// 	tree := binaryTree.BinaryTree{}
	// 	// nodes
	// 	values := []int{1, 2, 3, 4, 5}
	// 	for _, value := range values {
	// 		tree.InsertNode(tree.Root, value)
	// 	}
	// 	tree.InOrderTraverse(tree.Root)
	// })

	It("In order transversal", func() {
		// Given tree:

		//       10
		//      /  \
		//     20   30
		//    / \     \
		//   40 50     60
		//  /
		// 70

		tree := binaryTree.BinaryTree{}
		tree.Root = &binaryTree.BinaryTreeNode{10, nil, nil}
		tree.Root.Left = &binaryTree.BinaryTreeNode{20, nil, nil}
		tree.Root.Right = &binaryTree.BinaryTreeNode{30, nil, nil}
		tree.Root.Right.Right = &binaryTree.BinaryTreeNode{60, nil, nil}
		tree.Root.Left.Left = &binaryTree.BinaryTreeNode{40, nil, nil}
		tree.Root.Left.Right = &binaryTree.BinaryTreeNode{50, nil, nil}
		tree.Root.Left.Left.Left = &binaryTree.BinaryTreeNode{70, nil, nil}
		tree.InOrderTraverse(tree.Root)

		// Add a good expect here
		Expect(1).To(Equal(2))
	})

})
