package controllers

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type Image struct{}

func NewImageController() *Image {
	return &Image{}
}

func (i *Image) Upload(c echo.Context) error {
	// Read from fields
	name := c.FormValue("name")
	email := c.FormValue("email")

	// -----------
	// Read File
	// -----------

	// Source
	file, err := c.FormFile("image")
	if err != nil {
		return c.JSON(404, echo.Map{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
		})
	}
	src, err := file.Open()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
		})
	}
	defer src.Close()

	// Destination
	uuid := uuid.New()
	fileName := uuid.String() + "_" + file.Filename
	dst, err := os.Create(fmt.Sprintf("./assets/server_storage/%v", fileName))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
		})
	}
	defer dst.Close()

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	return c.JSON(200, echo.Map{
		"code":    http.StatusOK,
		"message": "OK",
		"data": map[string]string{
			"name":       name,
			"email":      email,
			"image_name": fileName,
		},
	})
}

func (i *Image) Download(c echo.Context) error {
	imgId := c.Param("id")

	return c.Attachment(fmt.Sprintf("./assets/server_storage/%v", imgId), imgId)
}
