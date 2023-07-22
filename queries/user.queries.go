package queries

import (
	"context"
	"errors"
	"fmt"
	"log"

	"products/db"
	"products/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection = db.OpenCollection(db.Client, "users")

func CreateUserQuery(user *models.User) error {
	var ctx, cancel = context.WithTimeout(context.Background(), 100 * time.Second)
	defer cancel()

	_, err := userCollection.InsertOne(ctx, user)
	if err != nil {
		log.Fatal(err)
	}

	return err
}

func UpdateUserQuery(id *string, user *models.User) error {
	var ctx, cancel = context.WithTimeout(context.Background(), 100 * time.Second)
	defer cancel()

	primitveId, _ := primitive.ObjectIDFromHex(*id)
	filter := bson.D{primitive.E{Key: "_id", Value: primitveId}}

	update := bson.D{bson.E{ Key: "$set", Value: bson.D {
		bson.E{ Key: "fullname", Value: user.FullName },
		bson.E{ Key: "pass", Value: user.Pass },
		bson.E{ Key: "email", Value: user.Email },
		bson.E{ Key: "phone", Value: user.Phone },
		bson.E{ Key: "role", Value: user.Role },
		bson.E{ Key: "token", Value: user.Token },
		bson.E{ Key: "rtoken", Value: user.Rtoken },
		bson.E{ Key: "created_at", Value: user.Created_at },
		bson.E{ Key: "updated_at", Value: user.Updated_at },
	} }}

	fmt.Println("FullName", *user.FullName)

	result, _ := userCollection.UpdateOne(ctx, filter, update)
	if result.MatchedCount != 1 {
		return errors.New("No matched users found for update")
	}

	return nil
}

func DeleteUserQuery(id *string) error {
	var ctx, cancel = context.WithTimeout(context.Background(), 100 * time.Second)
	defer cancel()

	// ID of the document to delete
	primitiveId, err := primitive.ObjectIDFromHex(*id)
	if err != nil {
		log.Fatal(err)
	}

	filter := bson.M{"_id": primitiveId}

	result, _ := userCollection.DeleteOne(ctx, filter)
	if result.DeletedCount != 1 {
		return errors.New("no matched document found for delete")
	}

	return nil
}

func GetUserByIdQuery(id *string) (primitive.M, error) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100 * time.Second)
	defer cancel()
	var user bson.M
	
	primitiveId, _ := primitive.ObjectIDFromHex(*id)

	query := bson.D{primitive.E{Key: "_id", Value: primitiveId}}

	userCollection.FindOne(ctx, query).Decode(&user)

	if len(user) == 0 {
		return nil, errors.New("product not found")
	}

	return user, nil
}

func GetAllUsersQuery() ([]primitive.M, error) {
	var users []primitive.M
	var ctx, cancel = context.WithTimeout(context.Background(), 100 * time.Second)
	defer cancel()

	cursor, err := userCollection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	if err = cursor.All(ctx, &users); err != nil {
		return nil, err
	}

	if err != nil {
		return nil, err
	}

	if len(users) == 0 {
		return nil, errors.New("documents not found")
	}

	return users, nil
}