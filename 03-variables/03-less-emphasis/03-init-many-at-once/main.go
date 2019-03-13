package main

import "fmt"

func main() {
	var message = "hello golang"
	var a, b, c = 10, 20, 30

	fmt.Printf("%v %T \n", message, message)
	fmt.Printf("%v %T \n", a, a)
	fmt.Printf("%v %T \n", b, b)
	fmt.Printf("%v %T \n", c, c)
}
