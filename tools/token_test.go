package tools

import (
	"testing"
	"time"
)

func TestGenerateToken(t *testing.T) {
	data := map[string]interface{}{"userId": "abc"}
	str := GenerateToken(data, 1, []byte("abc"))

	token, err := ParseToken(str, []byte("abc"))
	if err != nil {
		t.Error(err)
	}

	if !token.Valid {
		t.Error("not valid")
	}

	time.Sleep(time.Second * 2)
	_, err = ParseToken(str, []byte("abc"))
	if err != ErrTokenExpired {
		t.Error("expected to be expired")
	}
}
