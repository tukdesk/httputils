package jsonutils

import (
	"bytes"
	"encoding/json"
	"net/http"
)

func outputJsonData(w http.ResponseWriter, status int, body []byte) {
	w.Header().Set("Content-Type", "application/json")
	if status > 0 {
		w.WriteHeader(status)
	}
	if len(body) > 0 {
		w.Write(body)
	}
	return
}

func OutputJson(data interface{}, w http.ResponseWriter, r *http.Request) {
	if data == nil {
		return
	}

	buf := new(bytes.Buffer)
	encoder := json.NewEncoder(buf)
	err := encoder.Encode(data)
	if err != nil {
		OutputJsonError(err, w, r)
		return
	}

	outputJsonData(w, http.StatusOK, buf.Bytes())
	return
}

func OutputJsonError(err error, w http.ResponseWriter, r *http.Request) {
	if err == nil {
		return
	}

	apierr, ok := err.(*APIError)
	if !ok {
		apierr = NewInternalError(err)
	}

	outputJsonData(w, apierr.StatusCode, apierr.Bytes())
	return
}
