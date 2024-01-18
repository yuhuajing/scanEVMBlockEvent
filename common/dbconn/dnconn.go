package dbconn

import (
	"context"
	"fmt"
	//"github.com/ethereum/go-ethereum/log"
	"log"
	"main/common/config"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetDB() (*mongo.Database, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	mongourl := fmt.Sprintf("mongodb://%s:%s@%s:%d",
		config.MongodbCon.Username, config.MongodbCon.Password, config.MongodbCon.Addr, config.MongodbCon.Port)

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(
		mongourl,
	))

	if err != nil {
		return nil, err
	}

	//Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return nil, err
	}

	database := client.Database(config.MongodbCon.Db)
	return database, nil
}

func GetCollection() (transfer_collection *mongo.Collection, approval_collection *mongo.Collection, approvalforall_collection *mongo.Collection, owner_collection *mongo.Collection) {
	db, err := GetDB()
	if err != nil {
		log.Fatalf("Err in connecting MongoDB", err)
	}
	transfer_collection = db.Collection(config.Transfer_collections)
	approval_collection = db.Collection(config.Approval_collections)
	approvalforall_collection = db.Collection(config.Approvalforall_collections)
	owner_collection = db.Collection(config.Owner_collections)
	return
}
