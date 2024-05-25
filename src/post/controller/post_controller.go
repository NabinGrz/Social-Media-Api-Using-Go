package postController

import (
	"context"
	"net/http"
	"time"

	"github.com/NabinGrz/SocialMediaApi/src/dbConnection"
	postModel "github.com/NabinGrz/SocialMediaApi/src/post/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreatePost(c *gin.Context) {
	var post postModel.SocialMediaPost
	userID, _ := primitive.ObjectIDFromHex(c.GetString("userid"))
	if err := c.BindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	post.Date = time.Now()
	post.User = userID

	result, uploadError := dbConnection.PostCollection.InsertOne(c, post)
	if uploadError != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": uploadError.Error()})
		return
	}

	if result.InsertedID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Post created successfully"})
}
func UpdatePost(c *gin.Context) {
	userID, _ := primitive.ObjectIDFromHex(c.GetString("userid"))

	var updatedPost postModel.SocialMediaPost
	var foundPost postModel.SocialMediaPost

	id := c.Param("id")

	objID, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": objID}

	result := dbConnection.PostCollection.FindOne(context.Background(), filter)

	if result.Err() != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}

	err := c.ShouldBindJSON(&updatedPost)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	result.Decode(&foundPost)

	if userID != foundPost.User {
		c.JSON(http.StatusNotFound, gin.H{"error": "You cannot update others post"})
		return
	}
	update := bson.M{"$set": bson.M{"caption": updatedPost.Caption}}
	updateResult, _ := dbConnection.PostCollection.UpdateMany(context.Background(), filter, update)

	if updateResult.MatchedCount == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Post updated successfully"})
}
func DeletePost(c *gin.Context) {
	userID, _ := primitive.ObjectIDFromHex(c.GetString("userid"))

	var post postModel.SocialMediaPost

	id := c.Param("id")

	objID, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": objID}

	err := dbConnection.PostCollection.FindOne(context.Background(), filter).Decode(&post)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	if userID != post.User {
		c.JSON(http.StatusNotFound, gin.H{"error": "You cannot delete others post"})
		return
	}
	result, _ := dbConnection.PostCollection.DeleteOne(context.Background(), filter)

	if result.DeletedCount == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Post deleted successfully"})
}