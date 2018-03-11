package color

import (
	"fmt"
	"strings"
)

const closeFgColor = "\u001B[39m"
const closeBgColor = "\u001B[49m"

type color struct {
	r, g, b int
}

type message struct {
	fg            color
	bg            color
	text          string
	hasForeground bool
	hasBackground bool
	isUnderlined  bool
	isItalic      bool
	isDim         bool
	isBold        bool
}

func Foreground(r, g, b int) *message {
	return &message{
		hasForeground: true,
		fg: color{
			r: r,
			g: g,
			b: b,
		},
	}
}

func Background(r, g, b int) *message {
	return &message{
		hasBackground: true,
		bg: color{
			r: r,
			g: g,
			b: b,
		},
	}
}

func White() *message {
	return &message{
		hasForeground: true,
		fg:            color{255, 255, 255},
	}
}

func Black() *message {
	return &message{
		hasForeground: true,
		fg:            color{0, 0, 0},
	}
}

func (m *message) Foreground(r, g, b int) *message {
	m.hasForeground = true
	m.fg = color{r, g, b}
	return m
}

func (m *message) Background(r, g, b int) *message {
	m.bg = color{r, g, b}
	m.hasBackground = true
	return m
}

func (m *message) Underline() *message {
	m.isUnderlined = true
	return m
}

func (m *message) Dim() *message {
	m.isDim = true
	return m
}

func (m *message) Italic() *message {
	m.isItalic = true
	return m
}

func (m *message) Bold() *message {
	m.isBold = true
	return m
}

func Text(text string) *message {
	return &message{
		text: text,
	}
}

func (m *message) Print(text string) {
	m.text = text
	fmt.Print(m)
}

func (m *message) Println(text string) {
	m.text = text
	fmt.Println(m)
}

func (m *message) Printf(format string, a ...interface{}) {
	m.text = fmt.Sprintf(format, a...)
	fmt.Print(m)
}

func (m *message) String() string {
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
