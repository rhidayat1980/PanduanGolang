package main

import "fmt"

const (
	a = iota
	b
	c
)

func main() {
	fmt.Printf("%v %T\n", a, a)
	fmt.Printf("%v %T\n", b, b)
	fmt.Printf("%v %T\n", c, c)
}
