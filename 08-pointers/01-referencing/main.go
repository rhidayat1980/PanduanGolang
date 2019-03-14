package main

import "fmt"

func main() {
	a := 43
	fmt.Println("nilai a adalah", a)
	fmt.Println("alamat memory a di", &a)

	var b = &a
	fmt.Println("nilai b adalah", *b) // dereferencing untuk dapat nilai nya
	fmt.Println("alamat memory b di", &b)

	a = 100
	fmt.Println("nilai a adalah", a)
	fmt.Println("alamat memory a di", &a)

	fmt.Println("nilai b adalah", *b)
	fmt.Println("alamat memory b di", &b)

	*b = 200

	fmt.Println("nilai a adalah", a)
	fmt.Println("alamat memory a di", &a)

	fmt.Println("nilai b adalah", *b)
	fmt.Println("alamat memory b di", &b)

}
