package routes

import (
	"log"
	"os"
	"products/controllers"

	"github.com/gin-gonic/gin"
)

var(
	pc		controllers.ProductControllers
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
	router.GET("/products", pc.GetProducts)
	router.POST("/products", pc.CreateProduct)
	router.GET("/products/:id", pc.GetProductById)
	router.PATCH("/products/:id", pc.UpdateProduct)
	router.DELETE("/products/:id", pc.DeleteProduct)
}