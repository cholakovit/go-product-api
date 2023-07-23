package handlers

import (
	"net/http"
	"products/models"
	"products/queries"
	vs "products/validationMessages"
	"time"

	"github.com/gin-gonic/gin"
)

var category *models.Category

func GetCategories(c *gin.Context) {
	categories, err := queries.GetCategoiesQuery()
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, categories)
}

func CreateCategory(c *gin.Context) {
	if err := c.ShouldBindJSON(&category); err != nil {
		errMsg := vs.CategoryMessageValidate(err)
		c.JSON(http.StatusBadRequest, gin.H{"message": errMsg})
		return
	}

	category.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	category.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

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
		return
	}

	c.JSON(http.StatusOK, categoryById)
}

func UpdateCategoryById(c *gin.Context) {
	id := c.Param("id")
	if err := c.ShouldBindJSON(&category); err != nil {
		errMsg := vs.CategoryMessageValidate(err)
		c.JSON(http.StatusBadRequest, gin.H{"message": errMsg})
		return
	}

	category.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

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