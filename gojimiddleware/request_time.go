package gojimiddleware

import (
	"fmt"
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

		if lw.Status() == 0 {
			lw.WriteHeader(http.StatusOK)
		}

		logger.Infof("[%d] %s | %s | %s", lw.Status(), r.Method, r.URL.String(), SinceStr(start))
	}

	return http.HandlerFunc(fn)
}

var (
	intSecond      = int64(time.Second)
	intMillisecond = int64(time.Millisecond)
	intMicrosecond = int64(time.Microsecond)
	intNanosecond  = int64(time.Nanosecond)

	floatSecond      = float64(time.Second)
	floatMillisecond = float64(time.Millisecond)
	floatMicrosecond = float64(time.Microsecond)
	floatNanosecond  = float64(time.Nanosecond)
)

func SinceStr(then time.Time) string {
	duration := time.Now().UnixNano() - then.UnixNano()
	switch {
	case duration > intSecond:
		return fmt.Sprintf("%.2f s", float64(duration)/floatSecond)
	case duration > intMillisecond:
		return fmt.Sprintf("%.2f ms", float64(duration)/floatMillisecond)
	case duration > intMicrosecond:
		return fmt.Sprintf("%.2f us", float64(duration)/floatMicrosecond)
	default:
		return fmt.Sprintf("%.2f ns", float64(duration)/floatNanosecond)
	}
}
