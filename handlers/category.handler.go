package handlers

import (
	"fmt"
	"net/http"
	"products/models"
	"products/queries"

	"github.com/gin-gonic/gin"
)

var category *models.Category

func GetCategories(c *gin.Context) {
	categories, err := queries.GetCategoiesQuery()
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
	}

	fmt.Println("GetCategories", categories)

	c.JSON(http.StatusOK, categories)
}

func CreateCategory(c *gin.Context) {
	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	err := queries.CreateCategoryQuery(category)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{ "message": err.Error() })
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Create category SUCCESS!"})
}

func GetCategoryById(c *gin.Context) {
	id := c.Param("id")

	categoryById, err := queries.GetCategoryByIdQuery(&id)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
	}

	c.JSON(http.StatusOK, categoryById)
}

func UpdateCategoryById(c *gin.Context) {
	id := c.Param("id")
	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}

	err := queries.UpdateCategoryByIdQuery(&id, category)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{ "message": err.Error() })
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Update category SUCCESS!"})
}

func DeleteCategoryById(c *gin.Context) {
	id := c.Param("id")
	err := queries.DeleteCategoryByIdQuery(&id)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{ "message": err.Error() })
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Delete category SUCCESS!"})
}