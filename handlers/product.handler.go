package handlers

import (
	"net/http"
	"products/models"
	"products/queries"
	vs "products/validationMessages"
	"time"

	"sync"

	"github.com/gin-gonic/gin"
)

var (
	product		*models.Product
	wg 				sync.WaitGroup
)

func GetProducts(c *gin.Context) {
	products, err := queries.GetProductsQuery()
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, products)
}

func CreateProduct(c *gin.Context) {
	if err := c.ShouldBindJSON(&product); err != nil {
		errMsg := vs.ProductMessageValidate(err)
		c.JSON(http.StatusBadRequest, gin.H{"message": errMsg})
		return
	}

	product.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	product.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

	err := queries.CreateProductQuery(product)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{ "message": err.Error() })
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "create product success"})
}

func GetProductById(c *gin.Context) {
	id := c.Param("id")
	productById, err := queries.GetProductByIdQuery(&id)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, productById)
}

func UpdateProductById(c *gin.Context) {
	id := c.Param("id")
	if err := c.ShouldBindJSON(&product); err != nil {
		errMsg := vs.ProductMessageValidate(err)
		c.JSON(http.StatusBadRequest, gin.H{"message": errMsg})
		return
	}

	product.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		
	err := queries.UpdateProductByIdQuery(&id, product)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{ "message": err.Error() })
		return
	}	

	c.JSON(http.StatusOK, gin.H{"message": "update product success"})
}

func DeleteProductById(c *gin.Context) {
	id := c.Param("id")

	err := queries.DeleteProductByIdQuery(&id)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{ "message": err.Error() })
		return
	}
	
	c.JSON(http.StatusOK, gin.H{"message": "delete product success"})
}