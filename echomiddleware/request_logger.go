package echomiddleware

import (
	"github.com/tukdesk/httputils/xlogger"

	"github.com/labstack/echo"
)

const (
	RequestLoggerKey = "_logger"
)

func RequestLogger() func(*echo.Context) error {
	return func(c *echo.Context) error {
		GetRequestLogger(c)
		return nil
	}
}

func GetRequestLogger(c *echo.Context) *xlogger.XLogger {
	if logger, ok := c.Get(RequestLoggerKey).(*xlogger.XLogger); ok {
		return logger
	}

	logger := xlogger.NewXLogger(c.Response().Writer(), c.Request())
	c.Set(RequestLoggerKey, logger)

	return logger
}
