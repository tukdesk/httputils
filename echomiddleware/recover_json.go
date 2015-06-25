package echomiddleware

import (
	"runtime/debug"

	"github.com/tukdesk/httputils/jsonutils"

	"github.com/labstack/echo"
)

func JSONRecoverForAPIError() echo.Middleware {
	return func(h echo.HandlerFunc) echo.HandlerFunc {
		return func(c *echo.Context) error {
			logger := GetRequestLogger(c)
			defer func() {
				if err := recover(); err != nil {
					switch err.(type) {
					case *jsonutils.APIError, *echo.HTTPError:
						c.Error(err.(error))
					default:
						logger.Error(err)
						debug.PrintStack()
						JSONOutputAPIError(ErrInternalServerError, c)
					}
				}
			}()
			return h(c)
		}
	}
}
