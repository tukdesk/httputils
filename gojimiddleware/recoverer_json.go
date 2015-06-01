package gojimiddleware

import (
	"net/http"
	"runtime/debug"

	"github.com/tukdesk/httputils/jsonutils"
	"github.com/zenazn/goji/web"
)

var (
	ErrInternalServerError = jsonutils.NewAPIError(http.StatusInternalServerError, http.StatusInternalServerError, "internal server error")
)

func RecovererJson(c *web.C, h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		logger := GetRequestLogger(c, w, r)

		defer func() {
			if err := recover(); err != nil {
				// maybe use panic to abort request
				if apierr, ok := err.(*jsonutils.APIError); ok {
					jsonutils.OutputJsonError(apierr, w, r)
					return
				}

				logger.Error(err)
				debug.PrintStack()
				jsonutils.OutputJsonError(ErrInternalServerError, w, r)
			}
		}()

		h.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}
