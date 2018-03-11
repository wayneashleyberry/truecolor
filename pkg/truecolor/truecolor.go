package truecolor

import (
	"fmt"
	"strings"
)

const closeFgColor = "\u001B[39m"
const closeBgColor = "\u001B[49m"

type Color struct {
	r, g, b int
}

type Message struct {
	fg            Color
	bg            Color
	text          string
	hasForeground bool
	hasBackground bool
	isUnderlined  bool
	isItalic      bool
	isDim         bool
	isBold        bool
}

func Foreground(r, g, b int) *Message {
	return &Message{
		hasForeground: true,
		fg: Color{
			r: r,
			g: g,
			b: b,
		},
	}
}

func Background(r, g, b int) *Message {
	return &Message{
		hasBackground: true,
		bg: Color{
			r: r,
			g: g,
			b: b,
		},
	}
}

func White() *Message {
	return &Message{
		hasForeground: true,
		fg:            Color{255, 255, 255},
	}
}

func Black() *Message {
	return &Message{
		hasForeground: true,
		fg:            Color{0, 0, 0},
	}
}

func (m *Message) Foreground(r, g, b int) *Message {
	m.hasForeground = true
	m.fg = Color{r, g, b}
	return m
}

func (m *Message) Background(r, g, b int) *Message {
	m.bg = Color{r, g, b}
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
		fmt.Fprintf(&b, "\u001B[%d;2;%d;%d;%dm", mod, m.fg.r, m.fg.g, m.fg.b)
	}
	if m.hasBackground {
		fmt.Fprintf(&b, "\u001B[%d;2;%d;%d;%dm", mod+10, m.bg.r, m.bg.g, m.bg.b)
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
