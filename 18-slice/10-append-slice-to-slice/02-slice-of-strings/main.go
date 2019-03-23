package main

import "fmt"

func main() {
	mySlice := []string{"Monday", "Tuesday"}

	myOtherSlice := []string{"Wednesday", "Thursday", "Friday"}

	mySlice = append(mySlice, myOtherSlice...)

	fmt.Println(mySlice)

	for _, v := range mySlice {
		fmt.Println("the day of my life is", v)
	}
}
