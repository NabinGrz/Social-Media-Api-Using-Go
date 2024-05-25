package cloudinaryController

import (
	"net/http"

	cloudinaryService "github.com/NabinGrz/SocialMediaApi/src/cloudinary/service"
	"github.com/gin-gonic/gin"
)

func UploadFile(c *gin.Context) {
	//!! Get the image from request body
	file, err := c.FormFile("file")

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//!! Upload the image locally
	err = c.SaveUploadedFile(file, "assets/uploads/"+file.Filename)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
		return
	}

	//!! Using UploadToCloudinary Function
	fileUrl, err := cloudinaryService.UploadCloudinary(file, "file")

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to upload file"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":  "Upload successful",
		"imageUrl": fileUrl,
	})

}
