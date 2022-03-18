package binarySearchTree_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/dbgoytia/datastructures/trees/binarySearchTree"
)

var _ = Describe("BinarySearchTree", func() {

	Describe("The module is able to", func() {

		Context("insert the node in the correct position", func() {

			It("Inserting a value when root is nil becomes the root", func() {
				// Expected tree:
				//		  61
				// Basic tree
				tree := binarySearchTree.BinarySearchTree{}
				// A node
				tree.InsertNode(61)
				Expect(tree.Root.Val).To(Equal(61))
			})

			It("of perfect binary tree", func() {
				// Given tree:
				//         50
				//       /    \
				//     20      70
				//    /  \    /   \
				//   10  30  60    80
				tree := binarySearchTree.BinarySearchTree{}
				tree.InsertNode(50)
				tree.InsertNode(20)
				tree.InsertNode(70)
				tree.InsertNode(10)
				tree.InsertNode(60)
				tree.InsertNode(80)
				Expect(tree.Root.Right.Right.Val).To(Equal(80))
			})
		})

		Context("removes values", func() {

			It("on leaf values", func() {
				// Given tree:
				//         50
				//       /    \
				//     20      70
				//    /  \    /   \
				//   10  30  60    80
				tree := binarySearchTree.BinarySearchTree{}
				tree.InsertNode(50)
				tree.InsertNode(20)
				tree.InsertNode(70)
				tree.InsertNode(10)
				tree.InsertNode(60)
				tree.Remove(80)
				Expect(tree.Root.Right.Right).To(Equal(nil))
			})

			It("on nodes", func() {
				// Given tree:
				//         50
				//       /    \
				//     20      70
				//    /  \    /   \
				//   10  30  60    80
				tree := binarySearchTree.BinarySearchTree{}
				tree.InsertNode(50)
				tree.InsertNode(20)
				tree.InsertNode(70)
				tree.InsertNode(10)
				tree.InsertNode(60)
				tree.Remove(20)
				Expect(tree.Root.Left).To(Equal(10))
			})
		})

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
				tree.InsertNode(50)
				tree.InsertNode(20)
				tree.InsertNode(70)
				tree.InsertNode(10)
				tree.InsertNode(10)
				tree.InsertNode(60)
				tree.InsertNode(80)
				tree.InsertNode(5)
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
				tree.InsertNode(50)
				tree.InsertNode(20)
				tree.InsertNode(70)
				tree.InsertNode(10)
				tree.InsertNode(10)
				tree.InsertNode(60)
				tree.InsertNode(80)
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
				tree.InsertNode(1)
				tree.InsertNode(3)
				tree.InsertNode(5)
				tree.InsertNode(7)
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
				tree.InsertNode(10)
				tree.InsertNode(30)
				tree.InsertNode(60)
				tree.InsertNode(80)
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
				tree.InsertNode(50)
				tree.InsertNode(20)
				tree.InsertNode(70)
				tree.InsertNode(30)
				tree.InsertNode(10)
				tree.InsertNode(60)
				tree.InsertNode(80)
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
				tree.InsertNode(8)
				tree.InsertNode(4)
				tree.InsertNode(9)
				tree.InsertNode(3)
				tree.InsertNode(6)
				tree.InsertNode(6)
				tree.InsertNode(6)
				tree.InsertNode(6)
				tree.InsertNode(6)
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
				tree.InsertNode(1)
				tree.InsertNode(3)
				tree.InsertNode(5)
				tree.InsertNode(7)
				res := []int{1, 3, 5, 7}
				traverse := tree.InOrderTraverseArray(tree.Root)
				Expect(traverse).To(Equal(res))
			})

			It("for semi-degenerate tree", func() {
				// Given tree:
				//        8
				//       /
				//      3
				//   	 \
				//   	  5
				// 		   \
				// 			7
				tree := binarySearchTree.BinarySearchTree{}
				tree.InsertNode(8)
				tree.InsertNode(3)
				tree.InsertNode(5)
				tree.InsertNode(7)
				res := []int{3, 5, 7, 8}
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
				tree.InsertNode(50)
				tree.InsertNode(20)
				tree.InsertNode(70)
				tree.InsertNode(30)
				tree.InsertNode(10)
				tree.InsertNode(60)
				tree.InsertNode(80)
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
				tree.InsertNode(8)
				tree.InsertNode(4)
				tree.InsertNode(9)
				tree.InsertNode(3)
				tree.InsertNode(6)
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
				tree.InsertNode(1)
				tree.InsertNode(3)
				tree.InsertNode(5)
				tree.InsertNode(7)
				res := []int{1, 3, 5, 7}
				Expect(tree.PreOrderTraverseArray(tree.Root)).To(Equal(res))
			})

			It("for semi-degenerate tree", func() {
				// Given tree:
				//        8
				//       /
				//      3
				//   	 \
				//   	  5
				// 		   \
				// 			7
				tree := binarySearchTree.BinarySearchTree{}
				tree.InsertNode(8)
				tree.InsertNode(3)
				tree.InsertNode(5)
				tree.InsertNode(7)
				res := []int{3, 5, 7, 8}
				traverse := tree.InOrderTraverseArray(tree.Root)
				Expect(traverse).To(Equal(res))
			})
		})

		Context("perform post-order traversal array of a tree", func() {
			It("for perfect binary tree", func() {
				// Given tree:
				//         50
				//       /    \
				//     20      70
				//    /  \    /   \
				//   10  30  60    80
				tree := binarySearchTree.BinarySearchTree{}
				tree.InsertNode(50)
				tree.InsertNode(20)
				tree.InsertNode(70)
				tree.InsertNode(30)
				tree.InsertNode(10)
				tree.InsertNode(60)
				tree.InsertNode(80)
				res := []int{10, 30, 20, 60, 80, 70, 50}
				Expect(tree.PostOrderTraverseArray(tree.Root)).To(Equal(res))
			})

			It("for complete binary tree", func() {
				// Given tree:
				//       8
				//      /  \
				//     4    9
				//    / \
				//   3   6
				tree := binarySearchTree.BinarySearchTree{}
				tree.InsertNode(8)
				tree.InsertNode(4)
				tree.InsertNode(9)
				tree.InsertNode(3)
				tree.InsertNode(6)
				res := []int{3, 6, 4, 9, 8}
				Expect(tree.PostOrderTraverseArray(tree.Root)).To(Equal(res))
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
				tree.InsertNode(1)
				tree.InsertNode(3)
				tree.InsertNode(5)
				tree.InsertNode(7)
				res := []int{7, 5, 3, 1}
				Expect(tree.PostOrderTraverseArray(tree.Root)).To(Equal(res))
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
				tree.InsertNode(1)
				tree.InsertNode(3)
				tree.InsertNode(5)
				tree.InsertNode(7)
				res := []int{7, 5, 3, 1}
				traverse := tree.PostOrderTraverseArray(tree.Root)
				Expect(traverse).To(Equal(res))
			})
		})

		Context("perform breadth-first traversal array of a tree", func() {

			It("for perfect binary tree", func() {
				// Given tree:
				//         50
				//       /    \
				//     20      70
				//    /  \    /   \
				//   10  30  60    80
				tree := binarySearchTree.BinarySearchTree{}
				tree.InsertNode(50)
				tree.InsertNode(20)
				tree.InsertNode(70)
				tree.InsertNode(30)
				tree.InsertNode(10)
				tree.InsertNode(60)
				tree.InsertNode(80)
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
				tree.InsertNode(8)
				tree.InsertNode(4)
				tree.InsertNode(9)
				tree.InsertNode(3)
				tree.InsertNode(6)
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
				tree.InsertNode(1)
				tree.InsertNode(3)
				tree.InsertNode(5)
				tree.InsertNode(7)
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
				tree.InsertNode(1)
				tree.InsertNode(3)
				tree.InsertNode(5)
				tree.InsertNode(7)
				res := []int{1, 3, 5, 7}
				traverse := tree.BreadFirstArray(tree.Root)
				Expect(traverse).To(Equal(res))
			})

		})

		Context("calculates Max value in a tree", func() {

			It("for perfect binary tree", func() {
				// Given tree:
				//         50
				//       /    \
				//     20      70
				//    /  \    /   \
				//   10  30  60    80
				tree := binarySearchTree.BinarySearchTree{}
				tree.InsertNode(50)
				tree.InsertNode(20)
				tree.InsertNode(70)
				tree.InsertNode(30)
				tree.InsertNode(10)
				tree.InsertNode(60)
				tree.InsertNode(80)
				Expect(tree.Max(tree.Root)).To(Equal(80))
			})

			It("for complete binary tree", func() {
				// Given tree:
				//       8
				//      /  \
				//     4    9
				//    / \
				//   3   6
				tree := binarySearchTree.BinarySearchTree{}
				tree.InsertNode(8)
				tree.InsertNode(4)
				tree.InsertNode(9)
				tree.InsertNode(3)
				tree.InsertNode(6)
				Expect(tree.Max(tree.Root)).To(Equal(9))
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
				tree.InsertNode(1)
				tree.InsertNode(3)
				tree.InsertNode(5)
				tree.InsertNode(7)
				Expect(tree.Max(tree.Root)).To(Equal(7))
			})

			It("for semi-degenerate tree", func() {
				// Given tree:
				//        8
				//       /
				//      3
				//   	 \
				//   	  5
				// 		   \
				// 			7
				tree := binarySearchTree.BinarySearchTree{}
				tree.Root = &binarySearchTree.BinarySearchTreeNode{12, nil, nil}
				tree.Root.Left = &binarySearchTree.BinarySearchTreeNode{3, nil, nil}
				tree.Root.Left.Right = &binarySearchTree.BinarySearchTreeNode{5, nil, nil}
				tree.Root.Left.Right.Right = &binarySearchTree.BinarySearchTreeNode{7, nil, nil}
				Expect(tree.Max(tree.Root)).To(Equal(12))
			})

		})

		Context("performs binary search", func() {

			It("on perfect binary tree", func() {
				// Given tree:
				//         50
				//       /    \
				//     20      70
				//    /  \    /   \
				//   10  30  60    80
				tree := binarySearchTree.BinarySearchTree{}
				tree.InsertNode(50)
				tree.InsertNode(20)
				tree.InsertNode(70)
				tree.InsertNode(30)
				tree.InsertNode(10)
				tree.InsertNode(60)
				tree.InsertNode(80)
				Expect(tree.BinarySearch(tree.Root, 60)).To(Equal(true))
			})

			It("on complete binary tree", func() {
				// Given tree:
				//       8
				//      /  \
				//     4    9
				//    / \
				//   3   6
				tree := binarySearchTree.BinarySearchTree{}
				tree.InsertNode(8)
				tree.InsertNode(4)
				tree.InsertNode(9)
				tree.InsertNode(3)
				tree.InsertNode(6)
				Expect(tree.BinarySearch(tree.Root, 8)).To(Equal(true))
			})

			It("on degenerate tree", func() {
				// Given tree:
				//       1
				//        \
				//         3
				//   		\
				//   		 5
				// 			  \
				// 			   7
				tree := binarySearchTree.BinarySearchTree{}
				tree.InsertNode(1)
				tree.InsertNode(3)
				tree.InsertNode(5)
				tree.InsertNode(7)
				Expect(tree.BinarySearch(tree.Root, 22)).To(Equal(false))
			})

			It("on semi-degenerate tree", func() {
				// Given tree:
				//        12
				//       /
				//      3
				//   	 \
				//   	  5
				// 		   \
				// 			7
				tree := binarySearchTree.BinarySearchTree{}
				tree.InsertNode(12)
				tree.InsertNode(3)
				tree.InsertNode(5)
				tree.InsertNode(7)
				Expect(tree.BinarySearch(tree.Root, 33)).To(Equal(false))
			})

			It("on single-node tree", func() {
				// Given tree:
				//        12
				tree := binarySearchTree.BinarySearchTree{}
				tree.InsertNode(12)
				Expect(tree.BinarySearch(tree.Root, 12)).To(Equal(true))
			})

		})

		Context("perform breadth-first traversal array of a tree - iterative", func() {

			It("for perfect binary tree", func() {
				// Given tree:
				//         50
				//       /    \
				//     20      70
				//    /  \    /   \
				//   10  30  60    80
				tree := binarySearchTree.BinarySearchTree{}
				tree.InsertNode(50)
				tree.InsertNode(20)
				tree.InsertNode(70)
				tree.InsertNode(30)
				tree.InsertNode(10)
				tree.InsertNode(60)
				tree.InsertNode(80)
				res := []int{50, 20, 70, 10, 30, 60, 80}
				Expect(tree.BFTIterative()).To(Equal(res))
			})

			It("for complete binary tree", func() {
				// Given tree:
				//       8
				//      /  \
				//     4    9
				//    / \
				//   3   6
				tree := binarySearchTree.BinarySearchTree{}
				tree.InsertNode(8)
				tree.InsertNode(4)
				tree.InsertNode(9)
				tree.InsertNode(3)
				tree.InsertNode(6)
				res := []int{8, 4, 9, 3, 6}
				Expect(tree.BFTIterative()).To(Equal(res))
			})

			It("on degenerate tree", func() {
				// Given tree:
				//       1
				//        \
				//         3
				//   		\
				//   		 5
				// 			  \
				// 			   7
				tree := binarySearchTree.BinarySearchTree{}
				tree.InsertNode(1)
				tree.InsertNode(3)
				tree.InsertNode(5)
				tree.InsertNode(7)
				res := []int{1, 3, 5, 7}
				Expect(tree.BFTIterative()).To(Equal(res))
			})

			It("on semi-degenerate tree", func() {
				// Given tree:
				//       1
				//        \
				//         3
				//   		\
				//   		 5
				// 			  \
				// 			   7
				tree := binarySearchTree.BinarySearchTree{}
				tree.InsertNode(1)
				tree.InsertNode(3)
				tree.InsertNode(5)
				tree.InsertNode(7)
				res := []int{1, 3, 5, 7}
				Expect(tree.BFTIterative()).To(Equal(res))
			})
		})

	})
})
