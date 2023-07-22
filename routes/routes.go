package routes

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func InitRoutes() {
	router := gin.Default()

	ProductRoutes(router)
	UserRoutes(router)

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT environment variable is not set")
	}

	router.Run("localhost:" + port)
}

