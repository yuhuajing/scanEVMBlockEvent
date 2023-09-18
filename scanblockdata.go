package main

import (
	"context"
	"fmt"
	"log"
	"main/common/config"
	"main/common/dbconn"
	"main/common/ethconn"
	"main/common/tabletypes"
	"main/core/ethclientevent"
	"main/explorer"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	startBlockHeight, latestblockNum uint64
	client                           *ethclient.Client
)

func main() {
	client = ethconn.ConnBlockchain(config.EthServer)

	headers := make(chan *types.Header) // listening new blocks
	subheaders, err := client.SubscribeNewHead(context.Background(), headers)
	if err != nil {
		fmt.Println(fmt.Errorf("Subscribe Block error: %v", err))
		client = ethconn.ConnBlockchain(config.EthServer)
	}

	go explorer.Explorer()

	client = ethconn.ConnBlockchain(config.EthServer)
	transfer_collection, approval_collection, approvalforall_collection, owner_collection := dbconn.GetCollection()

	_tablelatestBlockNum := uint64(0)

	filter := bson.D{{Key: "address", Value: config.Address}}
	opts := options.Find().SetSort(bson.D{{Key: "blocknumber", Value: -1}}).SetLimit(1)
	cur, err := transfer_collection.Find(context.TODO(), filter, opts)
	if err != nil {
		log.Fatal(err)
	}
	var res []tabletypes.Transfer

	if err = cur.All(context.Background(), &res); err != nil {
		log.Fatal(err)
	}

	if len(res) > 0 {
		_tablelatestBlockNum = res[0].Blocknumber
	}

	latestblockNum, _ = client.BlockNumber(context.Background())

	if _tablelatestBlockNum == 0 {
		startBlockHeight = uint64(17948500)
	} else {
		startBlockHeight = _tablelatestBlockNum
	}

	eventlogs := make(chan []types.Log)
	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(int64(startBlockHeight)), //big.NewInt(int64(startBlockHeight)),
		ToBlock:   big.NewInt(int64(latestblockNum)),
		Addresses: []common.Address{common.HexToAddress(config.Address)},
		//Topics:    [][]common.Hash{{common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef")}},
	}
	go ethclientevent.GetAllTxInfoFromEtheClient(client, query, eventlogs)

	for {
		select {
		case err := <-subheaders.Err():
			fmt.Println(fmt.Errorf("Parse Block error: %v", err))
			client = ethconn.ConnBlockchain(config.EthServer)
		case header := <-headers:
			query.FromBlock = header.Number
			query.ToBlock = header.Number
			go ethclientevent.GetAllTxInfoFromEtheClient(client, query, eventlogs)
		case logs := <-eventlogs:
			ethclientevent.ParseEventLogs(transfer_collection, approval_collection, approvalforall_collection, owner_collection, logs)
		}
	}

}
