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

	collection := TaskCollection(db)

	_, err := collection.InsertOne(context.TODO(), t)

	if err != nil {
		fmt.Printf("Could not create Task")
		return nil
	}

	return t
}

func GetTaskById(db *mongo.Database, id string) *Task {
	collection := TaskCollection(db)

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

func DeleteTaskById(db *mongo.Database, id string) *Task {

	collection := TaskCollection(db)

	objectId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		fmt.Printf("Could not convert int string to primitive.Object: %v", err)
	}

	filter := bson.M{"_id": objectId}
	findOptions := options.FindOneAndDelete()

	var resultTask *Task
	err = collection.FindOneAndDelete(context.TODO(), filter, findOptions).Decode(&resultTask)

	if err != nil {
		fmt.Printf("Could not delete Task from database: %v", err)
	}

	return resultTask

}

func (t *Task) EditTaskCompletionById(db *mongo.Database, id string, updatedCompletion bool) *Task {

	collection := TaskCollection(db)
	objectId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		fmt.Printf("Could not Convert id type string to primitive.Object: %v", err)
	}

	filter := bson.M{"_id": objectId}
	update := bson.M{"$set": bson.M{"Completed": updatedCompletion}}
	updateOptions := options.FindOneAndUpdate().SetUpsert(false)

	var resultTask *Task
	err = collection.FindOneAndUpdate(context.TODO(), filter, update, updateOptions).Decode(&resultTask)

	if err != nil {
		fmt.Printf("Could not Update Task by Id provided: %v", err)
	}

	return resultTask

}

func (t *Task) EditDescriptionById(db *mongo.Database, id string, description string) *Task {

	collection := TaskCollection(db)
	objectId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		fmt.Printf("Could not Convert id type string to primitive.Object: %v", err)
	}

	filter := bson.M{"_id": objectId}
	update := bson.M{"$set": bson.M{"Description": description}}
	updateOptions := options.FindOneAndUpdate().SetUpsert(false)

	var resultTask *Task
	err = collection.FindOneAndUpdate(context.TODO(), filter, update, updateOptions).Decode(&resultTask)

	if err != nil {
		fmt.Printf("Could not Update Task by Id provided: %v", err)
	}

	return resultTask
}
