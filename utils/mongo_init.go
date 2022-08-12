package utils

import (
	"context"
	"fmt"
	"time"

	"github.com/Dreamerryao/prometheus-server/common"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var ConnQuery *mongo.Client
var MongoCtx context.Context
var MongoCtxCancel context.CancelFunc

// MongoDBInit 数据库连接
func MongoDBInit() {
	var err error

	MongoCtx, MongoCtxCancel = context.WithTimeout(context.Background(), 10*time.Second)
	uri := fmt.Sprintf(
		"mongodb://%s:%s@%s:%d/?authSource=admin&authMechanism=SCRAM-SHA-256",
		common.MongodbUser,
		common.MongodbPwd,
		// "127.0.0.1",
		"120.55.12.109",
		27017,
	)
	ConnQuery, err = mongo.Connect(MongoCtx, options.Client().ApplyURI(uri))
	if err != nil {
		fmt.Printf("connect mongodb error: %v\n", err)
	}
	fmt.Println("Connected to MongoDB!")
}

func CloseMongoClient() {
	MongoCtxCancel()
	fmt.Print("close mongodb client\n")
	if err := ConnQuery.Disconnect(MongoCtx); err != nil {
		panic(err)
	}
}

func OpenCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	prometheusDb := client.Database("prometheus")
	var collection *mongo.Collection = prometheusDb.Collection(collectionName)

	return collection
}
