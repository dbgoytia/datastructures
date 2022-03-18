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

Trasversals of a given tree are one of the most important topics to understand in order to move forward to more complex data structures, like graphs of heaps, so let's spend some time in each of them.


## In-order traversal

Algorithm Inorder(tree)
1. Traverse the left subtree, i.e., call Inorder(left-subtree)
2. Visit the root.
3. Traverse the right subtree, i.e., call Inorder(right-subtree)


### Applications
* Retrieve the sorted contents of a Binary Search Tree (ascending orders)

```
 Given tree:

         50
       /    \
     20      70
    /  \    /   \
   10  30  60    80

Represented as:

[10, 20, 30, 50, 60, 70, 80] <--- This is called in order traversal

```

## Pre-order traversal - Also called Depth-First Search

Algorithm PreOrder(tree) (Depth-First Search)
2. Visit the root.
1. Traverse the left subtree, i.e., call PreOrder(left-subtree)
3. Traverse the right subtree, i.e., call PreOrder(right-subtree)

### Applications

* You'll notice that you'll start descending the tree with the **left-most child first**
* You can see that this way it's super easy to create a copy of the tree.
* Databases often perform pre-order traversal to traverse B-tree indexes during search operations
* The Linux Makefile utility also makes use of this algorithm.

```
 Given tree:

         50
       /    \
     20      70
    /  \    /   \
   10  30  60    80

Represented as:

[50, 20, 10, 30, 70, 60, 80] <--- This is called in depth-first traversal

```

## Post-order traversal 

Algorithm Post order(tree)
1. Traverse the left subtree, i.e., call Inorder(left-subtree)
2. Visit the root.
3. Traverse the right subtree, i.e., call Inorder(right-subtree)

###  Aplications

* Super useful for deleting a tree, because it goes from the buttom-up.

```
 Given tree:

         50
       /    \
     20      70
    /  \    /   \
   10  30  60    80

Represented as:

[10, 30, 20, 60, 80, 70, 50] <--- This is called in post order traversal

```


## Breadth-First Search or Level Order

```
Algorithm (Recursive):
-  for every level of a tree:
    - printLevel

printLevel:
- if tree is Nil return

- if level is 1
    - print data

- else
    - printLevel(left, level -1);
    - printLevel(left, level -1);
```

```
Algorithm (Iterative) with a queue: # This is a ton important to learn graphs

- Create a queue
- mark the node as discovered
- enqueue node
- while the queue has something (not empty):
    - dequeue front node and print it
    - for every sub-node
        - mark as discovered
        - enqueue it

Sample run:

 Given tree:

         50
       /    \
     20      70
    /  \    /   \
   10  30  60    80

1. Q = [] # Queue is representedas as a list
2. Discovered = [False, False, False, False, False, False, False] # 2 ** 3 - 1 = 7. You'll create as much false as possible nodes in the Tree.
4. D = True # Mark Node as discovered

5. while the queue has elements:
  
  
  6. print(50) and dequeue                    # Q = [50]
  7. for 20, 70 ,mark as discovered
  8. append 20, 7                             # Q = [20, 70]

  9. print(20, 70) and dequeue                # Q = []
  10. for 10, 30, 60, 80
  11. append them                             # Q = [10, 30, 60, 80]

  12. print(10, 30, 80, 80) and dequeue       # Q = []
  13. since there are no adjacent nodes...
  14. Queue no longer has any elements.


Result: [50, 20, 70, 10, 30, 60, 80]


```

###  Aplications
* It's a popular search algorithm in graphs as well.

```
 Given tree:

         50
       /    \
     20      70
    /  \    /   \
   10  30  60    80

Represented as:

[50, 20, 70, 10, 30, 60, 80] <--- This is called breadth-first search
```





---------------------------------------------------------------







# Tree array representation

When representing a tree as an array or a list, it's important to keep two things:

1. The value of the node
2. The relationship of the node (parents, and childs)

Example:

```
 Given tree:

         50
       /    \
     20      70
    /  \    /   \
   10  30  60    80

Represented as:

[50, 20, 70, 10, 30, 60, 80]

```

To calculate the position of a given node in *i*,
* The left child is at 2 * *i*
* The right child is at 2 * *i* + 1
* The parent is at *i*/2

Also notice that the representation of the array is in **Level Order** also known as **Breadth First Search**.


```
Demonstration:

[50, 20, 70, 10, 30, 60, 80]   <--- Data
  1 , 2 , 3 , 4 , 5 , 6 , 7    <--- Positions

 Given node in position i = 3:
 - Left child is in 3 * 2 = 6
 - Right child is in 3 * 2 + 1 = 7 
 - Parent is in 3 / 2 = 1  
```


```
Demonstration:

[50, 20, 70, 10, 30, 60, 80]   <--- Data
 0,  1 ,  2,  3,  4,  5,  6    <--- Positions

 Given node in position i = 3:
 - Left child is in 3 * 2 = 6
 - Right child is in 3 * 2 + 1 = 7 
 - Parent is in 3 / 2 = 1  
```


###  Example two

```
 Given tree where you're missing pieces, the property will
 still add up.

         50
       /    \
     20      70
    /  \     
   10  30     

Represented as:

[50, 20, 70, 10, 30]
 1 ,  2,  3,  4,  5

 For a node in position 4
 * parent is at 4/2 = 2
 * left child is at 4 * 2 = 8 -> doesn't exist
 * right child is at 4 * 2 + 1 =  9 -> doesn't exist.

```


## Example three

This one is the **most** important one to see because it's a basis for understanding heaps, see how the tree would be represented

```
 Given tree:

          50
        /    \
       20    70
            /   \
           60    80
Represented as:

[50, 20, 70, -, -, 60, 80] <--- Notice that you need to keep track of missing data.


for a given node in i = 2
* parent is at 2/2 = 1
* left child is at 2 * 2 = 4 -> doesn't exist
* right child isat 2 * 2 + 1 = 5 -> doesn't exist
```







---

## Sources

* https://www.baeldung.com/cs/depth-first-traversal-methods
* https://www.youtube.com/watch?v=HqPJF2L5h9U&ab_channel=AbdulBari
* https://www.geeksforgeeks.org/binary-tree-array-implementation/
* https://dev.to/abdisalan_js/4-ways-to-traverse-binary-trees-with-animations-5bi5
* https://www.techiedelight.com/breadth-first-search/





----
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
``