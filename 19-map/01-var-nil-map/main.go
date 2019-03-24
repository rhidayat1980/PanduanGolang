package main

import "fmt"

func main() {
	var myGreeting map[string]string
	fmt.Println(myGreeting)
	fmt.Println(myGreeting == nil)

	/*
		updating empty map will create panic!!
	*/
	// myGreeting["Tim"] = "Good morning."
	// myGreeting["Jenny"] = "Bonjour."
	// fmt.Println(myGreeting)
}
