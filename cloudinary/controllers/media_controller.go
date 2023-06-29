package controllers

import (
	"myapp/cloudinary/models"
	"myapp/cloudinary/services"
	"myapp/dtos"
	"net/http"

	echo "github.com/tbxark/g4vercel"
)

func Status(c *echo.Context) {
	c.JSON(
		http.StatusOK,
		dtos.MediaDto{
			StatusCode: http.StatusOK,
			Message:    "success",
			Data:       &echo.H{"data": "OK"},
		})
}

/*
Upload file
*/
func FileUpload(c *echo.Context) {
	_, formHeader, err := c.Req.FormFile("file")
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			dtos.MediaDto{
				StatusCode: http.StatusInternalServerError,
				Message:    "error",
				Data:       &echo.H{"data": "Select a file to upload"},
			})
	}

	//get file from header
	formFile, err := formHeader.Open()

	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			dtos.MediaDto{
				StatusCode: http.StatusInternalServerError,
				Message:    "error",
				Data:       &echo.H{"data": err.Error()},
			})
	}

	uploadUrl, err := services.NewMediaUpload().FileUpload(models.File{File: formFile})
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			dtos.MediaDto{
				StatusCode: http.StatusInternalServerError,
				Message:    "error",
				Data:       &echo.H{"data": err.Error()},
			})
	}

	c.JSON(
		http.StatusOK,
		dtos.MediaDto{
			StatusCode: http.StatusOK,
			Message:    "success",
			Data:       &echo.H{"data": uploadUrl},
		})
}
