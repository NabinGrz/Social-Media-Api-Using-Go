package main

import (
	"net/http"

	authController "github.com/NabinGrz/SocialMediaApi/src/authentication/controller"
	"github.com/NabinGrz/SocialMediaApi/src/router"
	"github.com/gin-gonic/gin"
)

func main() {

	r := router.Router()

	authorized := r.Group("/api")
	authorized.Use(authController.Authenticate)
	{
		authorized.GET("/", func(ctx *gin.Context) {
			ctx.IndentedJSON(http.StatusOK, "Hello World")
		})
	}

	r.Run()

}
