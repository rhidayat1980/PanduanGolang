package main

import "fmt"

func main() {
	for i := 10; i <= 100; i++ {
		fmt.Println(i)
	}

	i := 20
	for i < 30 {
		fmt.Println(i)
		i++
	}
	fmt.Println("after looping")
	fmt.Println(i)

	fmt.Println("looping a:")
	a := 1
	for a <= 10 {
		fmt.Println(a)
		a++
	}

	for a = 100; a <= 200; a++ {
		fmt.Printf("%d \t %b \t %#x", a, a, a)
	}

}
