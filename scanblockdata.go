package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"log"
	"main/common/config"
	"main/core/database"
	"main/core/ethclientevent"
	"main/openseaorder"
	"math/big"
	"strings"
	"sync"
)

var (
	headers   = make(chan *types.Header)
	eventlogs = make(chan []types.Log)
)

//func init() {
//	latestblockNum, err := config.Client.BlockNumber(context.Background())
//	if err != nil {
//		log.Fatalf("Eth connect error:%v", err)
//	}
//
//	//database.CreOrUpdateLatestBlock(latestblockNum)
//}

func main() {
	//go explorer.Explorer()
	var wg sync.WaitGroup
	wg.Add(3)

	go func() {
		defer wg.Done()
		listenBlocks()
	}()

	go func() {
		defer wg.Done()
		parseOpenseaOrdersByCollection()
	}()

	go func() {
		defer wg.Done()
		updateOwnerTimestamp()
	}()

	wg.Wait()

	//parseMarketOrders()
	//ids, holdTime, _ := openseaorder.HoldTime(config.Contracts[1], "0x08d8db85ad681fa6a80c0d1fab9312f00d1a1888")
	//fmt.Println(ids)
	//fmt.Println(holdTime)
	//database.UpdateOwner("0x1aae1a668c92eb411eafd80dd0c60ca67ad17a10")
	//for _, contract := range config.Contracts {
	//	database.UpdateOwner(contract)
	//}
}

func listenBlocks() {
	//DealHistoryTx()
	subheaders, err := config.Client.SubscribeNewHead(context.Background(), headers)
	if err != nil {
		log.Fatalf("Subscribe Block error: %v", err)
	}
	for {
		select {
		case err := <-subheaders.Err():
			fmt.Errorf("Parse Block error: %v\n", err)
		case header := <-headers:
			config.BlockWithTimestamp[header.Number.Uint64()] = header.Time
			log.Printf("newblocks: %d\n ", header.Number)
			for _, contract := range config.Contracts {
				expectBlockNum := database.GetStartBlockNumber(contract)
				if header.Number.Uint64() > expectBlockNum {
					log.Printf("newblocks: %d with filter contracts: %s from: %d to: %d\n", header.Number, contract, expectBlockNum, header.Number)
					query := ethereum.FilterQuery{
						FromBlock: big.NewInt(int64(expectBlockNum)),
						ToBlock:   header.Number,
						Addresses: []common.Address{common.HexToAddress(contract)},
						Topics:    [][]common.Hash{{common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef")}},
					}
					go ethclientevent.GetAllTxInfoFromEtheClient(query, eventlogs)
					database.CreOrUpdateStartBlock(contract, header.Number.Uint64())
				}
			}
		case logs := <-eventlogs:
			ethclientevent.ParseEventLogs(logs)
		}
	}
}

func parseOpenseaOrdersByCollection() {
	for _, contract := range config.Contracts {
		coll := config.Collections[strings.ToLower(contract)]
		log.Printf("parsing opensea orders with collection: %s", coll)
		openseaorder.CreatOrUpdateOpenseaListingByhash(coll)
	}
	openseaorder.SubOpensea()
}

func updateOwnerTimestamp() {
	//time.Sleep(30 * time.Second)
	for {
		if database.UpdateOwnerTimestamp() {
			break
		}
	}
}

//
//func parseMarketOrders() {
//	marketorder.ParseOrderListing()
//	//marketorder.SubOrder()
//}
//
//func parseOpenseaOrders() {
//	for _, contractaddress := range config.Contracts {
//		fmt.Println("parsing contract opensea orders", contractaddress)
//		config.NftOwners[strings.ToLower(contractaddress)] = make(map[string]bool)
//		for i := 0; i < config.ContractSupply[strings.ToLower(contractaddress)]; i++ {
//			if contractaddress == "0x1aae1a668c92eb411eafd80dd0c60ca67ad17a1c" {
//				break
//			}
//			allids, owner, err := database.GetOwnerByNFTId(contractaddress, i)
//			if err != nil {
//				log.Fatalf("database.GetOwnerByNFTId error: %v", err)
//			}
//			if len(allids) > 0 {
//				openseaorder.ParseOpenseaListing(contractaddress, owner, allids)
//			}
//		}
//	}
//}

//func parseHistoryTx(StartTimes [2]int) {
//	var wg sync.WaitGroup
//	latestblockNum := database.GetLatestBlockNumber()
//	for index, contract := range config.Contracts {
//		wg.Add(1)
//		tmpContract := contract
//		tmpStartTime := StartTimes[index]
//		go func(contract string, startTime int) {
//			defer wg.Done()
//			log.Printf("parsing history with filter contracts: %s from: %d to: %d\n", contract, startTime, latestblockNum)
//			query := ethereum.FilterQuery{
//				FromBlock: big.NewInt(int64(startTime)),
//				ToBlock:   big.NewInt(int64(latestblockNum)),
//				Addresses: []common.Address{common.HexToAddress(contract)},
//				Topics:    [][]common.Hash{{common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef")}},
//			}
//			go ethclientevent.GetAllTxInfoFromEtheClient(query, eventlogs)
//		}(tmpContract, tmpStartTime)
//	}
//	wg.Wait()
//}

//func DealHistoryTx() {
//	for _, contract := range config.Contracts {
//		startTime := database.GetStartBlockNumber(contract)
//		//latestblockNum := database.GetLatestBlockNumber()
//		log.Printf("parsing history with filter contracts: %s from: %d to: %d\n", contract, startTime, latestblockNum)
//		query := ethereum.FilterQuery{
//			FromBlock: big.NewInt(int64(startTime)),
//			ToBlock:   big.NewInt(int64(latestblockNum)),
//			Addresses: []common.Address{common.HexToAddress(contract)},
//			Topics:    [][]common.Hash{{common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef")}},
//		}
//		go ethclientevent.GetAllTxInfoFromEtheClient(query, eventlogs)
//	}
//}
//
//func getStartBlockFromTable() [2]int {
//	var resTime [2]int
//	for index, contract := range config.Contracts {
//		filter := bson.M{"address": strings.ToLower(contract)}
//		//opts := options.Find().SetSort(bson.M{"blocknumber": -1}).SetLimit(1)
//		err, idres := database.GetDocuments(config.DbcollectionSB, filter, &tabletypes.Startblocks{})
//		if err != nil {
//			log.Fatalf("Err in getStartBlockFromTable: %s", err)
//		}
//		if len(idres) > 0 {
//			res := idres[0].(*tabletypes.Startblocks)
//			resTime[index] = int(res.Historyblocknumber)
//		} else {
//			resTime[index] = int(config.ContractDeployHeight[strings.ToLower(contract)])
//		}
//	}
//	return resTime
//}
