package db

import (
	"context"
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Db struct {
}

func (d *Db) Connect() (*mongo.Client, context.Context) {
	ctx := context.TODO()

	mongoURI := os.Getenv("MONGO_URI")
	if mongoURI == "" {
		log.Fatal("MONGO_URI environment variable is not set")
	}

	mongoconn := options.Client().ApplyURI(mongoURI)
	mongoclient, err := mongo.Connect(ctx, mongoconn)
	if err != nil {
		log.Fatal(err)
	}

	err = mongoclient.Ping(ctx, readpref.Primary())

	fmt.Println("mongo connection established")

	return mongoclient, ctx
}