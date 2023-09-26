package dbconn

import (
	"context"
	"fmt"
	"main/common/config"
	"time"

	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetDB() (*mongo.Database, error) {
	log.SetFormatter(&log.JSONFormatter{})
	// client, err := mongo.NewClient((options.Client().ApplyURI("mongodb://localhost:27017")))
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	dsn := fmt.Sprintf("mongodb://%s:%s@%s:%d",
		config.MongodbCon.Username, config.MongodbCon.Password, config.MongodbCon.Addr, config.MongodbCon.Port)
	//fmt.Println(dsn)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(
		dsn,
		//"mongodb://"+config.Username+":"+config.Password+"@localhost:27017",
		//"mongodb+srv://standard:"+password+"@cluster0.pdpui.mongodb.net/"+dbname+"?retryWrites=true&w=majority",
	))

	if err != nil {
		log.Error(err.Error())
		return nil, err
	}

	//Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}

	database := client.Database(config.MongodbCon.Db)
	return database, nil
}

func GetCollection() (transfer_collection *mongo.Collection, approval_collection *mongo.Collection, approvalforall_collection *mongo.Collection, owner_collection *mongo.Collection) {
	db, _ := GetDB()
	transfer_collection = db.Collection(config.Transfer_collections)
	approval_collection = db.Collection(config.Approval_collections)
	approvalforall_collection = db.Collection(config.Approvalforall_collections)
	owner_collection = db.Collection(config.Owner_collections)
	return
}

// func Buildconnect() *gorm.DB {
// 	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local&timeout=%s",
// 		config.MysqlCon.Username, config.MysqlCon.Password, config.MysqlCon.Addr, config.MysqlCon.Port, config.MysqlCon.Db, "10s")
// 	//mysql connection
// 	dba, err := gorm.Open("mysql", dsn)
// 	if err != nil {
// 		fmt.Printf("Connect error:%s\n", err)
// 	}
// 	return dba
// }
