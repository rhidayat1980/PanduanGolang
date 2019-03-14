package main

import "fmt"

func main() {
	i := 0 // start from here
	for {
		fmt.Println(i)
		if i >= 10 {
			break
		}
		i++
	}
}
