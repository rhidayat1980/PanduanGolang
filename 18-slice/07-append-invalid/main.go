package main

import "fmt"

func main() {
	greeting := make([]string, 3, 5)
	fmt.Printf("%T\n", greeting)
	fmt.Println("length of greeting:", len(greeting))
	fmt.Println("capacity of greeting:", cap(greeting))

	greeting[0] = "Good morning!"
	greeting[1] = "Bonjour!"
	greeting[2] = "Buenos dias!"
	greeting[3] = "suprabadham"

	// the statement aabove is invalid, we have length of 2 but we wantg to give 3, go complain

	fmt.Println(greeting[2])

}
