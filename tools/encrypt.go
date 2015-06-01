package tools

import (
	"crypto/sha1"
	"fmt"
	"io"
)

func Encrypt(text string, salt string) string {
	t := sha1.New()
	if salt == "" {
		salt = text[:len(text)-1]
	}
	io.WriteString(t, text)
	io.WriteString(t, salt)
	return fmt.Sprintf("%x", t.Sum(nil))
}
