package queries

import (
	"context"
	"errors"
	"sync"

	"products/db"
	"products/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection = db.OpenCollection(db.Client, "users")

func GetUsersQuery() ([]primitive.M, error) {
	var users []primitive.M
	var wg sync.WaitGroup
	var mutex sync.Mutex
	var err error

	wg.Add(1)
	go func() {
		defer wg.Done()
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		cursor, err := userCollection.Find(ctx, bson.M{})
		if err != nil {
			// Synchronize access to the error variable
			mutex.Lock()
			defer mutex.Unlock()
			// Store the error in a shared variable
			err = errors.New("failed to find documents")
			return
		}
		defer cursor.Close(ctx)

		if err = cursor.All(ctx, &users); err != nil {
			// Synchronize access to the error variable
			mutex.Lock()
			defer mutex.Unlock()
			// Store the error in a shared variable
			err = errors.New("failed to decode documents")
			return
		}
	}()

	wg.Wait() // Wait for goroutine to finish

	// Check if any errors occurred during the goroutine
	mutex.Lock()
	defer mutex.Unlock()
	if err != nil {
		return nil, err
	}

	if len(users) == 0 {
		return nil, errors.New("documents not found")
	}

	return users, nil
}

func CreateUserQuery(user *models.User) error {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	_, err := userCollection.InsertOne(ctx, user)
	if err != nil {
		return err
	}

	return nil
}

func GetUserByIdQuery(id *string) (*models.User, error) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	var user *models.User

	primitiveID, _ := primitive.ObjectIDFromHex(*id)

	query := bson.D{primitive.E{Key: "_id", Value: primitiveID}}

	err := userCollection.FindOne(ctx, query).Decode(&user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func UpdateUserByIdQuery(id *string, user *models.User) error {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	primitveId, _ := primitive.ObjectIDFromHex(*id)
	filter := bson.D{primitive.E{Key: "_id", Value: primitveId}}

	update := bson.D{bson.E{Key: "$set", Value: bson.D{
			bson.E{Key: "fullname", Value: user.FullName},
			bson.E{Key: "pass", Value: user.Pass},
			bson.E{Key: "email", Value: user.Email},
			bson.E{Key: "phone", Value: user.Phone},
			bson.E{Key: "role", Value: user.Role},
			bson.E{Key: "token", Value: user.Token},
			bson.E{Key: "rtoken", Value: user.Rtoken},
			bson.E{Key: "created_at", Value: user.Created_at},
			bson.E{Key: "updated_at", Value: user.Updated_at},
	}}}

	result, _ := userCollection.UpdateOne(ctx, filter, update)
	if result.MatchedCount != 1 {
			return errors.New("No matched users found for update")
	}

	return nil
}

func DeleteUserByIdQuery(id *string) error {
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

			result, _ := userCollection.DeleteOne(ctx, filter)
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

func FindUserByEmailQuery(email *string) (int64, error) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	count, err := userCollection.CountDocuments(ctx, bson.M{"email": email})
	if err != nil {
			return 0, errors.New("error occurred while checking for the user")
	}

	return count, nil
}

func FindOneQuery(user *models.Auth) (*models.User, error) {
	done := make(chan struct{})
	var result *models.User
	var err error

	go func() {
			defer close(done)
			var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
			defer cancel()

			err = userCollection.FindOne(ctx, bson.M{"email": user.Email}).Decode(&result)
			if err != nil {
					err = errors.New("error occurring while checking for the user")
			}
	}()

	// Wait for goroutine to finish
	<-done

	return result, err
}