package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

type doubleZero struct {
	person
	LicenceToKill bool
}

func main() {
	p1 := doubleZero{
		person: person{
			first: "James",
			last:  "Bond",
			age:   30,
		},
		LicenceToKill: true,
	}

	p2 := doubleZero{
		person: person{
			first: "Lady",
			last:  "gaga",
			age:   27,
		},
		LicenceToKill: false,
	}

	fmt.Println(p1.first, p1.last, p1.age, p1.LicenceToKill)
	fmt.Println(p2.first, p2.last, p2.age, p2.LicenceToKill)
}
