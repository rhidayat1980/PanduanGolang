package main

import "fmt"

func main() {
	student := []string{}
	students := []string{}

	fmt.Println(student)
	fmt.Println(students)
	fmt.Println(student == nil)

	fmt.Println("lenght of empty slice:", len(student))
	fmt.Println("capacity of empty slice:", cap(student))
}
