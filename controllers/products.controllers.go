package controllers

import (
	"net/http"
	"products/models"

	"github.com/gin-gonic/gin"
)

var (
	product		models.Product
)

type ProductControllers struct {}

func (pc *ProductControllers) CreateProduct(c *gin.Context) {
	var product *models.Product

	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{ "message": err.Error() })
		return
	}

	err := product.CreateProductQuery(product)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{ "message": err.Error() })
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "create product success"})
}

func (pc *ProductControllers) GetProducts(c *gin.Context) {
	products, err := product.GetAllQuery()
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{ "message": err.Error() })
		return
	}
	c.JSON(http.StatusOK, products)
}

func (pc *ProductControllers) GetProductById(c *gin.Context) {
	id := c.Param("id")
	
	productById, err := product.GetProductByIdQuery(&id)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{ "message": err.Error() })
		return	
	}

	c.JSON(http.StatusOK, productById)
}

func (pc *ProductControllers) UpdateProduct(c *gin.Context) {
	id := c.Param("id")
	var product models.Product

	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{ "message": err.Error() })
		return
	}

	err := product.UpdateProductQuery(&id, &product)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{ "message": err.Error() })
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "update product success"})
}

func (pc *ProductControllers) DeleteProduct(c *gin.Context) {
	id := c.Param("id")

	err := product.DeleteProductQuery(&id)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{ "message": err.Error() })
		return
	}
	
	c.JSON(http.StatusOK, gin.H{"message": "delete product success"})
}