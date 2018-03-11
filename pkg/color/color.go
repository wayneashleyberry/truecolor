// Package color provides methods for printing 24-bit truecolor in your terminal.
//
// Not all terminals support 24-bit color, see https://git.io/vxe0J for more.
package color

import (
	"fmt"
	"image/color"
	"strings"
)

// Message holds the formatting options for printing
type Message struct {
	fg            color.RGBA
	bg            color.RGBA
	text          string
	hasForeground bool
	hasBackground bool
	isUnderlined  bool
	isItalic      bool
	isDim         bool
	isBold        bool
}

// Color sets the foreground color
func Color(r, g, b uint8) *Message {
	return &Message{
		hasForeground: true,
		fg: color.RGBA{
			R: r,
			G: g,
			B: b,
			A: 0,
		},
	}
}

// Background sets the background color
func Background(r, g, b uint8) *Message {
	return &Message{
		hasBackground: true,
		bg:            color.RGBA{r, g, b, 0},
	}
}

// White is an alias to Color(255, 255, 255)
func White() *Message {
	return &Message{
		hasForeground: true,
		fg:            color.RGBA{255, 255, 255, 0},
	}
}

// Black is an alias to Color(0, 0, 0)
func Black() *Message {
	return &Message{
		hasForeground: true,
		fg:            color.RGBA{0, 0, 0, 0},
	}
}

// Color sets the foreground color
func (m *Message) Color(r, g, b uint8) *Message {
	m.hasForeground = true
	m.fg = color.RGBA{r, g, b, 0}
	return m
}

// Background sets the background color
func (m *Message) Background(r, g, b uint8) *Message {
	m.bg = color.RGBA{r, g, b, 0}
	m.hasBackground = true
	return m
}

// Underline will underline the text
func (m *Message) Underline() *Message {
	m.isUnderlined = true
	return m
}

// Dim will dim the text
func (m *Message) Dim() *Message {
	m.isDim = true
	return m
}

// Italic will format the text with italics
func (m *Message) Italic() *Message {
	m.isItalic = true
	return m
}

// Bold will format the text in bold
func (m *Message) Bold() *Message {
	m.isBold = true
	return m
}

// Print will print the formatted text, like fmt.Print
func (m *Message) Print(text string) {
	m.text = text
	fmt.Print(m)
}

// Sprint will print the formatted text, like fmt.Print
func (m *Message) Sprint(a ...interface{}) string {
	m.text = fmt.Sprint(a...)
	return fmt.Sprint(m)
}

// Println will print the formatted text, like fmt.Println
func (m *Message) Println(text string) {
	m.text = text
	fmt.Println(m)
}

// Printf will print the formatted text, like fmt.Printf
func (m *Message) Printf(format string, a ...interface{}) {
	m.text = fmt.Sprintf(format, a...)
	fmt.Print(m)
}

// String will return the formatted text, with ansii excape codes, as a string
func (m *Message) String() string {
	var b strings.Builder
	if m.hasForeground {
		fmt.Fprintf(&b, "\u001B[38;2;%d;%d;%dm", m.fg.R, m.fg.G, m.fg.B)
	}
	if m.hasBackground {
		fmt.Fprintf(&b, "\u001B[48;2;%d;%d;%dm", m.bg.R, m.bg.G, m.bg.B)
	}
	if m.isUnderlined {
		fmt.Fprintf(&b, "\u001B[4m")
	}
	if m.isDim {
		fmt.Fprintf(&b, "\u001B[2m")
	}
	if m.isItalic {
		fmt.Fprintf(&b, "\u001B[3m")
	}
	if m.isBold {
		fmt.Fprintf(&b, "\u001B[1m")
	}
	fmt.Fprintf(&b, "%s", m.text)
	if m.isBold {
		fmt.Fprintf(&b, "\u001B[21m")
	}
	if m.isItalic {
		fmt.Fprintf(&b, "\u001B[23m")
	}
	if m.isDim {
		fmt.Fprintf(&b, "\u001B[22m")
	}
	if m.isUnderlined {
		fmt.Fprintf(&b, "\u001B[24m")
	}
	if m.hasBackground {
		fmt.Fprintf(&b, "%s", "\u001B[49m")
	}
	if m.hasForeground {
		fmt.Fprintf(&b, "%s", "\u001B[39m")
	}
	return b.String()
}
