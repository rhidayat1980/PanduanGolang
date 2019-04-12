package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func cp(srcName, dstName string) error {
	src, err := os.Open(srcName)
	if err != nil {
		return fmt.Errorf("error opening source file: %v", err)
	}
	defer src.Close()

	dst, err := os.Create(dstName)
	if err != nil {
		return fmt.Errorf("error creating destination file: %v", err)
	}
	defer dst.Close()

	bs := bufio.NewReader(src)

	_, err = io.Copy(dst, bs)
	if err != nil {
		return fmt.Errorf("error writing to destination file: %v", err)
	}

	return nil
}

func main() {
	if len(os.Args) < 3 {
		log.Fatalln("Usage: 01-NewReader <src> <dst>")
	}

	srcName := os.Args[1]
	dstName := os.Args[2]

	err := cp(srcName, dstName)
	if err != nil {
		log.Fatalln(err)
	}
}
