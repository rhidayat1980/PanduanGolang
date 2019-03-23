package main

import "fmt"

func main() {
	m := average(1, 2, 3, 4, 5)
	fmt.Println(m)
}

func average(a ...float64) float64 {
	fmt.Println(a)
	fmt.Printf("%T \n", a)
	var total float64
	for _, v := range a {
		total += v
	}
	return total / float64(len(a))
}
