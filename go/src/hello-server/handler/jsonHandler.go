package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type jsonData struct {
	Number int    `json:"number,omitempty"`
	String string `json:"string,omitempty"`
	Bool   bool   `json:"bool,omitempty"`
}

func GetJsonHandler(c echo.Context) error {
	res := jsonData{
		Number: 10,
		String: "hoge",
		Bool:   false,
	}
	return c.JSON(http.StatusOK, &res)
}
