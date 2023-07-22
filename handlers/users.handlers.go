package handlers

import (
	"fmt"
	"net/http"
	"products/models"

	"products/queries"

	"github.com/gin-gonic/gin"
)

var user *models.User

func CreateUser(c *gin.Context) {
  
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	fmt.Println("CreateUser")
	err := queries.CreateUserQuery(user)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{ "message": err.Error() })
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "create user success"})
}

func UpdateUser(c *gin.Context) {
	id := c.Param("id")
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	err := queries.UpdateUserQuery(&id, user)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "update user success"})
}

func DeleteUser(c *gin.Context) {
	id := c.Param("id")

	err := queries.DeleteUserQuery(&id)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{ "message": err.Error() })
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "delete user success"})
}

func GetUserById(c *gin.Context) {
	id := c.Param("id")

	productById, err := queries.GetUserByIdQuery(&id)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
	}

	c.JSON(http.StatusOK, productById)
}

func GetUsers(c *gin.Context) {
	users, err := queries.GetAllUsersQuery()
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
	}	

	c.JSON(http.StatusOK, users)
}