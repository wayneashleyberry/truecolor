[![GoDoc](https://godoc.org/github.com/wayneashleyberry/truecolor/pkg/color?status.svg)](https://godoc.org/github.com/wayneashleyberry/truecolor/pkg/color)
[![Go Report Card](https://goreportcard.com/badge/github.com/wayneashleyberry/truecolor)](https://goreportcard.com/report/github.com/wayneashleyberry/truecolor)

```go
package main

import (
	"github.com/wayneashleyberry/truecolor/pkg/color"
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
}
```