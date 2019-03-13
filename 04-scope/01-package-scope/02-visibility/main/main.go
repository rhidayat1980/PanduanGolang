package main

import (
	"fmt"

	"github.com/rhidayat1980/PanduanGolang/04-scope/01-package-scope/02-visibility/vis"
)

func main() {
	fmt.Println(vis.MyName)
	// fmt.Println(vis.yourName) will error because yourName is not exproted
	vis.PrintVar()
}
