package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"main/common/config"
	"main/common/dbconn"
	"main/common/ethconn"
	"main/common/tabletypes"
	"main/core/checkaddress"
	"main/core/ethclientevent"
	"main/explorer"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	startBlockHeight, latestblockNum uint64
	client                           *ethclient.Client
	contractaddress                  string
	blockchain                       string
	rpcserver                        string
	database                         string
)

func main() {
	flag.StringVar(&blockchain, "blockchain", "ethereum", "Blockchain,for example: polygon,ethereum. default ethereum")
	flag.StringVar(&rpcserver, "rpcserver", "", "Blockchain RPC server")
	flag.StringVar(&contractaddress, "address", "0xff2b4721f997c242ff406a626f17df083bd2c568", "Smart contract address")
	flag.Uint64Var(&startBlockHeight, "startblock", 0, "if the transfer table is empty, the startblock is 17948500 by default, else the startblock equals to the blocknumer of the table last data filtered by contract address")
	flag.StringVar(&database, "database", "mongodb", "database type to store event logs, default mongodb")

	flag.Parse()

	if database != "mysql" && database != "mongodb" {
		fmt.Println("Not valid database, please select database from: mysql, mongodb")
		return
	}

	if blockchain != "" {
		switch blockchain {
		case "ethereum":
			client = ethconn.ConnBlockchain(config.EthServer)
		default:
			fmt.Println("This blockchain is not supported until now")
			return
		}
	} else {
		client = ethconn.ConnBlockchain(config.EthServer)
	}

	if !checkaddress.IsContractAddress(contractaddress, client) {
		fmt.Println("address should be a valid smart contract address")
		return
	}

	if rpcserver != "" {
		client = ethconn.ConnBlockchain(rpcserver)
	}

	headers := make(chan *types.Header) // listening new blocks
	subheaders, err := client.SubscribeNewHead(context.Background(), headers)
	if err != nil {
		fmt.Println(fmt.Errorf("Subscribe Block error: %v", err))
		client = ethconn.ConnBlockchain(config.EthServer)
	}

	go explorer.Explorer()
	_tablelatestBlockNum := uint64(0)
	latestblockNum, _ = client.BlockNumber(context.Background())
	res := []tabletypes.Transfer{}
	var dba *gorm.DB
	var transfer_collection *mongo.Collection
	var approval_collection *mongo.Collection
	var approvalforall_collection *mongo.Collection
	var owner_collection *mongo.Collection

	if database == "mysql" {
		dba = dbconn.Buildconnect()
		dba.AutoMigrate(&tabletypes.Transfer{}, &tabletypes.Approval{}, &tabletypes.ApprovalForAll{}, &tabletypes.Owner{})

		dba.Model(&tabletypes.Transfer{}).Where("address = ?", contractaddress).Order("blocknumber desc").Limit(1).Find(&res)

	} else {
		client = ethconn.ConnBlockchain(config.EthServer)
		transfer_collection, approval_collection, approvalforall_collection, owner_collection = dbconn.GetCollection()
		filter := bson.D{{Key: "address", Value: config.Address}}
		opts := options.Find().SetSort(bson.D{{Key: "blocknumber", Value: -1}}).SetLimit(1)
		cur, err := transfer_collection.Find(context.TODO(), filter, opts)
		if err != nil {
			log.Fatal(err)
		}
		if err = cur.All(context.Background(), &res); err != nil {
			log.Fatal(err)
		}
	}

	if len(res) > 0 {
		_tablelatestBlockNum = res[0].Blocknumber
	}

	if _tablelatestBlockNum == 0 {
		startBlockHeight = uint64(17948500)
	} else {
		startBlockHeight = _tablelatestBlockNum
	}

	eventlogs := make(chan []types.Log)
	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(int64(startBlockHeight)),
		ToBlock:   big.NewInt(int64(latestblockNum)),
		Addresses: []common.Address{common.HexToAddress(contractaddress)},
		//Topics:    [][]common.Hash{{common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef")}},
	}
	go ethclientevent.GetAllTxInfoFromEtheClient(client, query, eventlogs)

	// logs := make(chan types.Log)
	// sublogs, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	// if err != nil {
	// 	fmt.Println(fmt.Errorf("Subscribe Log error: %v", err))
	// 	client = ethconn.ConnBlockchain(config.EthServer)
	// 	//log.Fatal(err)
	// }

	for {
		select {
		case err := <-subheaders.Err():
			fmt.Println(fmt.Errorf("Parse Block error: %v", err))
			client = ethconn.ConnBlockchain(config.EthServer)
		// case err := <-sublogs.Err():
		// 	fmt.Println(fmt.Errorf("Parse Log error: %v", err))
		// 	client = ethconn.ConnBlockchain(config.EthServer)
		case header := <-headers:
			query.FromBlock = header.Number
			query.ToBlock = header.Number
			go ethclientevent.GetAllTxInfoFromEtheClient(client, query, eventlogs)
		// case log := <-logs:
		// 	ethclientevent.ParseEventLog(dba, log)
		case logs := <-eventlogs:
			if database == "mysql" {
				ethclientevent.ParseEventLogsMysql(dba, logs)
			} else {
				ethclientevent.ParseEventLogsMongo(transfer_collection, approval_collection, approvalforall_collection, owner_collection, logs)
			}

		}
	}

}
