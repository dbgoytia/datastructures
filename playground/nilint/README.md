## Nil Integer type

This come up to me as a question as I was building the "level order"
transversal of a Binary Tree. In my opinon, we should be able to track
missing values inside a tree so that the array representation that 
we get from it is an accurate representation of the tree -> at least in
my version of the tree that currently only supports integers.

It will also make the LevelOrder representation complete with the mathematical functions:

```
for a given node in i:

* the parent lives at i/2
* the left child lives at i*2
* the right child lives at (i*2)+1
```

As Go is a Statically typed language, and an slice of integers
will only allow that specific type to be included in the slice, I'm
thinking I could use this representation of a Nil integer in the 
future to improve that solution.


Even if my implementation of the tree has changed by the team you read this, I still found it interesting to add it to the datastructures training.

# Implementation

Might seem kind of obvious to most of you but I wanted to highlight some of
the best features of this very simple implementation:


## The NilInt struct

If you can see the Nil int struct has two attributes:
* holds a value of type int
* holds a null of type bool, that simply says true when you're referring to a nil

## The Value() method

Now this is where it gets very interesting, as you can see, the return value of this function is an empty **interface{}**.

# The laws of Reflection, main notes:

* Reflection in computing is the ability of a program to examine it's own
structure, particulary thorugh types.
* Every language reflection model is different, and many languages don't even support it at all.

There are two representations of an **empty interface** in go:

```
interface{}
```

or it's equivalent alias

```
any
```

This blog post features three different reflection rules, which I encourage you to read as well, for this scenario, I think we can get an idea of Rule #2

## Rule 2 Like physicall reflection, reflection in Go generates its own inverse.

Given a reflect.Value we can recover an interface value using the Interface Method (in our case, the Value() method). This method itself will pack the type and the value information back into an interface representation to return a result.

In our implementation, we also see that the arguments to **fmt.Println** and **fmt.Printf** are all passed as empty interfaces, which are then *unpacked* internally by the fmt package, that's why it only takes a call to either of those functions to get the value correctly



## Source:
* https://go.dev/blog/laws-of-reflection