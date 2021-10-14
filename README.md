<img width="954" alt="Screenshot 2020-05-26 at 1 33 47 pm" src="https://user-images.githubusercontent.com/727262/82901224-90499780-9f55-11ea-8681-82a5ac8ebab3.png">

[![Go](https://github.com/wayneashleyberry/truecolor/actions/workflows/go.yml/badge.svg)](https://github.com/wayneashleyberry/truecolor/actions/workflows/go.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/wayneashleyberry/truecolor)](https://goreportcard.com/report/github.com/wayneashleyberry/truecolor)
[![Go Reference](https://pkg.go.dev/badge/github.com/wayneashleyberry/truecolor.svg)](https://pkg.go.dev/github.com/wayneashleyberry/truecolor)

> Truecolor is a Go package that provides methods for printing 24-bit true color text in your terminal. Not all terminals support 24-bit color, see https://git.io/vxe0J for more.

### Installation

```sh
go get github.com/wayneashleyberry/truecolor
```

### Usage

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
}
```
