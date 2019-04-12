package main

import (
	"fmt"
	"strconv"
)

func main() {

	// convert int to ascii
	var x = 5
	str := "Hello world " + strconv.Itoa(x)
	fmt.Println(str)

	// via fmt.Sprint
	str = "Hello world " + fmt.Sprint(x)
	fmt.Println(str)

	// ascii to int
	i, _ := strconv.Atoi("32")
	fmt.Println(i + 10)
}
