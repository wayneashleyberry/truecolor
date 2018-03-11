package main

import (
	"github.com/wayneashleyberry/truecolor/pkg/truecolor"
)

func main() {
	truecolor.Foreground(186, 218, 85).Println("#bada55")
	truecolor.Black().Background(186, 218, 85).Println("#bada55")
	truecolor.White().Underline().Print("underline\n")
	truecolor.White().Dim().Println("dim")
	truecolor.White().Italic().Println("italic")
	truecolor.White().Bold().Println("bold")
	truecolor.Foreground(255, 165, 00).Printf("Hello, %s!", "World")
}
