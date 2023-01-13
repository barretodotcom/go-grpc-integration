package db

import (
	"context"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Collection

func Connect() error {
	uri := os.Getenv("MONGO_URI")
	fmt.Println(uri)
	conn, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))

	if err != nil {
		return err
	}

	DB = conn.Database("users").Collection("users")

	return nil
}
