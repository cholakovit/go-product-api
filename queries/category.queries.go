package queries

import (
	"context"
	"errors"
	"log"
	"products/db"
	"products/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var categoryCollection *mongo.Collection = db.OpenCollection(db.Client, "categories")

func GetCategoiesQuery() ([]primitive.M, error) {
	var categories []primitive.M
	var ctx, cancel = context.WithTimeout(context.Background(), 100 * time.Second)
	defer cancel()

	cursor, err := categoryCollection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	if err = cursor.All(ctx, &categories); err != nil {
		return nil, err
	}

	if len(categories) == 0 {
		return nil, errors.New("category not found")
	}

	return categories, nil
}

func CreateCategoryQuery(category *models.Category) error {
	var ctx, cancel = context.WithTimeout(context.Background(), 100 * time.Second)
	defer cancel()

	_, err := categoryCollection.InsertOne(ctx, category)
	if err != nil {
		log.Fatal(err)
	}
	return err
}

func GetCategoryByIdQuery(id *string) (primitive.M, error) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100 * time.Second)
	defer cancel()

	var category bson.M
	primitiveId, _ := primitive.ObjectIDFromHex(*id)

	query := bson.D{primitive.E{Key: "_id", Value: primitiveId}}

	categoryCollection.FindOne(ctx, query).Decode(&category)

	if len(category) == 0 {
		return nil, errors.New("category not found")
	}

	return category, nil
}

func UpdateCategoryByIdQuery(id *string, category *models.Category) error {
	var ctx, cancel = context.WithTimeout(context.Background(), 100 * time.Second)
	defer cancel()
	
	primitiveId, _ := primitive.ObjectIDFromHex(*id)
	filter := bson.D{primitive.E{Key: "_id", Value: primitiveId}}

	update := bson.D{bson.E{ Key: "$set", Value: bson.D {
		bson.E{ Key: "title", Value: category.Title },
		bson.E{ Key: "desc", Value: category.Desc },
		bson.E{ Key: "tags", Value: category.Tags },
		bson.E{ Key: "category_id", Value: category.Category_id },
		bson.E{ Key: "created_at", Value: category.Created_at },
		bson.E{ Key: "updated_at", Value: category.Updated_at },
		bson.E{ Key: "user_id", Value: category.User_id },
 	} } }

	result, _ := categoryCollection.UpdateOne(ctx, filter,update)
	if result.MatchedCount != 1 {
		return errors.New("no matched category found for update")
	}

	return nil
}

func DeleteCategoryByIdQuery(id *string) error {
	var ctx, cancel = context.WithTimeout(context.Background(), 100 * time.Second)
	defer cancel()

	// ID of the document to delete
	primitiveId, err := primitive.ObjectIDFromHex(*id)
	if err != nil {
		log.Fatal(err)
	}

	filter := bson.M{"_id": primitiveId}

	result, _ := categoryCollection.DeleteOne(ctx, filter)
	if result.DeletedCount != 1 {
		return errors.New("no matched document found for delete")
	}

	return nil
}