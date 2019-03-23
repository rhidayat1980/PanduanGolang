package main

import "fmt"

func main() {
	greet("jane")
	greet("badrun")
}

func greet(s string) {
	fmt.Println("hello", s)
}

// greet is declared with a parameter
// when calling greet, pass in an argument
