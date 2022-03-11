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

	Describe("The module is able to", func() {
		Context("calculate the height of a tree", func() {

			It("of complete binary tree", func() {
				// Given tree:
				//          50
				//        /    \
				//      20      70
				//     /  \    /   \
				//    10  30  60    80
				//   /
				//  5
				tree := binaryTree.BinaryTree{}
				tree.Root = &binaryTree.BinaryTreeNode{50, nil, nil}
				tree.Root.Left = &binaryTree.BinaryTreeNode{20, nil, nil}
				tree.Root.Right = &binaryTree.BinaryTreeNode{70, nil, nil}
				tree.Root.Left.Left = &binaryTree.BinaryTreeNode{10, nil, nil}
				tree.Root.Right.Left = &binaryTree.BinaryTreeNode{60, nil, nil}
				tree.Root.Right.Right = &binaryTree.BinaryTreeNode{80, nil, nil}
				tree.Root.Left.Left.Left = &binaryTree.BinaryTreeNode{5, nil, nil}
				height := tree.Height(tree.Root)
				Expect(height).To(Equal(4))
			})

			It("of perfect binary tree", func() {
				// Given tree:
				//         50
				//       /    \
				//     20      70
				//    /  \    /   \
				//   10  30  60    80
				tree := binaryTree.BinaryTree{}
				tree.Root = &binaryTree.BinaryTreeNode{50, nil, nil}
				tree.Root.Left = &binaryTree.BinaryTreeNode{20, nil, nil}
				tree.Root.Right = &binaryTree.BinaryTreeNode{70, nil, nil}
				tree.Root.Left.Left = &binaryTree.BinaryTreeNode{10, nil, nil}
				tree.Root.Right.Left = &binaryTree.BinaryTreeNode{60, nil, nil}
				tree.Root.Right.Right = &binaryTree.BinaryTreeNode{80, nil, nil}
				height := tree.Height(tree.Root)
				Expect(height).To(Equal(3))
			})

			It("of one node", func() {
				// Given tree:
				//       10
				tree := binaryTree.BinaryTree{}
				height := tree.Height(tree.Root)
				Expect(height).To(Equal(0))
			})

			It("of a semi-degenerate tree", func() {
				// Given tree:
				//        1
				//       /
				//      3
				//   	 \
				//   	  5
				// 		   \
				// 			7
				tree := binaryTree.BinaryTree{}
				tree.Root = &binaryTree.BinaryTreeNode{1, nil, nil}
				tree.Root.Left = &binaryTree.BinaryTreeNode{3, nil, nil}
				tree.Root.Left.Right = &binaryTree.BinaryTreeNode{5, nil, nil}
				tree.Root.Left.Right.Right = &binaryTree.BinaryTreeNode{7, nil, nil}
				Expect(tree.Height(tree.Root)).To(Equal(4))
			})

			It("of degenerate tree", func() {
				// Given tree:
				//       10
				//         \
				//          30
				//            \
				//             60
				//   			 \
				//       		  80
				tree := binaryTree.BinaryTree{}
				tree.Root = &binaryTree.BinaryTreeNode{10, nil, nil}
				tree.Root.Right = &binaryTree.BinaryTreeNode{30, nil, nil}
				tree.Root.Right.Right = &binaryTree.BinaryTreeNode{60, nil, nil}
				tree.Root.Right.Right.Right = &binaryTree.BinaryTreeNode{80, nil, nil}
				height := tree.Height(tree.Root)
				Expect(height).To(Equal(4))
			})
		})

		Context("perform in-order traversal array of a tree", func() {
			It("for perfect binary tree", func() {
				// Given tree:
				//         50
				//       /    \
				//     20      70
				//    /  \    /   \
				//   10  30  60    80
				tree := binaryTree.BinaryTree{}
				tree.Root = &binaryTree.BinaryTreeNode{50, nil, nil}
				tree.Root.Left = &binaryTree.BinaryTreeNode{20, nil, nil}
				tree.Root.Right = &binaryTree.BinaryTreeNode{70, nil, nil}
				tree.Root.Left.Right = &binaryTree.BinaryTreeNode{30, nil, nil}
				tree.Root.Left.Left = &binaryTree.BinaryTreeNode{10, nil, nil}
				tree.Root.Right.Left = &binaryTree.BinaryTreeNode{60, nil, nil}
				tree.Root.Right.Right = &binaryTree.BinaryTreeNode{80, nil, nil}
				res := []int{10, 20, 30, 50, 60, 70, 80}
				Expect(tree.InOrderTraverseArray(tree.Root)).To(Equal(res))
			})
		})

		It("for complete binary tree", func() {
			// Given tree:
			//       8
			//      /  \
			//     4    9
			//    / \
			//   3   6
			tree := binaryTree.BinaryTree{}
			tree.Root = &binaryTree.BinaryTreeNode{8, nil, nil}
			tree.Root.Left = &binaryTree.BinaryTreeNode{4, nil, nil}
			tree.Root.Right = &binaryTree.BinaryTreeNode{9, nil, nil}
			tree.Root.Left.Left = &binaryTree.BinaryTreeNode{3, nil, nil}
			tree.Root.Left.Right = &binaryTree.BinaryTreeNode{6, nil, nil}
			res := []int{3, 4, 6, 8, 9}
			traverse := tree.InOrderTraverseArray(tree.Root)
			Expect(traverse).To(Equal(res))
		})

		It("for degenerate tree", func() {
			// Given tree:
			//       1
			//        \
			//         3
			//   		\
			//   		 5
			// 			  \
			// 			   7
			tree := binaryTree.BinaryTree{}
			tree.Root = &binaryTree.BinaryTreeNode{1, nil, nil}
			tree.Root.Right = &binaryTree.BinaryTreeNode{3, nil, nil}
			tree.Root.Right.Right = &binaryTree.BinaryTreeNode{5, nil, nil}
			tree.Root.Right.Right.Right = &binaryTree.BinaryTreeNode{7, nil, nil}
			res := []int{1, 3, 5, 7}
			traverse := tree.InOrderTraverseArray(tree.Root)
			Expect(traverse).To(Equal(res))
		})

		It("for semi-degenerate tree", func() {
			// Given tree:
			//        1
			//       /
			//      3
			//   	 \
			//   	  5
			// 		   \
			// 			7
			tree := binaryTree.BinaryTree{}
			tree.Root = &binaryTree.BinaryTreeNode{1, nil, nil}
			tree.Root.Left = &binaryTree.BinaryTreeNode{3, nil, nil}
			tree.Root.Left.Right = &binaryTree.BinaryTreeNode{5, nil, nil}
			tree.Root.Left.Right.Right = &binaryTree.BinaryTreeNode{7, nil, nil}
			res := []int{3, 5, 7, 1}
			traverse := tree.InOrderTraverseArray(tree.Root)
			Expect(traverse).To(Equal(res))
		})
	})
})
