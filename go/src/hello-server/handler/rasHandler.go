package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetRasHandler(c echo.Context) error {
	return c.String(http.StatusOK, "GET!\nRasです、よろしく。\n")
}
