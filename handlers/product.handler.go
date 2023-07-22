package handlers

import (
	"net/http"
	"products/models"
	"products/queries"
	vs "products/validationMessages"

	"sync"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	product		*models.Product
	wg 				sync.WaitGroup
)

func GetProducts(c *gin.Context) {
	resultChan := make(chan []primitive.M)
	errChan := make(chan error)
	
	go func ()  {
		products, err := queries.GetProductsQuery()
		if err != nil {
			errChan <- err
		} else {
			resultChan <- products
		}
		close(resultChan) // Close the result channel after sending the result or error
		close(errChan) // Close the error channel after sending the result or error
	}()
	
	select {
		case products := <- resultChan:
			c.JSON(http.StatusOK, products)
		case err := <- errChan:
			c.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
	}
}

func CreateProduct(c *gin.Context) {
	if err := c.ShouldBindJSON(&product); err != nil {
		errMsg := vs.ProductMessageValidate(err)
		c.JSON(http.StatusBadRequest, gin.H{"message": errMsg})
		return
	}

	wg.Add(1)	
	go func() {
		defer wg.Done()

		err := queries.CreateProductQuery(product)
		if err != nil {
			c.JSON(http.StatusBadGateway, gin.H{ "message": err.Error() })
			return
		}

	}()
	wg.Wait()

	c.JSON(http.StatusOK, gin.H{"message": "create product success"})
}

func GetProductById(c *gin.Context) {
	id := c.Param("id")
	resultChan := make(chan primitive.M)
	errChan := make(chan error)

	go func() {
		productById, err := queries.GetProductByIdQuery(&id)
		if err != nil {
			errChan <- err
		}else {
			resultChan <- productById
		}
		close(resultChan)
		close(errChan)
	}()

	select {
		case productById := <- resultChan:
			c.JSON(http.StatusOK, productById)
		case err := <- errChan:
			c.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
	}	
}

func UpdateProductById(c *gin.Context) {
	id := c.Param("id")
	if err := c.ShouldBindJSON(&product); err != nil {
		errMsg := vs.ProductMessageValidate(err)
		c.JSON(http.StatusBadRequest, gin.H{"message": errMsg})
		return
	}

	wg.Add(1)
	go func ()  {
		defer wg.Done()
		err := queries.UpdateProductByIdQuery(&id, product)
		if err != nil {
			c.JSON(http.StatusBadGateway, gin.H{ "message": err.Error() })
			return
		}	
	}()
	wg.Wait()

	c.JSON(http.StatusOK, gin.H{"message": "update product success"})
}

func DeleteProductById(c *gin.Context) {
	id := c.Param("id")
	var wg sync.WaitGroup

	wg.Add(1)
	go func ()  {
		defer wg.Done()
		err := queries.DeleteProductByIdQuery(&id)
		if err != nil {
			c.JSON(http.StatusBadGateway, gin.H{ "message": err.Error() })
			return
		}
	}()
	wg.Wait()
	
	c.JSON(http.StatusOK, gin.H{"message": "delete product success"})
}