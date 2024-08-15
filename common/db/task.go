package db

import (
	"context"

	"github.com/TensoRaws/FinalRip/module/db"
	"go.mongodb.org/mongo-driver/bson"
)

// CheckTaskExist checks if a task exists in the database
func CheckTaskExist(videoKey string) bool {
	coll := db.DB.Collection(TASK_COLLECTION)
	count, _ := coll.CountDocuments(context.TODO(), Task{Key: videoKey})
	return count > 0
}

// CheckTaskStart checks if a task has started
func CheckTaskStart(videoKey string) bool {
	task, err := GetTask(videoKey)
	if err != nil {
		return false
	}
	return task.EncodeParam != ""
}

// InsertTask inserts a new uncompleted task into the database
func InsertTask(videoKey string) error {
	coll := db.DB.Collection(TASK_COLLECTION)
	_, err := coll.InsertOne(context.TODO(), Task{
		Key: videoKey,
	})
	return err
}

// UpdateTask updates an uncompleted task in the database
func UpdateTask(filter Task, update Task) error {
	coll := db.DB.Collection(TASK_COLLECTION)

	up := bson.D{{"$set", update}} //nolint:govet

	_, err := coll.UpdateOne(context.TODO(), filter, up)
	if err != nil {
		return err
	}
	return nil
}

// GetTask gets a completed Task from the database
func GetTask(videoKey string) (Task, error) {
	coll := db.DB.Collection(TASK_COLLECTION)
	var task Task
	err := coll.FindOne(context.TODO(), Task{Key: videoKey}).Decode(&task)

	return task, err
}
