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
	"main/common/tabletypes"
	"main/core/database"
	"main/core/ethclientevent"
	"math/big"
	"strings"
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
	parseHistoryTx(StartTimes)
	listenBlocks()
}

func parseHistoryTx(StartTimes [2]int) {
	log.Print("parsing history\n")
	for index, contract := range config.Contracts {
		go func(contract string, startTime int) {
			log.Printf("parsing history with filter contracts: %s from: %d to：%d \n", contract, startTime, expectBlockNum)
			query := ethereum.FilterQuery{
				FromBlock: big.NewInt(int64(startTime)),
				ToBlock:   big.NewInt(int64(expectBlockNum)),
				Addresses: []common.Address{common.HexToAddress(contract)},
				Topics:    [][]common.Hash{{common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef")}},
			}
			ethclientevent.GetAllTxInfoFromEtheClient(query, eventlogs)
		}(contract, StartTimes[index])
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
			log.Printf("newblocks: %d\n ", header.Number)
			for _, contract := range config.Contracts {
				go func(contract string) {
					log.Printf("newblocks: %d with filter contracts: %s from: %d to: %d\n", header.Number, contract, expectBlockNum, header.Number)
					query := ethereum.FilterQuery{
						FromBlock: big.NewInt(int64(expectBlockNum)),
						ToBlock:   header.Number,
						Addresses: []common.Address{common.HexToAddress(contract)},
						Topics:    [][]common.Hash{{common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef")}},
					}
					ethclientevent.GetAllTxInfoFromEtheClient(query, eventlogs)
				}(contract)
			}
			expectBlockNum = header.Number.Uint64() + 1
		case logs := <-eventlogs:
			ethclientevent.ParseEventLogs(logs)
		}
	}
}

func getStartBlockFromTable() [2]int {
	var resTime [2]int
	for index, contract := range config.Contracts {
		filter := bson.M{"address": strings.ToLower(contract)}
		opts := options.Find().SetSort(bson.M{"blocknumber": -1}).SetLimit(1)
		err, idres := database.GetDocuments(config.DbcollectionTrans, filter, &tabletypes.Transfer{}, opts)
		if err != nil {
			log.Fatalf("Err in getStartBlockFromTable: %s", err)
		}
		if len(idres) > 0 {
			res := idres[0].(*tabletypes.Transfer)
			resTime[index] = int(res.Blocknumber)
		} else {
			resTime[index] = config.StartBlockHeight
		}
	}
	return resTime
}
