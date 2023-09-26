package database

import (
	"context"
	"fmt"
	"log"
	"main/common/config"
	"main/common/tabletypes"

	"strconv"

	"github.com/ethereum/go-ethereum/core/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// func Insert(dba *gorm.DB, logdata *types.Log) {
func Insert(transfer_collection *mongo.Collection, approval_collection *mongo.Collection, approvalforall_collection *mongo.Collection, owner_collection *mongo.Collection, logdata types.Log) {
	topic := logdata.Topics[0].Hex()
	topicTable := config.Topic[topic]
	tokenIDInt := int64(0)
	approved := false
	firstAddr := ""
	secondAddr := ""
	if topicTable == 1 || topicTable == 2 || topicTable == 3 {
		firstAddr = "0x" + logdata.Topics[1].Hex()[26:]
		secondAddr = "0x" + logdata.Topics[2].Hex()[26:]
		if len(logdata.Topics) >= 4 {
			tokenIDInt, _ = strconv.ParseInt(logdata.Topics[3].Hex()[2:], 16, 64)
		} else {
			if logdata.Data[31:][0] == 1 {
				approved = true
			}
		}
	}

	if topicTable == 1 { //Approval
		var res tabletypes.Approval
		err := approval_collection.FindOne(context.TODO(), bson.D{{Key: "txhash", Value: logdata.TxHash.Hex()}, {Key: "logindex", Value: int(logdata.Index)}}).Decode(&res)
		if err != nil {
			_, err = approval_collection.InsertOne(context.TODO(), bson.D{
				{Key: "blocknumber", Value: logdata.BlockNumber},
				{Key: "timestamp", Value: config.BlockWithTimestamp[fmt.Sprint(logdata.BlockNumber)]},
				{Key: "address", Value: logdata.Address.Hex()},
				{Key: "func", Value: topic},
				{Key: "owner", Value: firstAddr},
				{Key: "approved", Value: secondAddr},
				{Key: "tokenid", Value: tokenIDInt},
				{Key: "txhash", Value: logdata.TxHash.Hex()},
				{Key: "logindex", Value: logdata.Index},
			})
			// dba.Create(&tabletypes.Approval{
			// 	Blocknumber: logdata.BlockNumber,
			// 	Timestamp:   config.BlockWithTimestamp[fmt.Sprint(logdata.BlockNumber)],
			// 	Address:     logdata.Address.Hex(),
			// 	Func:        topic,
			// 	Owner:       firstAddr,
			// 	Approved:    secondAddr,
			// 	Tokenid:     tokenIDInt,
			// 	Txhash:      logdata.TxHash.Hex(),
			// 	Logindex:    logdata.Index,
			// })
		}

	} else if topicTable == 2 { //Transfer
		var res tabletypes.Transfer
		err := transfer_collection.FindOne(context.TODO(), bson.D{{Key: "txhash", Value: logdata.TxHash.Hex()}, {Key: "logindex", Value: int(logdata.Index)}}).Decode(&res)
		if err != nil {
			_, err = transfer_collection.InsertOne(context.TODO(), bson.D{
				{Key: "blocknumber", Value: logdata.BlockNumber},
				{Key: "timestamp", Value: config.BlockWithTimestamp[fmt.Sprint(logdata.BlockNumber)]},
				{Key: "address", Value: logdata.Address.Hex()},
				{Key: "func", Value: topic},
				{Key: "from", Value: firstAddr},
				{Key: "to", Value: secondAddr},
				{Key: "tokenid", Value: tokenIDInt},
				{Key: "txhash", Value: logdata.TxHash.Hex()},
				{Key: "logindex", Value: logdata.Index},
			})

			ModifyOwner(owner_collection, logdata.Address.Hex(), int(tokenIDInt), secondAddr)
		}
	} else if topicTable == 3 { //ApprovalForAll
		var res tabletypes.ApprovalForAll
		err := approvalforall_collection.FindOne(context.TODO(), bson.D{{Key: "txhash", Value: logdata.TxHash.Hex()}, {Key: "logindex", Value: int(logdata.Index)}}).Decode(&res)
		if err != nil {
			_, err = approvalforall_collection.InsertOne(context.TODO(), bson.D{
				{Key: "blocknumber", Value: logdata.BlockNumber},
				{Key: "timestamp", Value: config.BlockWithTimestamp[fmt.Sprint(logdata.BlockNumber)]},
				{Key: "address", Value: logdata.Address.Hex()},
				{Key: "func", Value: topic},
				{Key: "owner", Value: firstAddr},
				{Key: "operator", Value: secondAddr},
				{Key: "approved", Value: approved},
				{Key: "txhash", Value: logdata.TxHash.Hex()},
				{Key: "logindex", Value: logdata.Index},
			})
		}
	}
}

func ModifyOwner(owner_collection *mongo.Collection, address string, id int, owner string) {
	var res tabletypes.Owner
	err := owner_collection.FindOne(context.TODO(), bson.D{{Key: "address", Value: address}, {Key: "tokenid", Value: id}}).Decode(&res)
	if err != nil {
		_, err = owner_collection.InsertOne(context.TODO(), bson.D{
			{Key: "address", Value: address},
			{Key: "owner", Value: owner},
			{Key: "tokenid", Value: id},
		})
	} else {
		filter := bson.D{{Key: "address", Value: address}, {Key: "tokenid", Value: id}}
		update := bson.D{{Key: "$set", Value: bson.D{{Key: "owner", Value: owner}}}}
		_, err := owner_collection.UpdateOne(context.TODO(), filter, update)
		if err != nil {
			log.Fatal(err)
		}
		//fmt.Println(updateres.ModifiedCount)
	}
}
