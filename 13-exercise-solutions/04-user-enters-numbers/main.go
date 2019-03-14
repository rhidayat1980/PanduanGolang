package main

import "fmt"

func main() {
	var numbOne int
	var numbTwo int
	fmt.Println("Please enter a large number: ")
	fmt.Scan(&numbOne)

	fmt.Println("Please enter a small number: ")
	fmt.Scan(&numbTwo)

	fmt.Println(numbOne, "%", numbTwo, " = ", numbOne%numbTwo)
}
