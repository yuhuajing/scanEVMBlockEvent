package ethclientevent

import (
	"context"
	"log"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func GetAllTxInfoFromEtheClient(nclient *ethclient.Client, q ethereum.FilterQuery, eventlogs chan<- []types.Log) {
	logs, err := nclient.FilterLogs(context.Background(), q)
	if err != nil {
		log.Fatal(err)
	}
	eventlogs <- logs
	//fmt.Println(len(logs))
}
