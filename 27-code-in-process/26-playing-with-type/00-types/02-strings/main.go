package main

import "fmt"

func main() {
	// escape sequences
	fmt.Println("Hello\tWorld\nHow are you?")

	// sequence of bytes
	intro := "Four score and seven years ago...."
	fmt.Println(intro)
	fmt.Println([]byte(intro))

	// immutable
	fmt.Println(intro)
	fmt.Println(&intro)
	intro = "Hahahaha!"
	fmt.Println(intro)
	fmt.Println(&intro)

	// len of string
	fmt.Println(len("hello"))

	// len of chinese
	fmt.Println(len("世界"))

	// binary
	intro = "Four 世"
	fmt.Printf("%T\n", intro)
	fmt.Println(intro)
	bs := []byte(intro)
	fmt.Println(bs)
	fmt.Printf("%T\n", bs)
	fmt.Println("********")
	fmt.Printf("%d\n", bs)

	for _, v := range bs {
		fmt.Printf("%d\t\t %#x\t %b\n", v, v, v)
	}

	fmt.Println("*********")
	y := 9999999999999999
	fmt.Printf("%d\t\t %#x\t %b\n", y, y, y)

	fmt.Println(&y)
	fmt.Println("*********")

	z := 'h'
	fmt.Printf("%T\n", z)

	// index access
	greeting := "Hello"
	fmt.Println(greeting)
	fmt.Println(greeting[0])
	fmt.Println(greeting[4])

	// slicing string
	greeting = "Hello"
	fmt.Println(greeting)
	fmt.Println(greeting[0])
	fmt.Println(greeting[4])
	fmt.Println("-------------")
	fmt.Println("What the ... ")
	fmt.Println(greeting[:4])
	fmt.Println("... did that just do?")

	// slicing string 2
	greeting = "Hello"
	fmt.Println(greeting)
	fmt.Println(greeting[:4])
	fmt.Println(greeting[0:4])
	fmt.Println(greeting[1:4])
	fmt.Println(greeting[1:])

	// invalid negative index
	greeting = "Hello"
	fmt.Println(greeting)
	fmt.Println(greeting[:4])
	fmt.Println(greeting[0:4])
	fmt.Println(greeting[1:4])
	fmt.Println(greeting[1:])
	// fmt.Println(greeting[:-2]) // invalid

	// string concatenation
	fname := "James"
	lname := "Bond"
	fmt.Println(fname + " " + lname)
}
