package main

import (
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func main() {
	e := echo.New()

	e.File("/swagger/raw", "pkg/api/openapi.yml")
	url := echoSwagger.URL("/swagger/raw")
	e.GET("/swagger/*", echoSwagger.EchoWrapHandler(url))
	e.Logger.Fatal(e.Start(":1323"))
}
