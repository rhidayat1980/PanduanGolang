package main

import "fmt"

// Person struct
type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func main() {
	// declaring a variable of a `struct` type
	var p Person // declaring with zero value
	fmt.Println(p)

	// Declaring and initializing a struct using a struct literal
	p1 := Person{"Azza", "Midaskha", 29}
	fmt.Println("Person1: ", p1)

	// Naming fields while initializing a struct
	p2 := Person{
		FirstName: "John",
		LastName:  "Doe",
		Age:       40,
	}
	fmt.Println(p2)

	// uninitialized fields are set to their corresponding zero-value
	p3 := Person{FirstName: "Inara"}
	fmt.Println(p3)
}
