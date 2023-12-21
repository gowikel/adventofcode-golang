package utils

import (
	"fmt"
)

var codes = map[string]int{
	"normal":     0,
	"bold":       1,
	"faint":      2,
	"italicized": 3,
	"underlined": 4,
	"blink":      5,
	"inverse":    7,
	"invisible":  8,
}

func wrapSGR(str string, mode string) string {
	code, ok := codes[mode]
	if !ok {
		code = 0
	}

	return fmt.Sprintf("\u001B[%dm%s\u001B[0m", code, str)
}

func Faint(str string) string {
	return wrapSGR(str, "faint")
}

func Bold(str string) string {
	return wrapSGR(str, "bold")
}

func Emph(str string) string {
	return wrapSGR(str, "italicized")
}

func Underline(str string) string {
	return wrapSGR(str, "underlined")
}

func Blink(str string) string {
	return wrapSGR(str, "blink")
}

func Inverse(str string) string {
	return wrapSGR(str, "inverse")
}

func Conceal(str string) string {
	return wrapSGR(str, "invisible")
}

func ConcealNegative(str string) string {
	return Inverse(Conceal(str))
}
