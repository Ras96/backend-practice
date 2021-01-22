package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Add struct {
	Right int `json:"right"`
	Left  int `json:"left"`
}

func GetAddHandler(c echo.Context) error {
	a := new(Add)
	if err := c.Bind(&a); err != nil {
		return c.JSON(
			http.StatusBadRequest,err,
			// map[string]string{"error": "Bad Request"},
		)
	}

	return c.JSON(
		http.StatusOK,
		map[string]int{"answer": a.Right + a.Left},
	)
}
