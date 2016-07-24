package logger

import (
	"fmt"
	"log"
	"strings"
)

const (
	DEBUG = iota
	INFO
	WARN
	ERROR
	FATAL
	UNKNOWN
)

var (
	indentLevel = 0
	logLevel    = INFO
)

func Info(msg string) {
	if logLevel <= INFO {
		lines := indentLines(msg)
		for _, line := range lines {
			fmt.Println(colorize(Clear, " INFO : "+line))
		}
	}
}

func Warn(msg string) {
	if logLevel <= WARN {
		lines := indentLines(msg)
		for _, line := range lines {
			fmt.Println(colorize(Magenta, " WARN : "+line))
		}
	}
}

func Error(msg string) {
	if logLevel <= ERROR {
		lines := indentLines(msg)
		for _, line := range lines {
			fmt.Println(colorize(Red, "ERROR : "+line))
		}
	}
}

func Debug(msg string) {
	if logLevel <= DEBUG {
		lines := indentLines(msg)
		for _, line := range lines {
			fmt.Println(colorize(Clear, "DEBUG : "+line))
		}
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

func SetLogLevel(level string) {
	if len(level) == 0 {
		logLevel = INFO
		return
	}

	switch strings.ToLower(level) {
	case "debug":
		logLevel = DEBUG
	case "info":
		logLevel = INFO
	case "warn":
		logLevel = WARN
	case "error":
		logLevel = ERROR
	case "fatal":
		logLevel = FATAL
	case "unknown":
		logLevel = UNKNOWN
	default:
		log.Fatal("unexpected log level: ", level)
	}
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
