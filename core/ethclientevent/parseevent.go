package ethclientevent

import (
	"main/core/database"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/jinzhu/gorm"
	"go.mongodb.org/mongo-driver/mongo"
)

func ParseEventLogsMysql(dba *gorm.DB, logs []types.Log) {
	for _, v := range logs {
		ParseEventLogMysql(dba, v)
	}
}

func ParseEventLogMysql(dba *gorm.DB, log types.Log) {
	database.InsertMysql(dba, log)
}

func ParseEventLogsMongo(transfer_collection *mongo.Collection, approval_collection *mongo.Collection, approvalforall_collection *mongo.Collection, owner_collection *mongo.Collection, logs []types.Log) {
	for _, v := range logs {
		ParseEventLogMongo(transfer_collection, approval_collection, approvalforall_collection, owner_collection, v)
	}
}

func ParseEventLogMongo(transfer_collection *mongo.Collection, approval_collection *mongo.Collection, approvalforall_collection *mongo.Collection, owner_collection *mongo.Collection, log types.Log) {
	database.InsertMongo(transfer_collection, approval_collection, approvalforall_collection, owner_collection, log)
}
