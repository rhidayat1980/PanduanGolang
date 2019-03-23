package main

import "fmt"

func main() {
	m := []float64{1, 2, 3, 4, 5}
	n := average(m...)
	fmt.Println(n)
}

func average(sf ...float64) float64 {
	total := 0.0
	for _, v := range sf {
		total = v + total
	}
	return total / float64(len(sf))
}
