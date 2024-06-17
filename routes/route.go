package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/osvaldosilitonga/payload-compression-gzip/controllers"
)

func Router(e *echo.Echo) {
	imagesController := controllers.NewImageController()

	v1 := e.Group("/api/v1")

	imageV1 := v1.Group("/images")
	{
		imageV1.GET("", imagesController.GetAll)
		imageV1.POST("/upload", imagesController.Upload)
		imageV1.GET("/download/:id", imagesController.Download)
	}
}
