package main

import "fmt"

// List represents a singly-linked list that holds
// values of any type.
type List[T any] struct {
	next  *List[T]
	value T
}

func (list *List[T]) String() string {
	// assign temp variable
	temp := list
	isFirst := true
	for temp != nil {
		// print each value until we reach the end of the list
		if isFirst {
			fmt.Printf("%v ", temp.value)
			isFirst = false
		} else {
			fmt.Printf("-> %v ", temp.value)
		}
		temp = temp.next
	}
	return "."
}

func (*List[T]) add(list *List[T], element T) *List[T] {
	// init first (or next link)
	temp := &List[T]{value: element, next: nil}
	if list == nil {
		// if this is the first item
		list = temp
	} else {
		pt := list
		// else iterate through the list until we're at the end, then
		// add the next one
		for pt.next != nil {
			pt = pt.next
		}
		pt.next = temp
	}
	return list
}

func main() {
	ll := List[int]{}
	ll.add(&ll, 1)
	ll.add(&ll, 2)
	fmt.Print(&ll)
}
