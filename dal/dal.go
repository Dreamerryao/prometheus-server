package dal

import (
	"context"
	"fmt"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var once sync.Once
var ConnQuery *mongo.Client
var MongoDefaultCtx context.Context

// MongoDBInit 数据库连接
func MongoDBInit() {
	once.Do(func() {
		MongoDefaultCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		ConnQuery, err := mongo.Connect(MongoDefaultCtx, options.Client().ApplyURI("mongodb://localhost:27017"))
		defer func() {
			if err = ConnQuery.Disconnect(MongoDefaultCtx); err != nil {
				panic(err)
			}
		}()
		fmt.Println("数据库连接成功")
	})
}
