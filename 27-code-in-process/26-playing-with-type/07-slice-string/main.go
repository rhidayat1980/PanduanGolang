package main

import "fmt"

type mySentence []string

func main() {
	var message mySentence = []string{"Hello World!", "More coffee"}
	fmt.Println(message)
	fmt.Printf("%T\n", message)
}
