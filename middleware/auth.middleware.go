package middleware

import (
	"fmt"
	"net/http"
	"products/helpers"

	"github.com/gin-gonic/gin"
)

func Auth(c *gin.Context) {

	token := c.Request.Header.Get("token")
	if token == ""{
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("No Authorization header provided")})
		c.Abort()
		return
	}

	claims, err := helpers.ValidateToken(token)
	if err != "" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		c.Abort()
		return
	}

	c.Set("FullName", claims.Full_name)
	c.Set("email", claims.Email)
	c.Set("role", claims.Role)
	c.Next()
}
