package logger

import (
	"fmt"
	"strings"
)

const Clear uint16 = 1 << 15

const (
	_ uint16 = iota
	Black
	Red
	Green
	Yellow
	Blue
	Magenta
	Cyan
	White
	fgOffset = 30 - 1
)

var (
	defaultColor = Clear
	indentLevel  = 0
)

func Info(msg string) {
	lines := indentLines(msg)
	for _, line := range lines {
		fmt.Println(colorize(Clear, " INFO : "+line))
	}
}

func Warn(msg string) {
	lines := indentLines(msg)
	for _, line := range lines {
		fmt.Println(colorize(Magenta, " WARN : "+line))
	}
}

func Error(msg string) {
	lines := indentLines(msg)
	for _, line := range lines {
		fmt.Println(colorize(Red, "ERROR : "+line))
	}
}

func Debug(msg string) {
	lines := indentLines(msg)
	for _, line := range lines {
		fmt.Println(colorize(Clear, "DEBUG : "+line))
	}
}

func Color(color uint16, fn func()) {
	org := defaultColor
	defaultColor = color
	fn()
	defaultColor = org
}

func WithIndent(fn func()) {
	indentLevel++
	fn()
	indentLevel--
}

func indentLines(str string) []string {
	indent := ""
	for i := 0; i < indentLevel; i++ {
		indent += "  "
	}

	lines := strings.Split(str, "\n")
	ret := make([]string, 0, len(lines))
	for _, line := range lines {
		ret = append(ret, indent+line)
	}
	return ret
}

func colorize(color uint16, msg string) string {
	if color == Clear {
		if defaultColor == Clear {
			return msg
		} else {
			color = defaultColor
		}
	}
	return fmt.Sprintf("\033[%dm%s\033[0m", fgOffset+color, msg)
}
