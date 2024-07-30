package db

import (
	"context"
	"strconv"
	"sync"

	"github.com/TensoRaws/FinalRip/module/config"
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
	credential := options.Credential{
		Username: config.DBConfig.Username,
		Password: config.DBConfig.Password,
	}
	applyURI := "mongodb://" + config.DBConfig.Host + ":" + strconv.Itoa(config.DBConfig.Port)
	log.Logger.Info("Connecting to MongoDB: " + applyURI)
	clientOpts := options.Client().ApplyURI(applyURI).SetAuth(credential)

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
