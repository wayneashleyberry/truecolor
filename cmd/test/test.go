package main

import (
	"fmt"

	"github.com/wayneashleyberry/truecolor/pkg/truecolor"
)

func main() {
	badass := truecolor.Foreground(186, 218, 85)
	fmt.Println(badass.Text("Hello, World!"))
}
