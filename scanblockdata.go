package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"main/common/config"
	"main/common/dbconn"
	"main/common/tabletypes"
	"main/core/ethclientevent"
	"math/big"
)

var (
	headers        = make(chan *types.Header)
	eventlogs      = make(chan []types.Log)
	expectBlockNum = uint64(0)
)

func main() {
	//go explorer.Explorer()
	latestblockNum, err := config.Client.BlockNumber(context.Background())
	if err != nil {
		log.Fatalf("Eth connect error:%v", err)
	}
	expectBlockNum = latestblockNum
	StartTimes := getStartBlockFromTable()
	parseHistoryTx(StartTimes, int(expectBlockNum))
	listenBlocks()
}

func parseHistoryTx(StartTimes [2]int, latestblockNum int) {
	for index, contract := range config.Contracts {
		query := ethereum.FilterQuery{
			FromBlock: big.NewInt(int64(StartTimes[index])),
			ToBlock:   big.NewInt(int64(latestblockNum)),
			Addresses: []common.Address{common.HexToAddress(contract)},
			Topics:    [][]common.Hash{{common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef")}},
		}
		go ethclientevent.GetAllTxInfoFromEtheClient(query, eventlogs)
	}
}

func listenBlocks() {
	subheaders, err := config.Client.SubscribeNewHead(context.Background(), headers)
	if err != nil {
		log.Fatalf("Subscribe Block error: %v", err)
	}
	for {
		select {
		case err := <-subheaders.Err():
			fmt.Errorf("Parse Block error: %v\n", err)
		case header := <-headers:
			for _, contract := range config.Contracts {
				query := ethereum.FilterQuery{
					FromBlock: big.NewInt(int64(expectBlockNum)),
					ToBlock:   big.NewInt(int64(expectBlockNum)),
					Addresses: []common.Address{common.HexToAddress(contract)},
					Topics:    [][]common.Hash{{common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef")}},
				}
				query.FromBlock = big.NewInt(int64(expectBlockNum))
				query.ToBlock = header.Number
				go ethclientevent.GetAllTxInfoFromEtheClient(query, eventlogs)
				expectBlockNum = header.Number.Uint64() + 1
			}
		case logs := <-eventlogs:
			ethclientevent.ParseEventLogs(logs)
		}
	}
}

func getStartBlockFromTable() [2]int {
	var resTime [2]int
	transferCollection := dbconn.GetCollection(config.DbcollectionTrans)
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
