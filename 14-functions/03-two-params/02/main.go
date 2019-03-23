package main

import "fmt"

func main() {
	greet("aco", "bobi")
	greet("bucek", "arman")
}

func greet(fname, lname string) {
	fmt.Println(fname, lname)
}
