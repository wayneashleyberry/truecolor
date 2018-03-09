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
	hasForeground bool
	bg            Color
	hasBackground bool
	text          string
	underline     bool
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
	m.underline = true
	return m
}

func Text(text string) *Message {
	return &Message{
		text: text,
	}
}

func (m *Message) Text(text string) *Message {
	m.text = text
	return m
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
	if m.underline {
		fmt.Fprintf(&b, "\u001B[4m")
	}
	fmt.Fprintf(&b, "%s", m.text)
	if m.underline {
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
