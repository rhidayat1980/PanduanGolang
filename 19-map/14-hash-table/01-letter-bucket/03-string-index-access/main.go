package main

import "fmt"

func main() {
	word := "Hello"
	letter := rune(word[0])
	fmt.Println(letter)

	letter = rune(word[1])
	fmt.Println(letter)
}
