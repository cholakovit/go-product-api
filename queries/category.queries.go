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

var categoryCollection *mongo.Collection = db.OpenCollection(db.Client, "categories")
     
func GetCategoriesQuery() ([]primitive.M, error) {
	var categories []primitive.M
	//var err error
 
	// Create a channel to receive the query result and error
	resultChan := make(chan []primitive.M)
	errChan := make(chan error)

	go func() {
		// Perform the query asynchronously
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		cursor, err := categoryCollection.Find(ctx, bson.M{})
		if err != nil {
			errChan <- err
			return
		}
		defer cursor.Close(ctx)

		if err = cursor.All(ctx, &categories); err != nil {
			errChan <- err
			return
		}

		resultChan <- categories
	}()

	// Wait for either the result or an error
	select {
	case categories := <-resultChan:
		if len(categories) == 0 {
			return nil, errors.New("category not found")
		}
		return categories, nil

	case err := <-errChan:
		return nil, err
	}
}

func CreateCategoryQuery(category *models.Category) error {
	//var err error

	// Create a channel to receive the insert result
	resultChan := make(chan interface{})
	errChan := make(chan error)

	go func() {
		// Perform the insert operation asynchronously
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		_, err := categoryCollection.InsertOne(ctx, category)
		if err != nil {
			errChan <- err
			return
		}

		resultChan <- nil
	}()

	// Wait for either the result or an error
	select {
	case <-resultChan:
		return nil

	case err := <-errChan:
		return err
	}
}

func GetCategoryByIdQuery(id *string) (primitive.M, error) {
	var category bson.M
	//var err error

	// Create a channel to receive the query result and error
	resultChan := make(chan bson.M)
	errChan := make(chan error)

	go func() {
		// Perform the query asynchronously
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		primitiveId, err := primitive.ObjectIDFromHex(*id)
		if err != nil {
			errChan <- err
			return
		}

		query := bson.D{primitive.E{Key: "_id", Value: primitiveId}}

		err = categoryCollection.FindOne(ctx, query).Decode(&category)
		if err != nil {
			errChan <- err
			return
		}

		resultChan <- category
	}()

	// Wait for either the result or an error
	select {
	case category := <-resultChan:
		if len(category) == 0 {
			return nil, errors.New("category not found")
		}
		return category, nil

	case err := <-errChan:
		return nil, err
	}
}

func UpdateCategoryByIdQuery(id *string, category *models.Category) error {
	var wg sync.WaitGroup
	var mutex sync.Mutex
	var err error

	primitiveId, _ := primitive.ObjectIDFromHex(*id)
	filter := bson.D{primitive.E{Key: "_id", Value: primitiveId}}

	update := bson.D{
			bson.E{Key: "$set", Value: bson.D{
					bson.E{Key: "title", Value: category.Title},
					bson.E{Key: "desc", Value: category.Desc},
					bson.E{Key: "tags", Value: category.Tags},
					bson.E{Key: "category_id", Value: category.Category_id},
					bson.E{Key: "created_at", Value: category.Created_at},
					bson.E{Key: "updated_at", Value: category.Updated_at},
					bson.E{Key: "user_id", Value: category.User_id},
			}},
	}

	wg.Add(1)
	go func() {
			defer wg.Done()
			var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
			defer cancel()

			result, _ := categoryCollection.UpdateOne(ctx, filter, update)
			if result.MatchedCount != 1 {
					// Synchronize access to the error variable
					mutex.Lock()
					defer mutex.Unlock()
					// Store the error in a shared variable
					err = errors.New("no matched category found for update")
			}
	}()

	wg.Wait() // Wait for goroutine to finish

	// Check if any errors occurred during the goroutine
	mutex.Lock()
	defer mutex.Unlock()
	if err != nil {
			return err
	}

	return nil
}

func DeleteCategoryByIdQuery(id *string) error {
	var wg sync.WaitGroup
	var mutex sync.Mutex
	var err error

	wg.Add(1)
	go func() {
			defer wg.Done()
			var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
			defer cancel()

			// ID of the document to delete
			primitiveId, err := primitive.ObjectIDFromHex(*id)
			if err != nil {
					// Synchronize access to the error variable
					mutex.Lock()
					defer mutex.Unlock()
					// Store the error in a shared variable
					err = errors.New("invalid ID")
					return
			}

			filter := bson.M{"_id": primitiveId}

			result, _ := categoryCollection.DeleteOne(ctx, filter)
			if result.DeletedCount != 1 {
					// Synchronize access to the error variable
					mutex.Lock()
					defer mutex.Unlock()
					// Store the error in a shared variable
					err = errors.New("no matched document found for delete")
			}
	}()

	wg.Wait() // Wait for goroutine to finish

	// Check if any errors occurred during the goroutine
	mutex.Lock()
	defer mutex.Unlock()
	if err != nil {
			return err
	}

	return nil
}