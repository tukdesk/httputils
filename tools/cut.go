package tools

import (
	"strings"
)

func CutRune(s string, l int) (int, string) {
	r := []rune(s)
	length := len(r)

	if l < 0 || l > length {
		return length, s
	}
	return length, string(r[:l])
}

func CutEmail(s string) (string, string) {
	index := strings.Index(s, "@")
	if index != -1 {
		return s[:index], s[index+1:]
	}
	return s, ""
}
