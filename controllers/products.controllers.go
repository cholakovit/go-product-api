package controllers

import (
	"net/http"
	"products/models"
	"sync"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	product		models.Product
)

type ProductControllers struct {}

func (pc *ProductControllers) CreateProduct(c *gin.Context) {
	var product *models.Product
	var wg sync.WaitGroup

	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{ "message": err.Error() })
		return
	}

	wg.Add(1)	
	go func() {
		defer wg.Done()

		err := product.CreateProductQuery(product)
		if err != nil {
			c.JSON(http.StatusBadGateway, gin.H{ "message": err.Error() })
			return
		}

	}()
	wg.Wait()

	c.JSON(http.StatusOK, gin.H{"message": "create product success"})
}

func (pc *ProductControllers) GetProducts(c *gin.Context) {
	resultChan := make(chan []primitive.M)
	errChan := make(chan error)
	
	go func ()  {
		products, err := product.GetAllQuery()
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

func (pc *ProductControllers) GetProductById(c *gin.Context) {
	id := c.Param("id")
	resultChan := make(chan primitive.M)
	errChan := make(chan error)

	go func() {
		productById, err := product.GetProductByIdQuery(&id)
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

func (pc *ProductControllers) UpdateProduct(c *gin.Context) {
	id := c.Param("id")
	var product models.Product
	var wg sync.WaitGroup

	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{ "message": err.Error() })
		return
	}

	wg.Add(1)
	go func ()  {
		defer wg.Done()
		err := product.UpdateProductQuery(&id, &product)
		if err != nil {
			c.JSON(http.StatusBadGateway, gin.H{ "message": err.Error() })
			return
		}	
	}()
	wg.Wait()

	c.JSON(http.StatusOK, gin.H{"message": "update product success"})
}

func (pc *ProductControllers) DeleteProduct(c *gin.Context) {
	id := c.Param("id")
	var wg sync.WaitGroup

	wg.Add(1)
	go func ()  {
		defer wg.Done()
		err := product.DeleteProductQuery(&id)
		if err != nil {
			c.JSON(http.StatusBadGateway, gin.H{ "message": err.Error() })
			return
		}
	}()
	wg.Wait()
	
	c.JSON(http.StatusOK, gin.H{"message": "delete product success"})
}



















// package controllers

// import (
// 	"net/http"
// 	"products/models"

// 	"github.com/gin-gonic/gin"
// )

// var (
// 	product		models.Product
// )

// type ProductControllers struct {}

// func (pc *ProductControllers) CreateProduct(c *gin.Context) {
// 	var product *models.Product

// 	if err := c.ShouldBindJSON(&product); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{ "message": err.Error() })
// 		return
// 	}

// 	err := product.CreateProductQuery(product)
// 	if err != nil {
// 		c.JSON(http.StatusBadGateway, gin.H{ "message": err.Error() })
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"message": "create product success"})
// }

// func (pc *ProductControllers) GetProducts(c *gin.Context) {
// 	products, err := product.GetAllQuery()
// 	if err != nil {
// 		c.JSON(http.StatusBadGateway, gin.H{ "message": err.Error() })
// 		return
// 	}
// 	c.JSON(http.StatusOK, products)
// }

// func (pc *ProductControllers) GetProductById(c *gin.Context) {
// 	id := c.Param("id")
	
// 	productById, err := product.GetProductByIdQuery(&id)
// 	if err != nil {
// 		c.JSON(http.StatusBadGateway, gin.H{ "message": err.Error() })
// 		return	
// 	}

// 	c.JSON(http.StatusOK, productById)
// }

// func (pc *ProductControllers) UpdateProduct(c *gin.Context) {
// 	id := c.Param("id")
// 	var product models.Product

// 	if err := c.ShouldBindJSON(&product); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{ "message": err.Error() })
// 		return
// 	}

// 	err := product.UpdateProductQuery(&id, &product)
// 	if err != nil {
// 		c.JSON(http.StatusBadGateway, gin.H{ "message": err.Error() })
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"message": "update product success"})
// }

// func (pc *ProductControllers) DeleteProduct(c *gin.Context) {
// 	id := c.Param("id")

// 	err := product.DeleteProductQuery(&id)
// 	if err != nil {
// 		c.JSON(http.StatusBadGateway, gin.H{ "message": err.Error() })
// 		return
// 	}
	
// 	c.JSON(http.StatusOK, gin.H{"message": "delete product success"})
// }