package db

import (
	"context"

	"github.com/TensoRaws/FinalRip/module/db"
	"go.mongodb.org/mongo-driver/bson"
)

// InsertUncompletedTask inserts a new uncompleted task into the database
func InsertUncompletedTask(videoKey string) error {
	coll := db.DB.Collection(COMPLETED_COLLECTION)
	_, err := coll.InsertOne(context.TODO(), CompletedTask{
		Key: videoKey,
	})
	return err
}

// UpdateUncompletedTask updates an uncompleted task in the database
func UpdateUncompletedTask(videoKey string, encodeKey string) error {
	coll := db.DB.Collection(COMPLETED_COLLECTION)

	filter := CompletedTask{
		Key: videoKey,
	}

	up := bson.D{{"$set", CompletedTask{ //nolint: govet
		EncodeKey: encodeKey,
	}}}

	_, err := coll.UpdateOne(context.TODO(), filter, up)
	if err != nil {
		return err
	}
	return nil
}

// GetCompletedEncode gets a completed encode video from the database
func GetCompletedEncode(videoKey string) (string, error) {
	coll := db.DB.Collection(COMPLETED_COLLECTION)
	var task CompletedTask
	err := coll.FindOne(context.TODO(), CompletedTask{Key: videoKey}).Decode(&task)
	return task.EncodeKey, err
}
