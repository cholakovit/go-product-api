package handlers

import (
	"net/http"
	"products/helpers"
	"products/models"
	"products/queries"
	vs "products/validationMessages"
	"time"

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

	emailMsg := helpers.VerifyEmail(&user.Email)
	if emailMsg != "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": emailMsg})
		return
	}

	token, refreshToken, _ := helpers.GenerateToken(user)

	user.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	user.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	user.Token = token
	user.Rtoken = refreshToken

	pass := helpers.HashPassword(user.Pass)
	user.Pass = pass

	err := queries.CreateUserQuery(user)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{ "message": err.Error() })
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "create user success"})
}

func GetUserById(c *gin.Context) {
	id := c.Param("id")

	userById, err := queries.GetUserByIdQuery(&id)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
	}

	c.JSON(http.StatusOK, userById)
}

func UpdateUserById(c *gin.Context) {
	id := c.Param("id")
	if err := c.ShouldBindJSON(&user); err != nil {
		errMsg := vs.UserMessageValidate(err)
		c.JSON(http.StatusBadRequest, gin.H{"message": errMsg})
		return
	}

	userById, err := queries.GetUserByIdQuery(&id)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
	}

	if user.Email != userById.Email {
		isValidEmail := helpers.VerifyEmail(&user.Email)
		if isValidEmail != "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": isValidEmail})
			return		
		}
	} 

	user.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

	err = queries.UpdateUserByIdQuery(&id, user)
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