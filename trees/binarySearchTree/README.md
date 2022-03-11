# Binary tree

The most basic implementations of a Binary tree are:

* Perfect binary tree
* Complete binary tree
* Degenerate binary tree


## Note

In this implementation we're not going to be dealing with generics, right now,
this binary search tree is only going to deal with integer values for learning
purposes. Later on, I might add some datastructure that is able to handle generics.

This datastructure is also not thread-safe, I might also implement a thread-safe
version of this data structure.

# Traversing the tree
## In-order traversal

Algorithm Inorder(tree)
1. Traverse the left subtree, i.e., call Inorder(left-subtree)
2. Visit the root.
3. Traverse the right subtree, i.e., call Inorder(right-subtree)

### Applications 
* Retrieve the sorted contents of a Binary Search Tree (ascending orders)

## Pre-order traversal

Algorithm PreOrder(tree)
2. Visit the root.
1. Traverse the left subtree, i.e., call PreOrder(left-subtree)
3. Traverse the right subtree, i.e., call PreOrder(right-subtree)

## Applications

* Databases often perform pre-order traversal to traverse B-tree indexes during search operations
* Useful when searching for an element within a binary search tree
* The Linux Makefile utility.¡


## Sources

* https://www.baeldung.com/cs/depth-first-traversal-methods


## Tests
```
Summarizing 2 Failures:

[Fail] BinarySearchTree The module is able to removes values [It] on leaf values 
/Users/diego_canizales/go/src/github.com/dbgoytia/datastructures/trees/binarySearchTree/binaryTree_test.go:60

[Fail] BinarySearchTree The module is able to removes values [It] on nodes 
/Users/diego_canizales/go/src/github.com/dbgoytia/datastructures/trees/binarySearchTree/binaryTree_test.go:77

Ran 34 of 34 Specs in 0.004 seconds
FAIL! -- 32 Passed | 2 Failed | 0 Pending | 0 Skipped
--- FAIL: TestBinaryTree (0.01s)
FAIL

Ginkgo ran 1 suite in 1.611383729s
Test Suite Failed
```