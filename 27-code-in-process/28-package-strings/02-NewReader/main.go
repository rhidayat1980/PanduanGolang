package main

import (
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	dst, err := os.Create(os.Args[1])
	if err != nil {
		log.Fatalln("error creating destination file:", err.Error())
	}
	defer dst.Close()

	rdr := strings.NewReader("Hello World")
	io.Copy(dst, rdr)
}

/*
io.Copy
func Copy(dst Writer, src Reader) (written int64, err error)
strings.NewReader
func NewReader(s string) *Reader
strings
func (r *Reader) Read(b []byte) (n int, err error)
*/

/*
step 1 - at command line enter:
go install
step 2 - at command line enter:
programName initial.txt second.txt
*/
