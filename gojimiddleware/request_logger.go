package gojimiddleware

import (
	"net/http"

	"github.com/zenazn/goji/web"
)

const (
	RequestLoggerKey = "_logger"
)

func RequestLogger(c *web.C, h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		GetRequestLogger(c, w, r)
		h.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}

func GetRequestLogger(c *web.C, w http.ResponseWriter, r *http.Request) *xlogger {
	if c.Env == nil {
		c.Env = map[interface{}]interface{}{}
	}

	if logger, ok := c.Env[RequestLoggerKey].(*xlogger); ok {
		return logger
	}

	logger := newXLogger(w, r)
	c.Env[RequestLoggerKey] = logger

	return logger
}
