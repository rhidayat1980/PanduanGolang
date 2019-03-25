package main

import "fmt"

type foo int

func main() {
	var myAge foo
	myAge = 40

	fmt.Printf("%T\n", myAge)
	fmt.Println(myAge)

}
