package main

import "fmt"

func main() {
	var message = "hello"
	var a, b, c = 1, false, "Go"
	fmt.Printf("%v %T\n", message, message)
	fmt.Printf("%v %T\n", a, a)
	fmt.Printf("%v %T\n", b, b)
	fmt.Printf("%v %T\n", c, c)

}
