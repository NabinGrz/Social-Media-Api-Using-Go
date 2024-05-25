package cloudinaryController

import (
	"errors"

	cloudinaryService "github.com/NabinGrz/SocialMediaApi/src/cloudinary/service"
	"github.com/gin-gonic/gin"
)

func UploadFile(c *gin.Context) (string, error) {
	//!! Get the image from request body
	file, err := c.FormFile("file")

	if err != nil {
		return "", err
	}

	//!! Upload the image locally
	err = c.SaveUploadedFile(file, "assets/uploads/"+file.Filename)

	if err != nil {
		return "", errors.New("Failed to save file")
	}

	//!! Using UploadToCloudinary Function
	fileUrl, err := cloudinaryService.UploadCloudinary(file, "file")

	if err != nil {
		return "", errors.New("Failed to upload file")
	}

	// c.JSON(http.StatusOK, gin.H{
	// 	"message":  "Upload successful",
	// 	"imageUrl": fileUrl,
	// })

	return fileUrl, nil

}
