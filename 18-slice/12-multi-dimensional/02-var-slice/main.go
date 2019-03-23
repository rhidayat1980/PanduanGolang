package main

import "fmt"

func main() {
	var student []string
	var students []string

	fmt.Println(student)
	fmt.Println(students)
	fmt.Println(student == nil)
	fmt.Println("lenght of empty slice with var:", len(student))
	fmt.Println("capacity of empty string with var:", cap(student))

}
