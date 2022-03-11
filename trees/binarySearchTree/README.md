# Binary tree

The most basic implementations of a Binary tree are:

* Perfect binary tree
* Complete binary tree
* Degenerate binary tree


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