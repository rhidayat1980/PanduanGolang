package main

import "fmt"

func main() {
	greet("jane", "badrun")
	greet("aco", "bobi")
}

func greet(m string, n string) {
	fmt.Println(m, n)
}
