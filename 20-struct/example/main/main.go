package main

import (
	"fmt"

	"github.com/rhidayat1980/PanduanGolang/20-struct/example/model"
)

func main() {
	c := model.Customer{
		Id:   1,
		Name: "Rachmat Hidayat",
	}
	c.Married = true

	a := model.Address{}
	fmt.Println(c)
	fmt.Println(a)
}
