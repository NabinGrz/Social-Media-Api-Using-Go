package postController

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/NabinGrz/SocialMediaApi/src/dbConnection"
	postModel "github.com/NabinGrz/SocialMediaApi/src/post/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetAllPost(c *gin.Context) {
	var allPost []postModel.SocialMediaPost
	cursor, err := dbConnection.PostCollection.Find(context.Background(), bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer cursor.Close(c)
	for cursor.Next(c) {
		var post postModel.SocialMediaPost

		if err := cursor.Decode(&post); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		allPost = append(allPost, post)
	}

	if err := cursor.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, allPost)
}
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

func LikePost(c *gin.Context) {
	userID, _ := primitive.ObjectIDFromHex(c.GetString("userid"))
	var foundPost postModel.SocialMediaPost

	id := c.Param("id")

	objID, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": objID}

	result := dbConnection.PostCollection.FindOne(context.Background(), filter)

	if result.Err() != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}
	result.Decode(&foundPost)
	update := bson.M{"$addToSet": bson.M{"likeby": userID}}
	updateResult, _ := dbConnection.PostCollection.UpdateMany(context.Background(), filter, update)
	fmt.Println(updateResult)

	c.JSON(http.StatusOK, gin.H{"message": "Post has been liked successfully"})
}
func SharePost(c *gin.Context) {
	userID, _ := primitive.ObjectIDFromHex(c.GetString("userid"))
	var foundPost postModel.SocialMediaPost

	id := c.Param("id")

	objID, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": objID}

	result := dbConnection.PostCollection.FindOne(context.Background(), filter)

	if result.Err() != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}
	result.Decode(&foundPost)
	update := bson.M{"$addToSet": bson.M{"shares": userID}}
	updateResult, _ := dbConnection.PostCollection.UpdateMany(context.Background(), filter, update)
	fmt.Println(updateResult)
	c.JSON(http.StatusOK, gin.H{"message": "Post has been liked successfully"})
}
func CommentPost(c *gin.Context) {
	userID, _ := primitive.ObjectIDFromHex(c.GetString("userid"))
	var foundPost postModel.SocialMediaPost
	var comment postModel.CommentDetail

	id := c.Param("id")

	objID, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": objID}

	result := dbConnection.PostCollection.FindOne(context.Background(), filter)

	if result.Err() != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}
	result.Decode(&foundPost)
	err := c.ShouldBindJSON(&comment)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if comment.Comment == "" {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"error": "Please enter your comment",
		})
		return
	}
	comment.User = userID
	comment.Date = time.Now()

	update := bson.M{"$push": bson.M{
		"comments": comment,
	}}
	updateResult, _ := dbConnection.PostCollection.UpdateOne(context.Background(), filter, update)
	fmt.Println(updateResult)
	c.JSON(http.StatusOK, gin.H{"message": "Post has been commented successfully"})
}
