package config

import (
	"context"
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Connect() (*mongo.Client, error) {
	connectionString := "mongodb://localhost:27017"

	clientOptions := options.Client().ApplyURI(connectionString)

	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		return nil, fmt.Errorf("could not Connect to MongoDB: %v", err)
	}

	return client, nil

}

func GetDB() *mongo.Database {
	db, err := Connect()

	if err != nil {
		fmt.Printf("Could not get database name")
	}

	return db.Database("task_manager")
}
