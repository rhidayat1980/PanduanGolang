package main

import (
	"fmt"

	"github.com/rhidayat1980/PanduanGolang/02-package/icomefromalaska"
	"github.com/rhidayat1980/PanduanGolang/02-package/stringutil"
)

func main() {

	fmt.Println(stringutil.Reverse("hello, golang!"))
	fmt.Println(stringutil.MyName)
	fmt.Println(icomefromalaska.BearName)
}
