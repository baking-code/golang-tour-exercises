package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	// init slice with 0 and 1 (first values)
	// (this is a closure belonging to the instance of this function)
	var tuple = [2]int{0, 1}
	return func() int {
		f := tuple[0]
		// next value in fibonacci is the previous two values summed
		// we keep the current and the next value in memory
		tuple[0], tuple[1] = tuple[1], tuple[1]+f
		// return the new current value
		return f
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
