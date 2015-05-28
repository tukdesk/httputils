package jsonutils

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type APIError struct {
	StatusCode int    `json:"-"`
	ErrorCode  int    `json:"error_code"`
	ErrorMsg   string `json:"error_msg"`
}

func (this *APIError) Error() string {
	return fmt.Sprintf("StatusCode: %d; ErrorCode: %d; Error: %s", this.StatusCode, this.ErrorCode, this.ErrorMsg)
}

func (this *APIError) Bytes() []byte {
	b, _ := json.Marshal(this)
	return b
}

func NewAPIError(statusCode, errorCode int, msg string) *APIError {
	return &APIError{
		StatusCode: statusCode,
		ErrorCode:  errorCode,
		ErrorMsg:   msg,
	}
}

func NewInternalError(err error) *APIError {
	return NewAPIError(http.StatusInternalServerError, http.StatusInternalServerError, err.Error())
}
