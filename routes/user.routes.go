package routes

import (
	"products/handlers"
	//"products/middleware"

	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine) {
	//router.Use(middleware.Auth)
	router.GET("/users", handlers.GetUsers)
	router.POST("/user", handlers.CreateUser)
	router.GET("/user/:id", handlers.GetUserById)
	router.PATCH("/user/:id", handlers.UpdateUserById)
	router.DELETE("/user/:id", handlers.DeleteUserById)
}