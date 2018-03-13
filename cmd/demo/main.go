package main

import (
	"fmt"
	col "image/color"

	color "github.com/wayneashleyberry/truecolor/pkg/color"
)

func main() {
	color.Color(186, 218, 85).Println("Hello, World!")
	color.Black().Background(186, 218, 85).Println("Hello, World!")
	color.White().Underline().Print("Hello, World!\n")
	color.White().Dim().Println("Hello, World!")
	color.White().Italic().Println("Hello, World!")
	color.White().Bold().Println("Hello, World!")
	color.Color(255, 165, 00).Printf("Hello, %s!\n", "World")
	fmt.Printf("Hello, %s\n", color.Color(255, 0, 0).Sprint("World", "!"))
	color.RGBA(col.RGBA{100, 200, 10, 0}).Println("Hello, World!")

	for i := 0; i <= 255; i++ {
		color.Background(uint8(i), 0, 0).Print(" ")
	}
	fmt.Println("")
	for i := 0; i <= 255; i++ {
		color.Background(0, uint8(i), 0).Print(" ")
	}
	fmt.Println("")
	for i := 0; i <= 255; i++ {
		color.Background(0, 0, uint8(i)).Print(" ")
	}
	fmt.Println("")
	for i := 0; i <= 255; i++ {
		color.Background(uint8(i), uint8(i), uint8(i)).Print(" ")
	}
	fmt.Println("")
}
