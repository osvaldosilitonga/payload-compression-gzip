package controllers

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
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
	log.Info("Starting to download file")

	imgId := c.Param("id")
	filePath := fmt.Sprintf("./assets/server_storage/%v", imgId)

	// check if the file exists
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return c.JSON(http.StatusNotFound, echo.Map{
			"code":    http.StatusNotFound,
			"message": "file not found",
		})
	} else if err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	log.Info(fmt.Sprintf("Serving file: %s", filePath))
	return c.Attachment(filePath, imgId)
}

type Data struct {
	ImageName string `json:"image_name"`
	Owner     string `json:"owner"`
	Email     string `json:"email"`
}

func (i *Image) GetAll(c echo.Context) error {
	datas := []Data{
		{
			ImageName: "44b45cc9-5db6-4034-933e-7cf30a93bc44_samurai.png",
			Owner:     "John Doe",
			Email:     "john.doe@mail.com",
		},
		{
			ImageName: "17240e88-73a4-4f51-bd7a-8afa0944c2bc_samurai.png",
			Owner:     "Jane Smith",
			Email:     "jane.smith@mail.com",
		},
		{
			ImageName: "f52b271e-430c-415d-a907-134d62e47b77_samurai.png",
			Owner:     "Eko Kurniawan",
			Email:     "eko.kurniawan@mail.com",
		},
	}

	return c.JSON(200, echo.Map{
		"code":    http.StatusOK,
		"message": "ok",
		"data":    datas,
	})
}
