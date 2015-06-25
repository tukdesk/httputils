package echomiddleware

import (
	"net/http"

	"github.com/tukdesk/httputils/jsonutils"

	"github.com/labstack/echo"
)

var (
	ErrInternalServerError = jsonutils.NewAPIError(http.StatusInternalServerError, http.StatusInternalServerError, "internal server error")
)

func JSONErrHandlerForAPIError() echo.HTTPErrorHandler {
	return func(err error, c *echo.Context) {
		apierr, ok := err.(*jsonutils.APIError)
		if ok {
			JSONOutputAPIError(apierr, c)
			return
		}

		if echoErr, ok := err.(*echo.HTTPError); ok {
			code := echoErr.Code()
			JSONOutputAPIError(jsonutils.NewAPIError(code, code, echoErr.Error()), c)
			return
		}

		logger := GetRequestLogger(c)
		logger.Error(err)
		JSONOutputAPIError(ErrInternalServerError, c)
		return
	}
}

func JSONOutputAPIError(err *jsonutils.APIError, c *echo.Context) {
	c.JSON(err.StatusCode, err)
	return
}
