package main

import "fmt"

func main() {
	var records [][]string

	// print initial record
	fmt.Println("record is", records)

	// student 1
	student1 := make([]string, 4)
	student1[0] = "Foster"
	student1[1] = "Nathan"
	student1[2] = "100.00"
	student1[3] = "74.00"

	fmt.Println("student1 :", student1)

	// store the record
	records = append(records, student1)

	fmt.Println("print record: ", records)

	// student 2
	student2 := make([]string, 4)
	student2[0] = "Gomez"
	student2[1] = "Lisa"
	student2[2] = "92.00"
	student2[3] = "96.00"

	fmt.Println("student2 :", student2)

	// store the record
	records = append(records, student2)

	// print record
	fmt.Println("record after student1 and student 2 added: ", records)
}
