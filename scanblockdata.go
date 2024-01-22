package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/foundVanting/opensea-stream-go/entity"
	"github.com/foundVanting/opensea-stream-go/opensea"
	openseatypes "github.com/foundVanting/opensea-stream-go/types"
	"github.com/mitchellh/mapstructure"
	"github.com/xiaowang7777/phx"
	"go.mongodb.org/mongo-driver/bson"
	"io"
	"log"
	"main/common/config"
	"main/common/tabletypes"
	"main/core/database"
	"main/core/ethclientevent"
	"math"
	"math/big"
	"net/http"
	"strings"
	"sync"
)

var (
	headers        = make(chan *types.Header)
	eventlogs      = make(chan []types.Log)
	expectBlockNum = uint64(0)
	latestblockNum = uint64(0)
	err            error
)

func init() {
	latestblockNum, err = config.Client.BlockNumber(context.Background())
	if err != nil {
		log.Fatalf("Eth connect error:%v", err)
	}
	expectBlockNum = latestblockNum + 1

}

func main() {
	//go explorer.Explorer()
	//subOpensea()
	parseOpenseaListing("0x1aae1A668c92Eb411eAfD80DD0c60ca67ad17a1c", "0x3ba331b7026b60e5f92f5a039cdded8e52c90cd3", []int{195})

	//var wg sync.WaitGroup
	//wg.Add(2)
	//
	//go func() {
	//	defer wg.Done()
	//	StartTimes := getStartBlockFromTable()
	//	parseHistoryTx(StartTimes)
	//}()
	//
	//go func() {
	//	defer wg.Done()
	//	listenBlocks()
	//}()
	//
	//wg.Wait()
}

type Order struct {
	ListingTime    int64  `json:"listing_time"`
	ExpirationTime int64  `json:"expiration_time"`
	OrderHash      string `json:"order_hash"`
}

type Response struct {
	Next     interface{} `json:"next"`
	Previous interface{} `json:"previous"`
	Orders   []struct {
		ListingTime    int    `json:"listing_time"`
		ExpirationTime int    `json:"expiration_time"`
		OrderHash      string `json:"order_hash"`
		ProtocolData   struct {
			Parameters struct {
				Offerer string `json:"offerer"`
				Offer   []struct {
					Token                string `json:"token"`
					IdentifierOrCriteria string `json:"identifierOrCriteria"`
				} `json:"offer"`
			} `json:"parameters"`
		} `json:"protocol_data"`
	} `json:"orders"`
}

func parseOpenseaListing(contractaddress, owner string, nftid []int) {
	openseaUrl := "https://api.opensea.io//api/v2/orders/ethereum/seaport/listings?asset_contract_address=%s&maker=%s"
	openseaUrl = fmt.Sprintf(openseaUrl, contractaddress, owner)
	tokenId := "&token_ids=%d"
	idLen := len(nftid)
	index := 0
	for idLen > 0 && index < idLen {
		targetIdUrl := ""
		targetV := math.Min(float64(10), float64(idLen))
		idLen -= int(targetV)
		for targetV > 0 {
			targetIdUrl += fmt.Sprintf(tokenId, nftid[index])
			index++
			targetV--
		}
		url := openseaUrl + targetIdUrl
		//fmt.Println(url)
		req, _ := http.NewRequest("GET", url, nil)
		req.Header.Add("accept", "application/json")
		req.Header.Add("x-api-key", "9602c2e9de24426196b5c317099155c7")
		res, _ := http.DefaultClient.Do(req)
		defer res.Body.Close()
		body, _ := io.ReadAll(res.Body)
		var response Response
		err = json.Unmarshal(body, &response)
		if err != nil {
			fmt.Println("Error:", err)
		}
		fmt.Println(response)
	}
}

