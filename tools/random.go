package tools

import (
	"math/rand"
	"time"
)

var (
	letters    = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	alphadigit = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
)

func RandInt(min int, max int) int {
	if max <= min {
		return min
	}
	rand.Seed(time.Now().UTC().UnixNano())
	return min + rand.Intn(max-min)
}

func RandStringFromSrc(l int, src []rune) string {
	rand.Seed(time.Now().UnixNano())
	srcLength := len(src)
	data := make([]rune, l)

	for i := range data {
		data[i] = src[rand.Intn(srcLength)]
	}
	return string(data)
}

func RandString(l int) string {
	return RandStringFromSrc(l, alphadigit)
}
