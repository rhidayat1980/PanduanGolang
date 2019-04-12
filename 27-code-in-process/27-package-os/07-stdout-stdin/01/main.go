package main

import (
	"io"
	"log"
	"os"
)

func main() {
	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatalln("my program broke: ", err.Error())
	}
	defer f.Close()

	io.Copy(os.Stdout, f)
}

/*
go build .
touch test.txt
echo "test me here" > test.txt

./01 test.txt
*/
