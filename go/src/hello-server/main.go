package main

import (
	"github.com/labstack/echo/v4"
	"github.com/Ras96/backend-practice/handler"
)

func main() {
	e := echo.New()

	hello := e.Group("/hello")
	{
		hello.GET("", handler.GetHelloHandler)
		hello.POST("", handler.PostHelloHandler)
		hello.GET("/:username", handler.GetUsernameHandler)
	}
	e.GET("/ping", handler.GetPingHandler)
	e.GET("/ras", handler.GetRasHandler)
	e.GET("/json", handler.jsonHandler)
	e.GET("/fizzbuzz", handler.fizzbuzzHandler)
	e.GET("/students/:class/:studentNumber", handler.GetStudentNumberHandler)
	e.POST("/add", handler.addHandler)

	e.Logger.Fatal(e.Start(":4000"))
}
