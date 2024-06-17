package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/osvaldosilitonga/payload-compression-gzip/middlewares"
	"github.com/osvaldosilitonga/payload-compression-gzip/routes"
)

func main() {
	e := echo.New()

	// Logrus log
	e.Use(middleware.RequestLoggerWithConfig(middlewares.LogrusConfig()))
	e.Use(middleware.Recover())

	routes.Router(e)

	e.Logger.Fatal(e.Start(":8080"))
}
