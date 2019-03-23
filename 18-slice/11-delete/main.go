package main

import "fmt"

func main() {
	mySlice := []string{"Monday", "Tuesday"}
	myOtherSlice := []string{"Wednesday", "Thursday", "Friday"}
	myOtherOtherSlice := []string{"Saturday", "Sunday"}

	myDays := append(mySlice, myOtherSlice...)
	myFullDays := append(myDays, myOtherOtherSlice...)

	fmt.Println("my full days ", myFullDays)

	for _, v := range myFullDays {
		fmt.Println("my every days", v)
	}

	// deleting slices
	// here we remove wednesday

	mySlice = append(myFullDays[:2], myFullDays[3:]...)
	fmt.Println(mySlice)

	for _, v := range mySlice {
		fmt.Println("my every days", v)
	}

}
