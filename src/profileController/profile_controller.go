package profileController

import (
	"net/http"

	userModel "github.com/NabinGrz/SocialMediaApi/src/authentication/models"
	cloudinaryController "github.com/NabinGrz/SocialMediaApi/src/cloudinary/controller"
	"github.com/NabinGrz/SocialMediaApi/src/dbConnection"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func UpdateDetails(c *gin.Context) {
	var updatedProfile userModel.User
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	if err := c.ShouldBindJSON(&updatedProfile); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Something went wrong"})
		return
	}

	objID, idErr := primitive.ObjectIDFromHex(id)
	if idErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	filter := bson.M{"_id": objID}
	update := bson.M{"$set": bson.M{"username": updatedProfile.Username, "fullname": updatedProfile.FullName}}

	result, err := dbConnection.UserCollection.UpdateOne(c, filter, update)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update profile URL"})
		return
	}

	if result.MatchedCount == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Profile Detail updated successfully"})
}
func UpdateProfileImage(c *gin.Context) {
	// var updatedProfile userModel.User

	fileUrl, err := cloudinaryController.UploadFile(c)
	id := c.Param("id")

	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	filter := bson.M{"_id": objID}
	update := bson.M{"$set": bson.M{"profileurl": fileUrl}}

	result, err := dbConnection.UserCollection.UpdateMany(c, filter, update)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update profile URL"})
		return
	}

	if result.MatchedCount == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Profile URL updated successfully"})
}
