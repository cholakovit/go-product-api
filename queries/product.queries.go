package queries

import (
	"context"
	"errors"
	"products/db"
	"products/models"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	collection *mongo.Collection = db.OpenCollection(db.Client, "products")
	wg 				sync.WaitGroup
)

func GetProductsQuery() ([]primitive.M, error) {
	//var products []primitive.M
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	// Create a channel to receive the fetched products from goroutines
	productChan := make(chan []primitive.M)

	go func() {
		var fetchedProducts []primitive.M
		for cursor.Next(ctx) {
			var product primitive.M
			err := cursor.Decode(&product)
			if err != nil {
				break
			}
			fetchedProducts = append(fetchedProducts, product)
		}
		productChan <- fetchedProducts
	}()

	// Wait for the products to be fetched from the goroutine
	fetchedProducts := <-productChan

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	if len(fetchedProducts) == 0 {
		return nil, errors.New("products not found")
	}

	return fetchedProducts, nil
}

func CreateProductQuery(product *models.Product) error {
	var wg sync.WaitGroup // To wait for goroutine to complete
	var resultErr error   // To store the error from the database operation

	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	wg.Add(1)
	go func() {
			defer wg.Done()
			_, err := collection.InsertOne(ctx, product)
			if err != nil {
					resultErr = err
			}
	}()

	wg.Wait() // Wait for the goroutine to finish
	return resultErr
}

func GetProductByIdQuery(id *string) (primitive.M, error) {
	var product bson.M

	// Create a channel to receive the result
	resultChan := make(chan primitive.M)
	errChan := make(chan error)

	go func() {
			// Perform the database operation asynchronously
			ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
			defer cancel()

			primitiveId, err := primitive.ObjectIDFromHex(*id)
			if err != nil {
					errChan <- err
					return
			}

			query := bson.D{primitive.E{Key: "_id", Value: primitiveId}}

			err = collection.FindOne(ctx, query).Decode(&product)
			if err != nil {
					errChan <- err
					return
			}

			resultChan <- product
	}()

	// Wait for either the result or an error
	select {
		case product := <-resultChan:
				if len(product) == 0 {
						return nil, errors.New("product not found")
				}
				return product, nil

		case err := <-errChan:
				return nil, err
	}
}

func UpdateProductByIdQuery(id *string, product *models.Product) error {
	//var err error

	// Create a channel to receive the update result
	resultChan := make(chan int64)
	errChan := make(chan error)

	go func() {
			// Perform the database update asynchronously
			ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
			defer cancel()

			primitiveId, err := primitive.ObjectIDFromHex(*id)
			if err != nil {
					errChan <- err
					return
			}

			filter := bson.D{primitive.E{Key: "_id", Value: primitiveId}}

			update := bson.D{bson.E{Key: "$set", Value: bson.D{
					bson.E{Key: "title", Value: product.Title},
					bson.E{Key: "desc", Value: product.Desc},
					bson.E{Key: "img", Value: product.Img},
					bson.E{Key: "short_desc", Value: product.Short_Desc},
					bson.E{Key: "manufacturer", Value: product.Manufacturer},
					bson.E{Key: "price", Value: product.Price},
					bson.E{Key: "stock", Value: product.Stock},
					bson.E{Key: "discount", Value: product.Discount},
					bson.E{Key: "active", Value: product.Active},
					bson.E{Key: "thumbs", Value: bson.D{
							bson.E{Key: "thumb1", Value: product.Thumbs.Thumb1},
							bson.E{Key: "thumb2", Value: product.Thumbs.Thumb2},
							bson.E{Key: "thumb3", Value: product.Thumbs.Thumb3},
							bson.E{Key: "thumb4", Value: product.Thumbs.Thumb4},
							bson.E{Key: "thumb5", Value: product.Thumbs.Thumb5},
					}},
			}}}

			result, err := collection.UpdateOne(ctx, filter, update)
			if err != nil {
					errChan <- err
					return
			}

			resultChan <- result.MatchedCount
	}()

	// Wait for either the result or an error
	select {
	case matchedCount := <-resultChan:
			if matchedCount != 1 {
					return errors.New("no matched product found for update")
			}
			return nil

	case err := <-errChan:
			return err
	}
}

func DeleteProductByIdQuery(id *string) error {
	//var err error

	// Create a channel to receive the delete result
	resultChan := make(chan int64)
	errChan := make(chan error)

	go func() {
			// Perform the delete operation asynchronously
			ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
			defer cancel()

			// ID of the document to delete
			primitiveId, err := primitive.ObjectIDFromHex(*id)
			if err != nil {
					errChan <- err
					return
			}

			filter := bson.M{"_id": primitiveId}

			result, err := collection.DeleteOne(ctx, filter)
			if err != nil {
					errChan <- err
					return
			}

			resultChan <- result.DeletedCount
	}()

	// Wait for either the result or an error
	select {
	case deletedCount := <-resultChan:
			if deletedCount != 1 {
					return errors.New("no matched document found for delete")
			}
			return nil

	case err := <-errChan:
			return err
	}
}