package main

import "fmt"

func main() {
	fmt.Println(100)
	fmt.Println(100.929292)

	a := 100
	b := 200
	fmt.Println("a + b = ", a+b)

	// we cannot direct process different type in go for example
	// add int with float
	c := 17
	d := 23.45
	// fmt.Println("c + d = ", c+d) will create an error
	// we have to convert one to make it similar type with other
	fmt.Println("c + d = ", float64(c)+d)
}
