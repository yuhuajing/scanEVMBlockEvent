package ethclientevent

import (
	"main/core/database"

	"github.com/ethereum/go-ethereum/core/types"
)

func ParseEventLogs(logs []types.Log) {
	for _, v := range logs {
		ParseEventLog(v)
	}
}

func ParseEventLog(log types.Log) {
	database.Insert(log)
}
