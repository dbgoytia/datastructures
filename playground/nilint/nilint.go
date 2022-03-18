package main

import (
	"fmt"
)

type NilInt struct {
	value int
	null  bool
}

func (n *NilInt) Value() interface{} {
	if n.null {
		return nil
	}
	return n.value
}

func NewInt(x int) NilInt {
	return NilInt{x, false}
}

func NewNil() NilInt {
	return NilInt{0, true}
}

func main() {
	var x = []NilInt{NewNil(), NewInt(10), NewNil(), NewInt(5)}
	for _, i := range x {
		fmt.Printf("%v, ", i.Value())
	}
}
