package router

import (
	"net/http"

	authController "github.com/NabinGrz/SocialMediaApi/src/authentication/controller"
	authServices "github.com/NabinGrz/SocialMediaApi/src/authentication/services"
	postController "github.com/NabinGrz/SocialMediaApi/src/post/controller"
	"github.com/NabinGrz/SocialMediaApi/src/profileController"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	router := gin.Default()

	//!! GENERATING JWT TOKEN AFTER LOGIN
	router.POST("/login", authController.LoginHandler)
	router.POST("/register", authController.RegisterHandler)

	router.Use(authServices.JWTMiddleware())

	authorized := router.Group("/api")
	authorized.Use(authController.AuthenticateHandler)
	{
		authorized.GET("/", func(ctx *gin.Context) {
			ctx.IndentedJSON(http.StatusOK, "Hello World")
		})
		authorized.PUT("/updateProfileUrl/:id", profileController.UpdateProfileImage)
		authorized.PUT("/updateProfileDetail/:id", profileController.UpdateDetails)
		authorized.POST("/post", postController.CreatePost)
		authorized.DELETE("/post/:id", postController.DeletePost)
	}

	return router
}