func subOpensea() {
	client := opensea.NewStreamClient(openseatypes.MAINNET, "9602c2e9de24426196b5c317099155c7", phx.LogInfo, func(err error) {
		fmt.Println("opensea.NewStreamClient err:", err)
	})
	if err := client.Connect(); err != nil {
		fmt.Println("client.Connect err:", err)
		return
	}

	client.OnItemListed("Enforcer Founder Edition Spaceship", func(response any) {
		var itemListedEvent entity.ItemListedEvent
		err := mapstructure.Decode(response, &itemListedEvent)
		if err != nil {
			fmt.Println("mapstructure.Decode err:", err)
		}
		fmt.Printf("%+v\n", itemListedEvent)
	})

	//client.OnItemSold("collection-slug", func(response any) {
	//	var itemSoldEvent entity.ItemSoldEvent
	//	err := mapstructure.Decode(response, &itemSoldEvent)
	//	if err != nil {
	//		fmt.Println("mapstructure.Decode err:", err)
	//	}
	//	fmt.Printf("%+v\n", itemSoldEvent)
	//})
	//
	//client.OnItemTransferred("collection-slug", func(response any) {
	//	var itemTransferredEvent entity.ItemTransferredEvent
	//	err := mapstructure.Decode(response, &itemTransferredEvent)
	//	if err != nil {
	//		fmt.Println("mapstructure.Decode err:", err)
	//	}
	//	fmt.Printf("%+v\n", itemTransferredEvent)
	//})

	client.OnItemCancelled("Enforcer Founder Edition Spaceship", func(response any) {
		var itemCancelledEvent entity.ItemCancelledEvent
		err := mapstructure.Decode(response, &itemCancelledEvent)
		if err != nil {
			fmt.Println("mapstructure.Decode err:", err)
		}
		fmt.Printf("%+v\n", itemCancelledEvent)
	})

	//client.OnItemReceivedBid("collection-slug", func(response any) {
	//	var itemReceivedBidEvent entity.ItemReceivedBidEvent
	//	err := mapstructure.Decode(response, &itemReceivedBidEvent)
	//	if err != nil {
	//		fmt.Println("mapstructure.Decode err:", err)
	//	}
	//	fmt.Printf("%+v\n", itemReceivedBidEvent)
	//})
	//client.OnItemReceivedOffer("collection-slug", func(response any) {
	//	var itemReceivedOfferEvent entity.ItemReceivedOfferEvent
	//	err := mapstructure.Decode(response, &itemReceivedOfferEvent)
	//	if err != nil {
	//		fmt.Println("mapstructure.Decode err:", err)
	//	}
	//	fmt.Printf("%+v\n", itemReceivedOfferEvent)
	//})
	//
	//client.OnItemMetadataUpdated("collection-slug", func(response any) {
	//	var itemMetadataUpdateEvent entity.ItemMetadataUpdateEvent
	//	err := mapstructure.Decode(response, &itemMetadataUpdateEvent)
	//	if err != nil {
	//		fmt.Println("mapstructure.Decode err:", err)
	//	}
	//	fmt.Printf("%+v\n", itemMetadataUpdateEvent)
	//})

	select {}
}

func parseHistoryTx(StartTimes [2]int) {
	log.Print("parsing history\n")
	//var wg sync.WaitGroup
	for index, contract := range config.Contracts {
		//wg.Add(1)
		tmpContract := contract
		tmpStartTime := StartTimes[index]
		go func(contract string, startTime int) {
			//	defer wg.Done()
			log.Printf("parsing history with filter contracts: %s from: %d to: %d\n", contract, startTime, latestblockNum)
			query := ethereum.FilterQuery{
				FromBlock: big.NewInt(int64(startTime)),
				ToBlock:   big.NewInt(int64(latestblockNum)),
				Addresses: []common.Address{common.HexToAddress(contract)},
				Topics:    [][]common.Hash{{common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef")}},
			}
			go ethclientevent.GetAllTxInfoFromEtheClient(query, eventlogs)
			database.CreOrUpdateStartBlock(contract, latestblockNum)

		}(tmpContract, tmpStartTime)
	}
	//	wg.Wait()
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
			if header.Number.Uint64() > latestblockNum {
				log.Printf("newblocks: %d\n ", header.Number)
				var wg sync.WaitGroup
				for _, contract := range config.Contracts {
					wg.Add(1)
					tmpContract := contract
					go func(contract string) {
						defer wg.Done()
						log.Printf("newblocks: %d with filter contracts: %s from: %d to: %d\n", header.Number, contract, expectBlockNum, header.Number)
						query := ethereum.FilterQuery{
							FromBlock: big.NewInt(int64(expectBlockNum)),
							ToBlock:   header.Number,
							Addresses: []common.Address{common.HexToAddress(contract)},
							Topics:    [][]common.Hash{{common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef")}},
						}
						go ethclientevent.GetAllTxInfoFromEtheClient(query, eventlogs)
						database.CreOrUpdateStartBlock(contract, header.Number.Uint64())
					}(tmpContract)
				}

				wg.Wait()
				expectBlockNum = header.Number.Uint64() + 1
			}
		case logs := <-eventlogs:
			ethclientevent.ParseEventLogs(logs)
		}
	}
}

func getStartBlockFromTable() [2]int {
	var resTime [2]int
	for index, contract := range config.Contracts {
		filter := bson.M{"address": strings.ToLower(contract)}
		//opts := options.Find().SetSort(bson.M{"blocknumber": -1}).SetLimit(1)
		err, idres := database.GetDocuments(config.DbcollectionSB, filter, &tabletypes.Startblocks{})
		if err != nil {
			log.Fatalf("Err in getStartBlockFromTable: %s", err)
		}
		if len(idres) > 0 {
			res := idres[0].(*tabletypes.Startblocks)
			resTime[index] = int(res.Blocknumber)
		} else {
			resTime[index] = config.StartBlockHeight
		}
	}
	return resTime
}
