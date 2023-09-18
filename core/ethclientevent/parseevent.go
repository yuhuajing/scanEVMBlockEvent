package ethclientevent

import (
	"main/core/database"

	"github.com/ethereum/go-ethereum/core/types"
	"go.mongodb.org/mongo-driver/mongo"
)

func ParseEventLogs(transfer_collection *mongo.Collection, approval_collection *mongo.Collection, approvalforall_collection *mongo.Collection, owner_collection *mongo.Collection, logs []types.Log) {
	for _, v := range logs {
		ParseEventLog(transfer_collection, approval_collection, approvalforall_collection, owner_collection, v)
	}
}

func ParseEventLog(transfer_collection *mongo.Collection, approval_collection *mongo.Collection, approvalforall_collection *mongo.Collection, owner_collection *mongo.Collection, log types.Log) {
	database.Insert(transfer_collection, approval_collection, approvalforall_collection, owner_collection, log)
}
