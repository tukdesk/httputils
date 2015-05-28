package jsonutils

import (
	"encoding/json"
	"net/http"
)

func GetJsonArgsFromRequest(r *http.Request, obj interface{}) error {
	if obj == nil {
		return nil
	}

	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()

	return decoder.Decode(obj)
}
