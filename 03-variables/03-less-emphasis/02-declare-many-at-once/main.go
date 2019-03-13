package main

import "fmt"

func main() {
	var message string
	var a, b, c int
	a = 1

	fmt.Printf("%v \n", message)
	fmt.Printf("%v %v %v\n", a, b, c)

	fmt.Printf("%T \n", message)
	fmt.Printf("%T %T %T\n", a, b, c)
}
