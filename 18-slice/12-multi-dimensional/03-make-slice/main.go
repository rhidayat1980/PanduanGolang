package main

import "fmt"

func main() {
	student := make([]string, 35)
	students := make([]string, 35, 60)

	fmt.Println(student)
	fmt.Println(students)

	fmt.Println(student == nil)

	fmt.Println("length of string slice with make: ", len(student))
	fmt.Println("capacity of string slice with make: ", cap(student))

	fmt.Println("length of string slice with make: ", len(students))
	fmt.Println("capacity of string slice with make: ", cap(students))
}
