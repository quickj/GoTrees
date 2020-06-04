package main

import (
	"tree";
	"fmt"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	if t != nil {
		Walk(t.Left, ch)
		ch <- t.Value
		Walk(t.Right, ch)
	}
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	b,c := make(chan int), make(chan int)
	
	go Walk(t1, b)
	go Walk(t2, c)
	
	for i := 0; i < 10; i++ {
		if <-b != <-c {
			return false
		}
	}
	return true
}

func main() {
	t := tree.New(1)
	ch := make(chan int)
	go Walk(t, ch)
	for i := 0; i < 10; i++ {
		fmt.Println(<- ch)
	}
	fmt.Println(Same(tree.New(1), tree.New(1)))
	fmt.Println(Same(tree.New(1), tree.New(2)))
}
