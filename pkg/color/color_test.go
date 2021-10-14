package color

import (
	"fmt"
	"strconv"
)

func ExampleColor() {
	fmt.Println(strconv.QuoteToASCII(Color(186, 218, 85).Sprint("Hello, World!")))
	// Output: "\x1b[38;2;186;218;85mHello, World!\x1b[39m"
}
