package models

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Task struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Description string             `bson:"description"`
	DueDate     string             `bson:"due_date"`
	Completed   bool               `bson:"completed"`
}

func TaskCollection(db *mongo.Database) *mongo.Collection {
	return db.Collection("tasks")
}

func GetAllTasks(db *mongo.Database) ([]Task, error) {

	findOptions := options.Find()

	cursor, err := TaskCollection(db).Find(context.TODO(), bson.D{}, findOptions)

	if err != nil {
		fmt.Printf("Could perform find operation")
		return nil, err
	}

	defer cursor.Close(context.TODO())

	var task []Task

	for cursor.Next(context.TODO()) {
		var result Task
		if err := cursor.Decode(&result); err != nil {
			return nil, err
		}

		task = append(task, result)
	}

	return task, nil
}

func (t *Task) CreateTask(db *mongo.Database) *Task {

	collection := db.Collection("tasks")

	_, err := collection.InsertOne(context.TODO(), t)

	if err != nil {
		fmt.Printf("Could not create Task")
		return nil
	}

	return t
}

func GetTaskById(db *mongo.Database, id string) *Task {
	collection := db.Collection("tasks")

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		fmt.Printf("Invalid ObjectID: %v", err)
		return nil
	}

	filter := bson.M{"_id": objectID}
	findOptions := options.FindOne()

	var result *Task
	err = collection.FindOne(context.TODO(), filter, findOptions).Decode(&result)

	if err != nil {
		fmt.Printf("Could not find that Task by ID: %v", err)
	}
	return result
}
