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
	// Context 被复用, store 中已有 logger, reset 后直接返回
	if logger, ok := c.Get(RequestLoggerKey).(*xlogger.XLogger); ok {
		logger.ResetForRequest(c.Response().Writer(), c.Request())
		return logger
	}

	// Context 为新创建的对象, set 一个 logger 到 store 中
	logger := xlogger.NewXLogger(c.Response().Writer(), c.Request())
	c.Set(RequestLoggerKey, logger)

	return logger
}
