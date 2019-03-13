package main

import "fmt"

const (
	pi       = 3.14
	language = "GO"
)

func main() {
	fmt.Println("pi - ", pi)
	fmt.Println("language - ", language)

	fmt.Printf("%v %T\n", pi, pi)
	fmt.Printf("%v %T\n", language, language)
}
