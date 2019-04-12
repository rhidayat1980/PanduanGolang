package main

import "fmt"

type mySentence []string

func main() {
	var message mySentence = []string{"Hello World!", "More coffee"}
	fmt.Println(message)

	fmt.Printf("%T\n", message)
	fmt.Printf("%T\n", []string(message))
	fmt.Println(message)

	for _, value := range message {
		fmt.Println(value)
	}
}
