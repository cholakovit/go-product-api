package handlers

import (
	"net/http"
	"products/helpers"
	"products/models"
	"products/queries"

	"github.com/gin-gonic/gin"
)

func SignIn(c *gin.Context) {
	var auth models.Auth

	if err := c.BindJSON(&auth); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	foundUser, err := queries.FindOneQuery(&auth)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	passwordIsValid, msg := helpers.VerifyPassword(*auth.Pass, foundUser.Pass)
	if passwordIsValid != true {
		c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
		return
	}
	c.JSON(http.StatusOK, foundUser)
}