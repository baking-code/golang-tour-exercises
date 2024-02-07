package main

import "fmt"

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func MakePerson(firstName, lastName string, age int) Person {
	return Person{
		FirstName: firstName, LastName: lastName, Age: age,
	}
}

func MakePersonPointer(firstName, lastName string, age int) *Person {
	return &Person{
		FirstName: firstName, LastName: lastName, Age: age,
	}
}

func main() {
	per := MakePerson("john", "doe", 30)
	perPointer := MakePersonPointer("jane", "doe", 20)
	fmt.Println(per, perPointer)
	// go build -gcflags="-m"
	// per escapes to heap
	// two separate instances of &Person{} from perPointer line escape to heap - pointer returned from function cannot be on stack

}
