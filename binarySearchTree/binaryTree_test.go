package binarySearchTree_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/dbgoytia/datastructures/binarySearchTree"
)

var _ = Describe("BinarySearchTree", func() {

	It("Inserting a value when root is nil becomes the root", func() {
		// Expected tree:
		//		  61
		// Basic tree
		tree := binarySearchTree.BinarySearchTree{}
		// A node
		node := binarySearchTree.BinarySearchTreeNode{
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
				tree := binarySearchTree.BinarySearchTree{}
				tree.Root = &binarySearchTree.BinarySearchTreeNode{50, nil, nil}
				tree.Root.Left = &binarySearchTree.BinarySearchTreeNode{20, nil, nil}
				tree.Root.Right = &binarySearchTree.BinarySearchTreeNode{70, nil, nil}
				tree.Root.Left.Left = &binarySearchTree.BinarySearchTreeNode{10, nil, nil}
				tree.Root.Right.Left = &binarySearchTree.BinarySearchTreeNode{60, nil, nil}
				tree.Root.Right.Right = &binarySearchTree.BinarySearchTreeNode{80, nil, nil}
				tree.Root.Left.Left.Left = &binarySearchTree.BinarySearchTreeNode{5, nil, nil}
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
				tree := binarySearchTree.BinarySearchTree{}
				tree.Root = &binarySearchTree.BinarySearchTreeNode{50, nil, nil}
				tree.Root.Left = &binarySearchTree.BinarySearchTreeNode{20, nil, nil}
				tree.Root.Right = &binarySearchTree.BinarySearchTreeNode{70, nil, nil}
				tree.Root.Left.Left = &binarySearchTree.BinarySearchTreeNode{10, nil, nil}
				tree.Root.Right.Left = &binarySearchTree.BinarySearchTreeNode{60, nil, nil}
				tree.Root.Right.Right = &binarySearchTree.BinarySearchTreeNode{80, nil, nil}
				height := tree.Height(tree.Root)
				Expect(height).To(Equal(3))
			})

			It("of one node", func() {
				// Given tree:
				//       10
				tree := binarySearchTree.BinarySearchTree{}
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
				tree := binarySearchTree.BinarySearchTree{}
				tree.Root = &binarySearchTree.BinarySearchTreeNode{1, nil, nil}
				tree.Root.Left = &binarySearchTree.BinarySearchTreeNode{3, nil, nil}
				tree.Root.Left.Right = &binarySearchTree.BinarySearchTreeNode{5, nil, nil}
				tree.Root.Left.Right.Right = &binarySearchTree.BinarySearchTreeNode{7, nil, nil}
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
				tree := binarySearchTree.BinarySearchTree{}
				tree.Root = &binarySearchTree.BinarySearchTreeNode{10, nil, nil}
				tree.Root.Right = &binarySearchTree.BinarySearchTreeNode{30, nil, nil}
				tree.Root.Right.Right = &binarySearchTree.BinarySearchTreeNode{60, nil, nil}
				tree.Root.Right.Right.Right = &binarySearchTree.BinarySearchTreeNode{80, nil, nil}
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
				tree := binarySearchTree.BinarySearchTree{}
				tree.Root = &binarySearchTree.BinarySearchTreeNode{50, nil, nil}
				tree.Root.Left = &binarySearchTree.BinarySearchTreeNode{20, nil, nil}
				tree.Root.Right = &binarySearchTree.BinarySearchTreeNode{70, nil, nil}
				tree.Root.Left.Right = &binarySearchTree.BinarySearchTreeNode{30, nil, nil}
				tree.Root.Left.Left = &binarySearchTree.BinarySearchTreeNode{10, nil, nil}
				tree.Root.Right.Left = &binarySearchTree.BinarySearchTreeNode{60, nil, nil}
				tree.Root.Right.Right = &binarySearchTree.BinarySearchTreeNode{80, nil, nil}
				res := []int{10, 20, 30, 50, 60, 70, 80}
				Expect(tree.InOrderTraverseArray(tree.Root)).To(Equal(res))
			})
			It("for complete binary tree", func() {
				// Given tree:
				//       8
				//      /  \
				//     4    9
				//    / \
				//   3   6
				tree := binarySearchTree.BinarySearchTree{}
				tree.Root = &binarySearchTree.BinarySearchTreeNode{8, nil, nil}
				tree.Root.Left = &binarySearchTree.BinarySearchTreeNode{4, nil, nil}
				tree.Root.Right = &binarySearchTree.BinarySearchTreeNode{9, nil, nil}
				tree.Root.Left.Left = &binarySearchTree.BinarySearchTreeNode{3, nil, nil}
				tree.Root.Left.Right = &binarySearchTree.BinarySearchTreeNode{6, nil, nil}
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
				tree := binarySearchTree.BinarySearchTree{}
				tree.Root = &binarySearchTree.BinarySearchTreeNode{1, nil, nil}
				tree.Root.Right = &binarySearchTree.BinarySearchTreeNode{3, nil, nil}
				tree.Root.Right.Right = &binarySearchTree.BinarySearchTreeNode{5, nil, nil}
				tree.Root.Right.Right.Right = &binarySearchTree.BinarySearchTreeNode{7, nil, nil}
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
				tree := binarySearchTree.BinarySearchTree{}
				tree.Root = &binarySearchTree.BinarySearchTreeNode{1, nil, nil}
				tree.Root.Left = &binarySearchTree.BinarySearchTreeNode{3, nil, nil}
				tree.Root.Left.Right = &binarySearchTree.BinarySearchTreeNode{5, nil, nil}
				tree.Root.Left.Right.Right = &binarySearchTree.BinarySearchTreeNode{7, nil, nil}
				res := []int{3, 5, 7, 1}
				traverse := tree.InOrderTraverseArray(tree.Root)
				Expect(traverse).To(Equal(res))
			})
		})

		Context("perform pre-order traversal array of a tree", func() {
			It("for perfect binary tree", func() {
				// Given tree:
				//         50
				//       /    \
				//     20      70
				//    /  \    /   \
				//   10  30  60    80
				tree := binarySearchTree.BinarySearchTree{}
				tree.Root = &binarySearchTree.BinarySearchTreeNode{50, nil, nil}
				tree.Root.Left = &binarySearchTree.BinarySearchTreeNode{20, nil, nil}
				tree.Root.Right = &binarySearchTree.BinarySearchTreeNode{70, nil, nil}
				tree.Root.Left.Right = &binarySearchTree.BinarySearchTreeNode{30, nil, nil}
				tree.Root.Left.Left = &binarySearchTree.BinarySearchTreeNode{10, nil, nil}
				tree.Root.Right.Left = &binarySearchTree.BinarySearchTreeNode{60, nil, nil}
				tree.Root.Right.Right = &binarySearchTree.BinarySearchTreeNode{80, nil, nil}
				res := []int{50, 20, 10, 30, 70, 60, 80}
				Expect(tree.PreOrderTraverseArray(tree.Root)).To(Equal(res))
			})
			It("for complete binary tree", func() {
				// Given tree:
				//       8
				//      /  \
				//     4    9
				//    / \
				//   3   6
				tree := binarySearchTree.BinarySearchTree{}
				tree.Root = &binarySearchTree.BinarySearchTreeNode{8, nil, nil}
				tree.Root.Left = &binarySearchTree.BinarySearchTreeNode{4, nil, nil}
				tree.Root.Right = &binarySearchTree.BinarySearchTreeNode{9, nil, nil}
				tree.Root.Left.Left = &binarySearchTree.BinarySearchTreeNode{3, nil, nil}
				tree.Root.Left.Right = &binarySearchTree.BinarySearchTreeNode{6, nil, nil}
				res := []int{8, 4, 3, 6, 9}
				Expect(tree.PreOrderTraverseArray(tree.Root)).To(Equal(res))
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
				tree := binarySearchTree.BinarySearchTree{}
				tree.Root = &binarySearchTree.BinarySearchTreeNode{1, nil, nil}
				tree.Root.Right = &binarySearchTree.BinarySearchTreeNode{3, nil, nil}
				tree.Root.Right.Right = &binarySearchTree.BinarySearchTreeNode{5, nil, nil}
				tree.Root.Right.Right.Right = &binarySearchTree.BinarySearchTreeNode{7, nil, nil}
				res := []int{1, 3, 5, 7}
				Expect(tree.PreOrderTraverseArray(tree.Root)).To(Equal(res))
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
				tree := binarySearchTree.BinarySearchTree{}
				tree.Root = &binarySearchTree.BinarySearchTreeNode{1, nil, nil}
				tree.Root.Left = &binarySearchTree.BinarySearchTreeNode{3, nil, nil}
				tree.Root.Left.Right = &binarySearchTree.BinarySearchTreeNode{5, nil, nil}
				tree.Root.Left.Right.Right = &binarySearchTree.BinarySearchTreeNode{7, nil, nil}
				res := []int{3, 5, 7, 1}
				traverse := tree.InOrderTraverseArray(tree.Root)
				Expect(traverse).To(Equal(res))
			})
		})

		Context("perform bread-first traversal array of a tree", func() {

			It("for perfect binary tree", func() {
				// Given tree:
				//         50
				//       /    \
				//     20      70
				//    /  \    /   \
				//   10  30  60    80
				tree := binarySearchTree.BinarySearchTree{}
				tree.Root = &binarySearchTree.BinarySearchTreeNode{50, nil, nil}
				tree.Root.Left = &binarySearchTree.BinarySearchTreeNode{20, nil, nil}
				tree.Root.Right = &binarySearchTree.BinarySearchTreeNode{70, nil, nil}
				tree.Root.Left.Right = &binarySearchTree.BinarySearchTreeNode{30, nil, nil}
				tree.Root.Left.Left = &binarySearchTree.BinarySearchTreeNode{10, nil, nil}
				tree.Root.Right.Left = &binarySearchTree.BinarySearchTreeNode{60, nil, nil}
				tree.Root.Right.Right = &binarySearchTree.BinarySearchTreeNode{80, nil, nil}
				res := []int{50, 20, 70, 10, 30, 60, 80}
				Expect(tree.BreadFirstArray(tree.Root)).To(Equal(res))
			})

			It("for complete binary tree", func() {
				// Given tree:
				//       8
				//      /  \
				//     4    9
				//    / \
				//   3   6
				tree := binarySearchTree.BinarySearchTree{}
				tree.Root = &binarySearchTree.BinarySearchTreeNode{8, nil, nil}
				tree.Root.Left = &binarySearchTree.BinarySearchTreeNode{4, nil, nil}
				tree.Root.Right = &binarySearchTree.BinarySearchTreeNode{9, nil, nil}
				tree.Root.Left.Left = &binarySearchTree.BinarySearchTreeNode{3, nil, nil}
				tree.Root.Left.Right = &binarySearchTree.BinarySearchTreeNode{6, nil, nil}
				res := []int{8, 4, 9, 3, 6}
				Expect(tree.BreadFirstArray(tree.Root)).To(Equal(res))
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
				tree := binarySearchTree.BinarySearchTree{}
				tree.Root = &binarySearchTree.BinarySearchTreeNode{1, nil, nil}
				tree.Root.Right = &binarySearchTree.BinarySearchTreeNode{3, nil, nil}
				tree.Root.Right.Right = &binarySearchTree.BinarySearchTreeNode{5, nil, nil}
				tree.Root.Right.Right.Right = &binarySearchTree.BinarySearchTreeNode{7, nil, nil}
				res := []int{1, 3, 5, 7}
				Expect(tree.BreadFirstArray(tree.Root)).To(Equal(res))
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
				tree := binarySearchTree.BinarySearchTree{}
				tree.Root = &binarySearchTree.BinarySearchTreeNode{1, nil, nil}
				tree.Root.Left = &binarySearchTree.BinarySearchTreeNode{3, nil, nil}
				tree.Root.Left.Right = &binarySearchTree.BinarySearchTreeNode{5, nil, nil}
				tree.Root.Left.Right.Right = &binarySearchTree.BinarySearchTreeNode{7, nil, nil}
				res := []int{1, 3, 5, 7}
				traverse := tree.BreadFirstArray(tree.Root)
				Expect(traverse).To(Equal(res))
			})

		})

	})
})
