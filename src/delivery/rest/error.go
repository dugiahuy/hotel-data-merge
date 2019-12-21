package rest

import (
	"net/http"

	"github.com/labstack/echo"

	"github.com/dugiahuy/hotel-data-merge/src/util/err_util"
)

type customResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Detail  string `json:"detail"`
}

func customHTTPErrorHandler(err error, c echo.Context) {
	if err == echo.ErrNotFound {
		noContent := map[string]string{"message": "url not found"}
		c.JSON(404, noContent)
		return
	}
	code := http.StatusInternalServerError
	if sc, ok := err.(err_util.StatusCoder); ok {
		code = sc.StatusCode()
	}
	if !c.Response().Committed {
		if c.Request().Method == http.MethodHead {
			err = c.NoContent(code)
		} else {
			rsp := &customResponse{}
			rsp.Code = code
			rsp.Message = err.Error()
			if code == 500 {
				rsp.Message = "internal server error"
			}
			rsp.Detail = err.Error()

			// log stuff
			c.JSON(code, rsp)

			return
		}
	}
}
