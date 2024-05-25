package authController

import (
	"net/http"

	userModel "github.com/NabinGrz/SocialMediaApi/src/authentication/models"
	authServices "github.com/NabinGrz/SocialMediaApi/src/authentication/services"
	"github.com/gin-gonic/gin"
)

var tokenString string

func Login(ctx *gin.Context) {
	var user userModel.User

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	token, error := authServices.Login(user.Email, user.Password)

	if error != nil {
		ctx.JSON(http.StatusBadRequest, error)
		return
	}

	ctx.IndentedJSON(http.StatusOK, token)

}

func Register(ctx *gin.Context) {
	var user userModel.User

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	_, err := authServices.Register(user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.IndentedJSON(http.StatusOK, gin.H{"message": "User registered successfully"})
}

func Authenticate(c *gin.Context) {
	tokenString = c.GetHeader("Authorization")
	if tokenString == "" {
		c.IndentedJSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized user",
		})
		c.Abort()
		return
	}

	tokenString = tokenString[len("Bearer "):]
	c.Next()
}
