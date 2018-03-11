package color

import (
	"fmt"
	"image/color"
	"strings"
)

const closeFgColor = "\u001B[39m"
const closeBgColor = "\u001B[49m"

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

func Background(r, g, b uint8) *Message {
	return &Message{
		hasBackground: true,
		bg:            color.RGBA{r, g, b, 0},
	}
}

func White() *Message {
	return &Message{
		hasForeground: true,
		fg:            color.RGBA{255, 255, 255, 0},
	}
}

func Black() *Message {
	return &Message{
		hasForeground: true,
		fg:            color.RGBA{0, 0, 0, 0},
	}
}

func (m *Message) Color(r, g, b uint8) *Message {
	m.hasForeground = true
	m.fg = color.RGBA{r, g, b, 0}
	return m
}

func (m *Message) Background(r, g, b uint8) *Message {
	m.bg = color.RGBA{r, g, b, 0}
	m.hasBackground = true
	return m
}

func (m *Message) Underline() *Message {
	m.isUnderlined = true
	return m
}

func (m *Message) Dim() *Message {
	m.isDim = true
	return m
}

func (m *Message) Italic() *Message {
	m.isItalic = true
	return m
}

func (m *Message) Bold() *Message {
	m.isBold = true
	return m
}

func Text(text string) *Message {
	return &Message{
		text: text,
	}
}

func (m *Message) Print(text string) {
	m.text = text
	fmt.Print(m)
}

func (m *Message) Println(text string) {
	m.text = text
	fmt.Println(m)
}

func (m *Message) Printf(format string, a ...interface{}) {
	m.text = fmt.Sprintf(format, a...)
	fmt.Print(m)
}

func (m *Message) String() string {
	var b strings.Builder
	mod := 38
	if m.hasForeground {
		fmt.Fprintf(&b, "\u001B[%d;2;%d;%d;%dm", mod, m.fg.R, m.fg.G, m.fg.B)
	}
	if m.hasBackground {
		fmt.Fprintf(&b, "\u001B[%d;2;%d;%d;%dm", mod+10, m.bg.R, m.bg.G, m.bg.B)
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
		fmt.Fprintf(&b, "%s", closeBgColor)
	}
	if m.hasForeground {
		fmt.Fprintf(&b, "%s", closeFgColor)
	}
	return b.String()
}
