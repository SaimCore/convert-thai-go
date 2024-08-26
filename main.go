package main

import (
	"fmt"
	"github.com/PurinPintakhiew/go-convert-thai/convert"
)

func main() {
	bath := convert.BathText(100.00)

	fmt.Println(bath)
}
