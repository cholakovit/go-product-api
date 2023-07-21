package routes

import (
	"log"
	"os"
	"products/handlers"

	"github.com/gin-gonic/gin"
)

type Routes struct {
}

func (r *Routes) InitRoutes() {
	router := gin.Default()

	r.ProductRoutes(router)

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT environment variable is not set")
	}

	router.Run("localhost:" + port)
}

func (r *Routes) ProductRoutes(router *gin.Engine) {
	router.GET("/products", handlers.GetProducts)
	router.POST("/products", handlers.CreateProduct)
	router.GET("/products/:id", handlers.GetProductById)
	router.PATCH("/products/:id", handlers.UpdateProduct)
	router.DELETE("/products/:id", handlers.DeleteProduct)
}