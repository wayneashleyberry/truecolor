[![GoDoc](https://godoc.org/github.com/wayneashleyberry/truecolor/pkg/color?status.svg)](https://godoc.org/github.com/wayneashleyberry/truecolor/pkg/color)
[![Go Report Card](https://goreportcard.com/badge/github.com/wayneashleyberry/truecolor)](https://goreportcard.com/report/github.com/wayneashleyberry/truecolor)

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