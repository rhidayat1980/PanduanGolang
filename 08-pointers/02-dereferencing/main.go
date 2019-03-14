package main

import "fmt"

func main() {
	a := 43
	fmt.Println("nilai a adalah", a)
	fmt.Println("alamat memory a di", &a)

	var b = &a
	fmt.Println("alamat memory b di", b)
	fmt.Println("nilai b adalah", *b)
}

// b is an int pointer;
// b points to the memory address where an int is stored
// to see the value in that memory address, add a * in front of b
// this is known as dereferencing
// the * is an operator in this case
