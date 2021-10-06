package mongodb

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func InitLog() *mongo.Client {
	var err error
	var DBLog *mongo.Client
	DBLog, err = mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017/ez_musix"))
	if err != nil {
		panic(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 1*time.Second)
	err = DBLog.Connect(ctx)
	if err != nil {
		panic(err)
	}

	DBLog.ListDatabaseNames(ctx, bson.M{})
	return DBLog
}
