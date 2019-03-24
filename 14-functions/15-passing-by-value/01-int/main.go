package main

import "fmt"

func changeMe(z int) {
	fmt.Println(z)
	z = 24
}

func main() {
	age := 44
	changeMe(age)
	fmt.Println(age)
}
