package models

import (
	"context"
	"errors"
	"log"

	"products/db"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Thumbs struct {
	Thumb1 string `json:"thumb1" bson:"thumb1"`
	Thumb2 string `json:"thumb2" bson:"thumb2"`
	Thumb3 string `json:"thumb3" bson:"thumb3"`
	Thumb4 string `json:"thumb4" bson:"thumb4"`
	Thumb5 string `json:"thumb5" bson:"thumb5"`
}

type Product struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	Title        string             `json:"title" binding:"required,min=2" 						bson:"title"`
	Img          string             `json:"img"																				bson:"img"`
	Desc         string             `json:"desc" binding:"required,min=20"						bson:"desc"`
	Short_Desc   string             `json:"short_desc"																bson:"short_desc"`
	Price        float64            `json:"price" binding:"required"									bson:"price"`
	Discount     float64            `json:"discount"																	bson:"discount"`
	Stock        int                `json:"stock"																			bson:"stock"`
	Active       bool               `json:"active"																		bson:"active"`
	Manufacturer string             `json:"manufacturer"															bson:"manufacturer"`
	Thumbs       Thumbs             `json:"thumbs"																		bson:"thumbs"`
}

var (
	connection *db.Db
)

func (p *Product) withCollection(fn func(*mongo.Collection) error) error {
	client, ctx := connection.Connect()
	defer client.Disconnect(ctx)

	productCollection := client.Database("demo").Collection("products")

	return fn(productCollection)
}

func (p *Product) CreateProductQuery(product *Product) error {
	ctx := context.TODO() // Create a new context for the operation
	return p.withCollection(func(collection *mongo.Collection) error {
		_, err := collection.InsertOne(ctx, product)
		if err != nil {
			log.Fatal(err)
		}
		return err
	})
}

func (p *Product) DeleteProductQuery(id *string) error {
	ctx := context.TODO() // Create a new context for the operation
	return p.withCollection(func(collection *mongo.Collection) error {
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
	})
}

func (p *Product) GetAllQuery() ([]primitive.M, error) {
	var products []primitive.M
	ctx := context.TODO() // Create a new context for the operation

	err := p.withCollection(func(collection *mongo.Collection) error {
		cursor, err := collection.Find(ctx, bson.M{})
		if err != nil {
			return err
		}
		defer cursor.Close(ctx)

		if err = cursor.All(ctx, &products); err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	if len(products) == 0 {
		return nil, errors.New("documents not found")
	}

	return products, nil
}

func (p *Product) GetProductByIdQuery(id *string) (primitive.M, error) {
	ctx := context.TODO() // Create a new context for the operation
	var product bson.M

	err := p.withCollection(func(collection *mongo.Collection) error {
		primitiveId, _ := primitive.ObjectIDFromHex(*id)

		query := bson.D{primitive.E{Key: "_id", Value: primitiveId}}

		collection.FindOne(ctx, query).Decode(&product)

		return nil
	})
	if err != nil {
		return nil, err
	}

	if len(product) == 0 {
		return nil, errors.New("document not found")
	}

	return product, nil

}

func (p *Product) UpdateProductQuery(id *string, product *Product) error {
	ctx := context.TODO()

	err := p.withCollection(func(c *mongo.Collection) error {
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

		result, _ := c.UpdateOne(ctx, filter,update)
		if result.MatchedCount != 1 {
			return errors.New("no matched document found for update")
		}

		return nil
	})
	if err != nil {
		return err
	}

	return nil
}