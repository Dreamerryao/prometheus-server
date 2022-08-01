package utils

import (
	"context"
	"fmt"
	"time"

	"github.com/Dreamerryao/prometheus-server/common"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoDBInit 数据库连接
func MongoDBInit() *mongo.Client {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	uri := fmt.Sprintf(
		"mongodb://%s:%s@%s:%d/?authSource=admin&authMechanism=SCRAM-SHA-256",
		common.MongodbUser,
		common.MongodbPwd,
		"127.0.0.1",
		27017,
	)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		fmt.Printf("connect mongodb error: %v\n", err)
	}
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
	fmt.Println("Connected to MongoDB!")
	return client
}

var ConnQuery *mongo.Client = MongoDBInit()

func OpenCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	prometheusDb := client.Database("prometheus")
	var collection *mongo.Collection = prometheusDb.Collection(collectionName)

	return collection
}
