package routes

import (
	"products/handlers"

	"github.com/gin-gonic/gin"
)

func ProductRoutes(router *gin.Engine) {
	router.GET("/products", handlers.GetProducts)
	router.POST("/product", handlers.CreateProduct)
	router.GET("/product/:id", handlers.GetProductById)
	router.PATCH("/product/:id", handlers.UpdateProductById)
	router.DELETE("/product/:id", handlers.DeleteProductById)
}