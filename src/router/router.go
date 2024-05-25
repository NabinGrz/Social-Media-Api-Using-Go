package router

import (
	"net/http"

	authController "github.com/NabinGrz/SocialMediaApi/src/authentication/controller"
	"github.com/NabinGrz/SocialMediaApi/src/profile"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	router := gin.Default()

	//!! GENERATING JWT TOKEN AFTER LOGIN
	router.POST("/login", authController.Login)
	router.POST("/register", authController.Register)

	authorized := router.Group("/api")
	authorized.Use(authController.Authenticate)
	{
		authorized.GET("/", func(ctx *gin.Context) {
			ctx.IndentedJSON(http.StatusOK, "Hello World")
		})
		authorized.PUT("/updateProfileUrl/:id", profile.UpdateProfileImage)
		authorized.PUT("/updateProfileDetail/:id", profile.UpdateDetails)
	}

	return router
}
