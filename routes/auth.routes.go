package routes

import (
	"products/handlers"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(router *gin.Engine) {
	router.POST("/signin", handlers.SignIn)
}