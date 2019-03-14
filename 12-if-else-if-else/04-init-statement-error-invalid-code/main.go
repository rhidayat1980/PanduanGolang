package main

import "fmt"

func main() {
	b := true

	if food := "Chocolate"; b {
		fmt.Println(food)
	}
	// fmt.Println(food) // this code is wrong, since variable food is only live in scope if
}
