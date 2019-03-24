package main

import "fmt"

func main() {
	myGreeting := map[string]string{
		"Tim":   "Good morning!",
		"Jenny": "Bonjour!",
	}

	myGreeting["Harleen"] = "Howdy"

	fmt.Println(myGreeting)

	for k, v := range myGreeting {
		fmt.Println(k, "said", v)
	}
}
