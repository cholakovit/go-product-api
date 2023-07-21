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

var collection *mongo.Collection = db.OpenCollection(db.Client, "products")

func CreateProductQuery(product *models.Product) error {
	var ctx, cancel = context.WithTimeout(context.Background(), 100 * time.Second)
	defer cancel()
	
	_, err := collection.InsertOne(ctx, product)
	if err != nil {
		log.Fatal(err)
	}
	return err
}

func GetAllQuery() ([]primitive.M, error) {
	var products []primitive.M
	var ctx, cancel = context.WithTimeout(context.Background(), 100 * time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	if err = cursor.All(ctx, &products); err != nil {
		return nil, err
	}

	if err != nil {
		return nil, err
	}

	if len(products) == 0 {
		return nil, errors.New("documents not found")
	}

	return products, nil
}

func DeleteProductQuery(id *string) error {
	var ctx, cancel = context.WithTimeout(context.Background(), 100 * time.Second)
	defer cancel()

	// ID of the document to delete
	primitiveId, err := primitive.ObjectIDFromHex(*id)
	if err != nil {
		log.Fatal(err)
	}

	filter := bson.M{"_id": primitiveId}

	result, _ := collection.DeleteOne(ctx, filter)
	if result.DeletedCount != 1 {
		return errors.New("no matched document found for delete")
	}

	return nil
}

func GetProductByIdQuery(id *string) (primitive.M, error) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100 * time.Second)
	defer cancel()
	var product bson.M
	
	primitiveId, _ := primitive.ObjectIDFromHex(*id)

	query := bson.D{primitive.E{Key: "_id", Value: primitiveId}}

	collection.FindOne(ctx, query).Decode(&product)

	if len(product) == 0 {
		return nil, errors.New("product not found")
	}

	return product, nil
}

func UpdateProductQuery(id *string, product *models.Product) error {
	var ctx, cancel = context.WithTimeout(context.Background(), 100 * time.Second)
	defer cancel()
	
	primitiveId, _ := primitive.ObjectIDFromHex(*id)
	filter := bson.D{primitive.E{Key: "_id", Value: primitiveId}}

	update := bson.D{bson.E{ Key: "$set", Value: bson.D {
		bson.E{ Key: "title", Value: product.Title },
		bson.E{ Key: "desc", Value: product.Desc },
		bson.E{ Key: "img", Value: product.Img },
		bson.E{ Key: "short_desc", Value: product.Short_Desc },
		bson.E{ Key: "manufacturer", Value: product.Manufacturer },
		bson.E{ Key: "price", Value: product.Price },
		bson.E{ Key: "stock", Value: product.Stock },
		bson.E{ Key: "discount", Value: product.Discount },
		bson.E{ Key: "active", Value: product.Active },
		//bson.E{ Key: "thumbs", Value: product.Thumbs},
		bson.E{ Key: "thumbs", Value: bson.D{
			bson.E{Key: "thumb1", Value: product.Thumbs.Thumb1},
			bson.E{Key: "thumb2", Value: product.Thumbs.Thumb2},
			bson.E{Key: "thumb3", Value: product.Thumbs.Thumb3},
			bson.E{Key: "thumb4", Value: product.Thumbs.Thumb4},
			bson.E{Key: "thumb5", Value: product.Thumbs.Thumb5},
		} },

	} }}

	result, _ := collection.UpdateOne(ctx, filter,update)
	if result.MatchedCount != 1 {
		return errors.New("no matched prroduct found for update")
	}

	return nil
}