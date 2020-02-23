package handler

import (
	"encoding/json"
	"net/http"

	echo "github.com/labstack/echo/v4"
)

type HTTPError struct {
	code int
	msg  string
}

func (e *HTTPError) Error() string {
	return e.msg
}

func (e *HTTPError) MarshalJSON() ([]byte, error) {
	return json.Marshal(echo.Map{
		"message": e.Error(),
	})
}

func ErrorHandler(err error, c echo.Context) {
	code := http.StatusInternalServerError
	if e, ok := err.(*HTTPError); ok {
		code = e.code
	}

	c.JSON(code, echo.Map{
		"message": err.Error(),
	})
}
