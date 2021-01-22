package main

import (
	"github.com/labstack/echo/v4"
	"github.com/Ras96/backend-practice/go/src/hello-server/handler"
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
	e.GET("/json", handler.GetJsonHandler)
	e.GET("/fizzbuzz", handler.GetFizzbuzzHandler)
	e.GET("/students/:class/:studentNumber", handler.GetStudentNumberHandler)
	e.POST("/add", handler.GetAddHandler)

	e.Logger.Fatal(e.Start(":4000"))
}
