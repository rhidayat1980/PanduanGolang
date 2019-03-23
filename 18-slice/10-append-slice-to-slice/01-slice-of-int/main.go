package main

import "fmt"

func main() {
	mySlice := []int{1, 2, 3, 4, 5}
	myOtherSlice := []int{6, 7, 8, 9}

	mySlice = append(mySlice, myOtherSlice...)
	fmt.Println(mySlice)

	for k, v := range mySlice {
		fmt.Println(k, " - ", v)
	}

	for _, v := range mySlice {
		fmt.Println(v)
	}

	for k := range mySlice {
		fmt.Println(k)
	}
}
