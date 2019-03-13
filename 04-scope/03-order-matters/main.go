package main

import "fmt"

func main() {
	x := 42
	fmt.Println(x)
	fmt.Println(y)
	// if we put x:= 42 here, and remove x:= 42 above, the code will error
}

var y = 50
