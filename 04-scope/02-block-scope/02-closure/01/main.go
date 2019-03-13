package main

import "fmt"

func main() {
	x := 42
	fmt.Println(x)
	{
		fmt.Println(x)
		y := "the credit is belongs with the one who is in the ring"
		fmt.Println(y)
	}
	// fmt.Println(y) will error, scope of y is outside the scope
}
