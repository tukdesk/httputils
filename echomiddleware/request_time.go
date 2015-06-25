package echomiddleware

import (
	"time"

	"github.com/labstack/echo"
)

func RequestTimer() echo.MiddlewareFunc {
	return func(h echo.HandlerFunc) echo.HandlerFunc {
		return func(c *echo.Context) error {
			logger := GetRequestLogger(c)

			start := time.Now()

			if err := h(c); err != nil {
				c.Error(err)
			}

			stop := time.Now()

			req := c.Request()
			resp := c.Response()
			logger.SimpleInfof("[%d] %s %s | %s | %d", resp.Status(), req.Method, req.URL.String(), stop.Sub(start), resp.Size())
			return nil
		}
	}

}
