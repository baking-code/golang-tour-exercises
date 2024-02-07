package main

import (
	"fmt"
	"time"
)

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

// GODEBUG=gctrace=1 go run main.go
func main() {
	start := time.Now()
	numPeople := 10_000_00
	// pre allocating the slice/array makes gc faster
	// arr := [10_000_00]Person{}
	arr := make([]Person, numPeople)
	for i := 0; i < numPeople; i++ {
		arr[i] = Person{"john", "doe", i}
		// appending to a slice each time takes much longer
		// arr = append(arr, Person{"john", "doe", i})
	}
	fmt.Println("Took:", time.Since(start))
}
