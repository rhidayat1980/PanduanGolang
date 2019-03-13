package main

import "fmt"

func main() {
	a := "stored in a"
	b := "stored in b"
	fmt.Println("a - ", a)
	fmt.Println("b - ", b) // we add this to used bariable b
}

/*
since b not in used by our code, the go compiler will complain, and we cannot run this code. variable in function level after created should be used.
the rules in variable inside function is, never created variable if you dont want to used it.
*/
