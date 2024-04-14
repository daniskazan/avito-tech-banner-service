package v1

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func Healthz() echo.HandlerFunc {
	type Response struct {
		Status string `json:"status"`
	}
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, Response{Status: "Ok"})
	}
}
