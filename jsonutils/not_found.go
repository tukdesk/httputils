package jsonutils

import (
	"fmt"
	"net/http"
)

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	err := NewAPIError(http.StatusNotFound, http.StatusNotFound, fmt.Sprintf("%s %s not found", r.Method, r.URL.String()))
	OutputJsonError(err, w, r)
	return
}
