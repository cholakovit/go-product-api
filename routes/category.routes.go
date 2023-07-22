package routes

import (
	"products/handlers"

	"github.com/gin-gonic/gin"
)

func CategoryRoutes(router *gin.Engine) {
	router.GET("/categories",  handlers.GetCategories)
	router.POST("/category", handlers.CreateCategory)
	router.GET("/category/:id", handlers.GetCategoryById)
	router.PATCH("/category/:id", handlers.UpdateCategoryById)
	router.DELETE("/category/:id", handlers.DeleteCategoryById)
}