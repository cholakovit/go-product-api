package handlers

import (
	"net/http"
	"products/models"
	"products/queries"
	vs "products/validationMessages"

	"github.com/gin-gonic/gin"
)

var user *models.User

func GetUsers(c *gin.Context) {
	users, err := queries.GetUsersQuery()
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
	}	

	c.JSON(http.StatusOK, users)
}

func CreateUser(c *gin.Context) {
  
	if err := c.ShouldBindJSON(&user); err != nil {
		errMsg := vs.UserMessageValidate(err)
		c.JSON(http.StatusBadRequest, gin.H{"message": errMsg})
		return
	}

	err := queries.CreateUserQuery(user)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{ "message": err.Error() })
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "create user success"})
}

func GetUserById(c *gin.Context) {
	id := c.Param("id")

	productById, err := queries.GetUserByIdQuery(&id)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
	}

	c.JSON(http.StatusOK, productById)
}

func UpdateUserById(c *gin.Context) {
	id := c.Param("id")
	if err := c.ShouldBindJSON(&user); err != nil {
		errMsg := vs.UserMessageValidate(err)
		c.JSON(http.StatusBadRequest, gin.H{"message": errMsg})
		return
	}

	err := queries.UpdateUserByIdQuery(&id, user)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "update user success"})
}

func DeleteUserById(c *gin.Context) {
	id := c.Param("id")

	err := queries.DeleteUserByIdQuery(&id)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{ "message": err.Error() })
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "delete user success"})
}