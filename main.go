package main

import (
	"fmt"
	"github.com/PurinPintakhiew/go-convert-thai/convert"
)

func main() {
	num := convert.NumberFormat(-10000.515, 2, true)

	fmt.Println(num)
}
