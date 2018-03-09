package main

import (
	"fmt"

	"github.com/wayneashleyberry/truecolor/pkg/truecolor"
)

func main() {
	x := truecolor.Text("Hello, World!").Foreground(186, 218, 85).Underline()
	fmt.Println(x)
}
