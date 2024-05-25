package router

import (
	authController "github.com/NabinGrz/SocialMediaApi/src/authentication/controller"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	router := gin.Default()

	//!! GENERATING JWT TOKEN AFTER LOGIN
	router.POST("/login", authController.Login)
	router.POST("/register", authController.Register)
	return router
}
