package handler

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type User struct {
	TraqID string `json:"traq_id"`
}

func GetHelloHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World.\n")
}

func PostHelloHandler(c echo.Context) error {
	user := new(User)
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, fmt.Errorf("failed to bind the request.: %v", err))
	}
	return c.String(http.StatusOK, fmt.Sprintf("POST!\n%vさん、よろしく。\n", user.TraqID))
}

func GetUsernameHandler(c echo.Context) error {
	userID := c.Param("username")
	return c.String(http.StatusOK, "Hello, "+userID+".\n")
}
