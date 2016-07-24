package logger

import (
	"fmt"
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
)

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
