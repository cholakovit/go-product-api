package routes

import (
	"products/handlers"

	"github.com/gin-gonic/gin"
)

func ProductRoutes(router *gin.Engine) {
	router.GET("/products", handlers.GetProducts)
	router.POST("/products", handlers.CreateProduct)
	router.GET("/products/:id", handlers.GetProductById)
	router.PATCH("/products/:id", handlers.UpdateProduct)
	router.DELETE("/products/:id", handlers.DeleteProduct)
}