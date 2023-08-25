package ethclientevent

import (
	"main/core/database"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/jinzhu/gorm"
)

func ParseEventLogs(dba *gorm.DB, logs []types.Log) {
	for _, v := range logs {
		ParseEventLog(dba, v)
	}
}

func ParseEventLog(dba *gorm.DB, log types.Log) {
	database.Insert(dba, log)
}
