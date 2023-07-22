package routes

import (
	"products/handlers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine) {
	router.POST("/user", handlers.CreateUser)
	router.PATCH("/user/:id", handlers.UpdateUser)
	router.DELETE("/user/:id", handlers.DeleteUser)
	router.GET("/user/:id", handlers.GetUserById)
	router.GET("/users", handlers.GetUsers)
}