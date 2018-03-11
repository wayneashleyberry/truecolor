[![GoDoc](https://godoc.org/github.com/wayneashleyberry/truecolor/pkg/color?status.svg)](https://godoc.org/github.com/wayneashleyberry/truecolor/pkg/color)
[![Go Report Card](https://goreportcard.com/badge/github.com/wayneashleyberry/truecolor)](https://goreportcard.com/report/github.com/wayneashleyberry/truecolor)

Package color provides methods for printing 24-bit truecolor in your terminal. The API is inspired by [chalk](https://github.com/chalk/chalk) and [color](https://github.com/fatih/color).

This is an experimental package and you probably shouldn't be using it.

```go
package main

import (
	"github.com/wayneashleyberry/truecolor/pkg/color"
)

func main() {
	color.Color(186, 218, 85).Println("#bada55")
	color.Black().Background(186, 218, 85).Println("#bada55")
	color.White().Underline().Print("underline\n")
	color.White().Dim().Println("dim")
	color.White().Italic().Println("italic")
	color.White().Bold().Println("bold")
	color.Color(255, 165, 00).Printf("Hello, %s!", "World")
}
```