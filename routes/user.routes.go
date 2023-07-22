package routes

import (
	"products/handlers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine) {
	router.GET("/users", handlers.GetUsers)
	router.POST("/user", handlers.CreateUser)
	router.GET("/user/:id", handlers.GetUserById)
	router.PATCH("/user/:id", handlers.UpdateUserById)
	router.DELETE("/user/:id", handlers.DeleteUserById)
}