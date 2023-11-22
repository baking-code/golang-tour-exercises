package main

import "fmt"

// List represents a doubly-linked list that holds
// values of any type.
type List[T any] struct {
	next  *List[T]
	prev  *List[T]
	value T
}

func (list *List[T]) String() string {
	// assign temp variable
	temp := list
	temp2 := list
	isFirst := true
	str := ""
	for temp != nil {
		// print each value until we reach the end of the list
		if isFirst {
			str = str + fmt.Sprintf("%v ", temp.value)
			isFirst = false
		} else {
			str = str + fmt.Sprintf("-> %v ", temp.value)
		}
		temp2 = temp
		temp = temp.next
	}
	for temp2 != nil {
		// print each value until we reach the start of the list
		str = str + fmt.Sprintf("<- %v ", temp2.value)
		temp2 = temp2.prev
	}
	return str

}

func (*List[T]) add(list *List[T], element T) *List[T] {
	// init first (or next link)
	temp := &List[T]{value: element, next: nil, prev: nil}
	if list == nil {
		// if this is the first item
		list = temp
	} else {
		head := list
		// else iterate through the list until we're at the end, then
		// add the next one
		for head.next != nil {
			head = head.next
		}
		temp.prev = head
		head.next = temp
	}
	return list
}

func main() {
	ll := List[int]{}
	ll.add(&ll, 1)
	ll.add(&ll, 2)
	fmt.Print(&ll)
}
