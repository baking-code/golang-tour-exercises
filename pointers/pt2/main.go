package main

import "fmt"

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func UpdateSlice(in []string, val string) {
	in[len(in)-1] = val
	fmt.Println("After UpdateSlice: ", in)
}

func GrowSlice(in []string, val string) {
	out := append(in, val)
	fmt.Println("After GrowSlice: ", out)
}

func main() {
	input := []string{"one", "two"}
	fmt.Println("Before UpdateSlice: ", input)
	UpdateSlice(input, "three")
	fmt.Println("Before GrowSlice: ", input)
	GrowSlice(input, "four")
	fmt.Println("End: ", input)

	/*
	   Prints
	   Before UpdateSlice:  [one two]
	   After UpdateSlice:  [one three]
	   Before GrowSlice:  [one three]
	   After GrowSlice:  [one three four]
	   End:  [one three]

	   we can modify the original slice in UpdateSlice because the slice remains of a constant length, Grow slice appends, so refers to a new pointer
	*/
}
