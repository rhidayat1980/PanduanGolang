package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	p1 := person{"james", 20}
	fmt.Println(p1)
	fmt.Println(p1.name)
	fmt.Println(p1.age)

	fmt.Printf("%T\n", p1)
}
