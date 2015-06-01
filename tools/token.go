package tools

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var (
	ErrTokenExpired = fmt.Errorf("token expired")
)

func GenerateToken(data map[string]interface{}, expiration time.Duration, key []byte) string {
	token := jwt.New(jwt.GetSigningMethod("HS256"))
	if data != nil {
		token.Claims = data
	} else {
		token.Claims = make(map[string]interface{})
	}
	token.Claims["exp"] = time.Now().Add(expiration).Unix()
	signed, _ := token.SignedString(key)
	return signed
}

func ParseToken(str string, key []byte) (*jwt.Token, error) {
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		return key, nil
	}
	return ParseTokenWithFunc(str, keyFunc)
}

func ParseTokenWithFunc(str string, keyFunc func(token *jwt.Token) (interface{}, error)) (*jwt.Token, error) {
	token, err := jwt.Parse(str, keyFunc)
	if err == nil {
		return token, nil
	}
	if validationErr, ok := err.(*jwt.ValidationError); ok {
		if validationErr.Errors == jwt.ValidationErrorExpired {
			return token, ErrTokenExpired
		}
	}
	return token, err
}
