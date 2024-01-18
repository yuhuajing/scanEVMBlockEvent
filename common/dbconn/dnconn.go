package dbconn

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"main/common/config"
	"time"
)

func init() {
	config.Mongoclient = GetMongoClient(config.Mongodburl)
}

func GetMongoClient(mongourl string) *mongo.Client {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(
		mongourl,
	))

	if err != nil {
		log.Fatalf("err in conn MonggoDB: %v", err)

	}
	return client
}

func GetCollection(dbcollectionname string) *mongo.Collection {
	err := config.Mongoclient.Ping(context.TODO(), nil)
	if err != nil {
		config.Mongoclient = GetMongoClient(config.Mongodburl)
		log.Fatalf("error in connecting mongodb: %v", err)
	}
	database := config.Mongoclient.Database(config.Dbname)
	txdata_collection := database.Collection(dbcollectionname)
	return txdata_collection
}
