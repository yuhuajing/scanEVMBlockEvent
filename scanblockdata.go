package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"main/common/config"
	"main/common/dbconn"
	"main/common/ethconn"
	"main/common/tabletypes"
	"main/core/ethclientevent"
	"math/big"
)

var (
	client    = ethconn.ConnBlockchain(config.EthServer)
	headers   = make(chan *types.Header)
	eventlogs = make(chan []types.Log)
)

func main() {
	//go explorer.Explorer()
	transferCollection, approvalCollection, approvalforallCollection, ownerCollection := dbconn.GetCollection()
	StartTimes := getStartBlockFromTable(transferCollection)
	//agStart := getStartBlockFromTable(transfer_collection, config.Agcontract)
	latestblockNum, err := client.BlockNumber(context.Background())
	if err != nil {
		log.Fatalf("Eth connect error:%v", err)
	}
	//go parseHistoryTx(config.Efescontract, efesStart, int(latestblockNum))
	go parseHistoryTx(StartTimes, int(latestblockNum))
	//go listenBlocks(transfer_collection, approval_collection, approvalforall_collection, owner_collection, config.Efescontract)
	go listenBlocks(transferCollection, approvalCollection, approvalforallCollection, ownerCollection)
}

func parseHistoryTx(StartTimes []int, latestblockNum int) {
	for index, contract := range config.Contracts {
		query := ethereum.FilterQuery{
			FromBlock: big.NewInt(int64(StartTimes[index])),
			ToBlock:   big.NewInt(int64(latestblockNum)),
			Addresses: []common.Address{common.HexToAddress(contract)},
		}
		go ethclientevent.GetAllTxInfoFromEtheClient(client, query, eventlogs)
	}
}

func listenBlocks(transferCollection, approvalCollection, approvalforallCollection, ownerCollection *mongo.Collection) {
	subheaders, err := client.SubscribeNewHead(context.Background(), headers)
	if err != nil {
		log.Fatalf("Subscribe Block error: %v", err)
	}

	for {
		select {
		case err := <-subheaders.Err():
			fmt.Errorf("Parse Block error: %v\n", err)
			client = ethconn.ConnBlockchain(config.EthServer)
		case header := <-headers:
			for _, contract := range config.Contracts {
				query := ethereum.FilterQuery{
					FromBlock: big.NewInt(0),
					ToBlock:   big.NewInt(0),
					Addresses: []common.Address{common.HexToAddress(contract)},
				}
				query.FromBlock = header.Number
				query.ToBlock = header.Number
				go ethclientevent.GetAllTxInfoFromEtheClient(client, query, eventlogs)
			}
		case logs := <-eventlogs:
			ethclientevent.ParseEventLogs(transferCollection, approvalCollection, approvalforallCollection, ownerCollection, logs)
		}
	}
}

func getStartBlockFromTable(transferCollection *mongo.Collection) []int {
	var resTime []int
	for index, contract := range config.Contracts {
		filter := bson.D{{Key: "address", Value: contract}}
		opts := options.Find().SetSort(bson.D{{Key: "blocknumber", Value: -1}}).SetLimit(1)
		cur, err := transferCollection.Find(context.TODO(), filter, opts)
		if err != nil {
			log.Fatalf("Err in getStartBlockFromTable: %s", err)
		}
		var res []tabletypes.Transfer
		if err = cur.All(context.Background(), &res); err != nil {
			log.Fatalf("Err parsing data in getStartBlockFromTable: %s", err)
		}
		if len(res) > 0 {
			resTime[index] = int(res[0].Blocknumber)
		} else {
			resTime[index] = config.StartBlockHeight
		}
	}
	return resTime
}
