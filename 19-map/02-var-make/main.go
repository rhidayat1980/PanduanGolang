package main

import "fmt"

func main() {
	var myGreeting = make(map[string]string)

	fmt.Println(myGreeting == nil)

	myGreeting["Tim"] = "Good morning."
	myGreeting["Jenny"] = "Bonjour."

	fmt.Println(myGreeting)

	for k, v := range myGreeting {
		fmt.Println(k, "said", v)
	}
}
