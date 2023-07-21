package db

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
//	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var Client *mongo.Client = DBinstance()

func DBinstance() *mongo.Client {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	MongoDb := os.Getenv("MONGO_URI")

	client, err := mongo.NewClient(options.Client().ApplyURI(MongoDb))
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	return client
}



func OpenCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	var collection *mongo.Collection = client.Database("demo").Collection(collectionName)
	return collection
}







// func (d *Db) Connect() (*mongo.Client, context.Context) {
// 	mongoURI := os.Getenv("MONGO_URI")
// 	if mongoURI == "" {
// 		log.Fatal("MONGO_URI environment variable is not set")
// 	}

// 	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// 	defer cancel()

// 	mongoconn := options.Client().ApplyURI(mongoURI)
// 	mongoclient, err := mongo.Connect(ctx, mongoconn)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	err = mongoclient.Ping(ctx, readpref.Primary())

// 	fmt.Println("mongo connection established")

// 	return mongoclient, ctx
// }