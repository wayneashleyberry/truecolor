```go
package main

import (
	"github.com/wayneashleyberry/truecolor/pkg/color"
)

func main() {
	color.Foreground(186, 218, 85).Println("#bada55")
	color.Black().Background(186, 218, 85).Println("#bada55")
	color.White().Underline().Print("underline\n")
	color.White().Dim().Println("dim")
	color.White().Italic().Println("italic")
	color.White().Bold().Println("bold")
	color.Foreground(255, 165, 00).Printf("Hello, %s!", "World")
}
```