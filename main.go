package main

import (
  "tree"
  "fmt"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
  //Walk Left
  if t.Left != nil {
    Walk(t.Left, ch)
  }
  //Return Value
  ch <- t.Value
  //Walk Right
  if t.Right != nil {
    Walk(t.Right, ch)
  }
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
  //initialize channels
  b, c := make(chan int, 10), make(chan int, 10)

  //Walk both trees
  go Walk(t1, b)
  go Walk(t2, c)

  //test equality
  for i := 0; i < 10; i++ {
    if <-b != <-c {
      return false
    }
  }
  return true
}

func main() {
  //test Walk
  ch := make(chan int, 10)
  go Walk(tree.New(1), ch)
  for i := 0; i < 10; i++ {
    i := <-ch
    fmt.Println(i)
  }

  //test Same
  if Same(tree.New(1), tree.New(2)) {
    fmt.Println("True")
  } else {
    fmt.Println("False")
  }
}
