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

	It("Inserts value in the correct position", func() {
		// Expected tree:
		//		  61
		//       /  \
		//      46  66
		//     / \  / \
		//    43  n n  88
		//   /
		//  39
		tree := binaryTree.BinaryTree{}
		// nodes
		values := []int{61, 46, 66, 43, 39, 88}
		for _, value := range values {
			tree.InsertNode(tree.Root, value)
		}
		Expect(tree.InOrderTraverse()).To(Equal(values))
	})

	It("Inserts value in the correct position when received in disorder", func() {
		// Expected tree:
		//		  4
		//       / \
		//      2   7
		//     / \  / \
		//    1  3
		tree := binaryTree.BinaryTree{}
		// nodes
		values := []int{4, 7, 2, 3, 1}
		for _, value := range values {
			tree.InsertNode(tree.Root, value)
		}
		Expect(tree.InOrderTraverse()).To(Equal([]int{4, 2, 7, 1, 3}))
	})

	It("Inserts value in the correct position when received in disorder", func() {
		// Expected tree:
		//		  4
		//       / \
		//      2   7
		//     / \  / \
		//    1  3
		tree := binaryTree.BinaryTree{}
		// nodes
		values := []int{4, 7, 2, 3, 1}
		for _, value := range values {
			tree.InsertNode(tree.Root, value)
		}
		Expect(tree.InOrderTraverse()).To(Equal([]int{4, 2, 7, 1, 3}))
	})

})
