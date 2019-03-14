package main

import "fmt"

func main() {
	x := 13 % 3
	fmt.Println(x)
	if x == 1 {
		fmt.Println("odd")
	} else {
		fmt.Println("even")
	}

	for i := 1; i < 70; i++ {
		if i%2 == 1 {
			fmt.Println(i, "is Odd")
		} else {
			fmt.Println(i, "is Even")
		}
	}
}
