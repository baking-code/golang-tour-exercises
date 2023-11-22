package main

import (
	"fmt"
	"slices"

	"golang.org/x/tour/tree"
)

func slicesEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	defer close(ch)
	var walkerFunc func(t *tree.Tree)
	walkerFunc = func(treeRef *tree.Tree) {
		ch <- treeRef.Value
		if treeRef.Left != nil {
			walkerFunc(treeRef.Left)
		}
		if treeRef.Right != nil {
			walkerFunc(treeRef.Right)
		}
	}
	walkerFunc(t)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)
	values1 := make([]int, 0)
	values2 := make([]int, 0)
	go Walk(t1, ch1)
	go Walk(t2, ch2)
	for {
		value1, ok1 := <-ch1
		values1 = append(values1, value1)
		value2, ok2 := <-ch2
		values2 = append(values2, value2)
		if !(ok1 || ok2) {
			break
		}
	}
	slices.Sort(values1)
	slices.Sort(values2)

	return slicesEqual(values1, values2)
}

func main() {
	fmt.Println(Same(tree.New(1), tree.New(1)))
	fmt.Println(Same(tree.New(1), tree.New(2)))
}
