package db

import (
	"context"
	"sync"

	"github.com/TensoRaws/FinalRip/module/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	DB   *mongo.Database
	once sync.Once
)

func Init() {
	once.Do(func() {
		initialize()
	})
}

func initialize() {
	// 初始化数据库
	credential := options.Credential{
		Username: "root",
		Password: "123456",
	}
	clientOpts := options.Client().ApplyURI("mongodb://localhost:27017").SetAuth(credential)

	cilent, err := mongo.Connect(context.TODO(), clientOpts)
	if err != nil {
		log.Logger.Error("Failed to connect to MongoDB: " + err.Error())
		panic(err)
	}

	// Send a ping to confirm a successful connection
	var result bson.M
	if err = cilent.Database("finalrip").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Decode(&result); err != nil { //nolint
		log.Logger.Error("Failed to connect to MongoDB: " + err.Error())
		panic(err)
	}

	DB = cilent.Database("finalrip")
}
