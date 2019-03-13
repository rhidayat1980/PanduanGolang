// init shorthand can only be used inside a function

package main

import "fmt"

func main() {
	message := "hello"
	a, b, c := 1, true, "go"
	d := 3.14
	e := false

	fmt.Printf("%v %T \n", message, message)
	fmt.Printf("%v %T \n", a, a)
	fmt.Printf("%v %T \n", b, b)
	fmt.Printf("%v %T \n", c, c)
	fmt.Printf("%v %T \n", d, d)
	fmt.Printf("%v %T \n", e, e)
}
