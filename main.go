package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/osvaldosilitonga/payload-compression-gzip/middlewares"
	"github.com/osvaldosilitonga/payload-compression-gzip/routes"
)

func main() {
	e := echo.New()

	e.Use(middleware.RequestLoggerWithConfig(middlewares.LogrusConfig()))
	e.Use(middleware.Recover())
	e.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Level:     5,
		MinLength: 1000,
	}))

	routes.Router(e)

	e.Logger.Fatal(e.Start(":8080"))
}
