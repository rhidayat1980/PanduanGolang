package main

import "fmt"

func main() {
	fmt.Println(greet("bucek", "deep"))
	fmt.Println(greet("arman", "maulana"))
}

func greet(fname, lname string) string {
	return fmt.Sprint(fname, " ", lname)
}
