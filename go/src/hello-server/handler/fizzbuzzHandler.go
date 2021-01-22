package handler

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetFizzbuzzHandler(c echo.Context) error {
	q := c.QueryParam("count")
	cnt, err := strconv.Atoi(q)
	if err != nil {
		return c.String(http.StatusBadRequest, "\"count\" is not integer.")
	}
	result := ""
	for i := 1; i <= cnt; i++ {
		if i%15 == 0 {
			result += "FizzBuzz"
		} else if i%3 == 0 {
			result += "Fizz"
		} else if i%5 == 0 {
			result += "Buzz"
		} else {
			result += strconv.Itoa(i)
		}
		result += "\n"
	}
	return c.String(http.StatusOK, result)
}
