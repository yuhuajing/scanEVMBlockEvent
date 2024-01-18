package ethclientevent

import (
	"context"
	"log"
	"main/common/config"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/core/types"
)

func GetAllTxInfoFromEtheClient(q ethereum.FilterQuery, eventlogs chan<- []types.Log) {
	logs, err := config.Client.FilterLogs(context.Background(), q)
	if err != nil {
		log.Fatal(err)
	}
	eventlogs <- logs
}
