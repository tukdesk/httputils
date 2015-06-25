package gojimiddleware

import (
	"net/http"
	"time"

	"github.com/zenazn/goji/web"
	"github.com/zenazn/goji/web/mutil"
)

func RequestTimer(c *web.C, h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		logger := GetRequestLogger(c, w, r)

		lw := mutil.WrapWriter(w)

		start := time.Now()
		h.ServeHTTP(lw, r)
		stop := time.Now()

		if lw.Status() == 0 {
			lw.WriteHeader(http.StatusOK)
		}

		logger.SimpleInfof("[%d] %s %s | %s | %d", lw.Status(), r.Method, r.URL.String(), stop.Sub(start), lw.BytesWritten())
	}

	return http.HandlerFunc(fn)
}
