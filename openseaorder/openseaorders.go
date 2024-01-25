package openseaorder

import (
	"encoding/json"
	"fmt"
	"github.com/foundVanting/opensea-stream-go/entity"
	"github.com/foundVanting/opensea-stream-go/opensea"
	openseatypes "github.com/foundVanting/opensea-stream-go/types"
	"github.com/mitchellh/mapstructure"
	"github.com/xiaowang7777/phx"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"io"
	"log"
	"main/common/config"
	"main/common/tabletypes"
	"main/core/database"
	"main/core/ethclientevent"
	"math"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type Listresponse struct {
	Orders []struct {
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

func ParseOpenseaListing(contractaddress, owner string, nftid []int) {
	openseaUrl := "https://api.opensea.io//api/v2/orders/ethereum/seaport/listings?asset_contract_address=%s&maker=%s"
	openseaUrl = fmt.Sprintf(openseaUrl, contractaddress, owner)
	tokenId := "&token_ids=%d"
	idLen := len(nftid)
	index := 0
	for idLen > 0 && index < len(nftid) {
		targetIdUrl := ""
		targetLen := math.Min(float64(10), float64(idLen))
		idLen -= int(targetLen)
		for targetLen > 0 {
			targetIdUrl += fmt.Sprintf(tokenId, nftid[index])
			index++
			targetLen--
		}
		url := openseaUrl + targetIdUrl
		//fmt.Println(url)
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			log.Fatalf("err in ParseOpenseaListingByCollection: %v", err)
		}
		req.Header.Add("accept", "application/json")
		req.Header.Add("x-api-key", config.OpenseaToken)
		res, _ := http.DefaultClient.Do(req)
		defer res.Body.Close()
		body, _ := io.ReadAll(res.Body)
		var response Listresponse
		err = json.Unmarshal(body, &response)
		if err != nil {
			log.Fatalf("unmarshal opensea data error: %v", err)
		}
		for _, order := range response.Orders {
			err = database.AddOpenSeaOrder(order.OrderHash, order.ProtocolData.Parameters.Offer[0].Token, order.ProtocolData.Parameters.Offerer, order.ProtocolData.Parameters.Offer[0].IdentifierOrCriteria, order.ListingTime, order.ExpirationTime)
			if err != nil {
				log.Fatalf("database.AddOpenSeaOrder error: %v", err)
			}
		}

		//	fmt.Println(response)
	}
}

type listbycoll struct {
	Listings []struct {
		OrderHash string `json:"order_hash"`
		Chain     string `json:"chain"`
		//	Type      string `json:"type"`
		Price struct {
			Current struct {
				Currency string `json:"currency"`
				Decimals int    `json:"decimals"`
				Value    string `json:"value"`
			} `json:"current"`
		} `json:"price"`
		ProtocolData struct {
			Parameters struct {
				Offerer string `json:"offerer"`
				Offer   []struct {
					//ItemType             int    `json:"itemType"`
					Token                string `json:"token"`
					IdentifierOrCriteria string `json:"identifierOrCriteria"`
					//StartAmount          string `json:"startAmount"`
					//	EndAmount            string `json:"endAmount"`
				} `json:"offer"`
				//Consideration []struct {
				//	ItemType             int    `json:"itemType"`
				//	Token                string `json:"token"`
				//	IdentifierOrCriteria string `json:"identifierOrCriteria"`
				//	StartAmount          string `json:"startAmount"`
				//	EndAmount            string `json:"endAmount"`
				//	Recipient            string `json:"recipient"`
				//} `json:"consideration"`
				StartTime string `json:"startTime"`
				EndTime   string `json:"endTime"`
				//OrderType                       int    `json:"orderType"`
				//Zone                            string `json:"zone"`
				//ZoneHash                        string `json:"zoneHash"`
				//Salt                            string `json:"salt"`
				//ConduitKey                      string `json:"conduitKey"`
				//TotalOriginalConsiderationItems int    `json:"totalOriginalConsiderationItems"`
				//Counter                         int    `json:"counter"`
			} `json:"parameters"`
			//Signature interface{} `json:"signature"`
		} `json:"protocol_data"`
		//ProtocolAddress string `json:"protocol_address"`
	} `json:"listings"`
}

func CreatOrUpdateOpenseaListingByhash(collection string) {
	openseaListingOrders := ParseOpenseaListingByCollection(collection)
	listingOrdersFromOpensea := make(map[string]bool)
	for _, order := range openseaListingOrders.Listings {
		listingOrdersFromOpensea[order.OrderHash] = true
		startTime, _ := strconv.ParseInt(order.ProtocolData.Parameters.StartTime, 10, 64)
		endTime, _ := strconv.ParseInt(order.ProtocolData.Parameters.EndTime, 10, 64)
		err := database.AddOpenSeaOrder(order.OrderHash, order.ProtocolData.Parameters.Offer[0].Token, order.ProtocolData.Parameters.Offerer, order.ProtocolData.Parameters.Offer[0].IdentifierOrCriteria, int(startTime), int(endTime))
		if err != nil {
			log.Fatalf("database.AddOpenSeaOrder error: %v", err)
		}
	}
	listingOrdersFromDB, _ := database.GetOpenSeaOrders()
	for _, hash := range listingOrdersFromDB {
		if listingOrdersFromOpensea[hash] {
			database.UpdateOpenSeaOrderByHash(hash, tabletypes.StatusInvalid)
		}
	}
}

func ParseOpenseaListingByCollection(collection string) *listbycoll {
	openseaUrl := "https://api.opensea.io/api/v2/listings/collection/%s/all"
	url := fmt.Sprintf(openseaUrl, collection)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalf("err in ParseOpenseaListingByCollection: %v", err)
	}
	req.Header.Add("accept", "application/json")
	req.Header.Add("x-api-key", config.OpenseaToken)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalf("err in Request OpenseaListingByCollection: %v", err)
	}
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	var response listbycoll
	err = json.Unmarshal(body, &response)
	if err != nil {
		log.Fatalf("unmarshal opensea data error: %v", err)
	}
	return &response
	//orderhashs := make([]string, 0)
	//for _, order := range response.Listings {
	//	//startTime, _ := strconv.ParseInt(order.ProtocolData.Parameters.StartTime, 10, 64)
	//	//endTime, _ := strconv.ParseInt(order.ProtocolData.Parameters.EndTime, 10, 64)
	//	//err = database.AddOpenSeaOrder(order.OrderHash, order.ProtocolData.Parameters.Offer[0].Token, order.ProtocolData.Parameters.Offerer, order.ProtocolData.Parameters.Offer[0].IdentifierOrCriteria, int(startTime), int(endTime))
	//	//if err != nil {
	//	//	log.Fatalf("database.AddOpenSeaOrder error: %v", err)
	//	//}
	//	orderhashs = append(orderhashs, order.OrderHash)
	//}
	//return orderhashs
}

func SubOpensea() {
	client := opensea.NewStreamClient(openseatypes.MAINNET, config.OpenseaToken, phx.LogInfo, func(err error) {
		fmt.Println("opensea.NewStreamClient err:", err)
	})
	if err := client.Connect(); err != nil {
		fmt.Println("client.Connect err:", err)
		return
	}

	client.OnItemListed(config.EfesCollections, func(response any) {
		var itemListedEvent entity.ItemListedEvent
		err := mapstructure.Decode(response, &itemListedEvent)
		if err != nil {
			fmt.Println("mapstructure.Decode listing err:", err)
		}
		orderhash := itemListedEvent.Payload.OrderHash

		//chainName := strings.Split(itemListedEvent.Payload.Item.NftId, "/")[0]
		contract := strings.Split(itemListedEvent.Payload.Item.NftId, "/")[1]
		nftID := strings.Split(itemListedEvent.Payload.Item.NftId, "/")[2]
		owner := itemListedEvent.Payload.Maker.Address
		listingTime, err := parseTime(itemListedEvent.Payload.ListingDate)
		if err != nil {
			fmt.Println("parseTime ListingDate err:", err)
		}
		expireTime, err := parseTime(itemListedEvent.Payload.ExpirationDate)
		if err != nil {
			fmt.Println("parseTime ExpirationDate err:", err)
		}
		err = database.AddOpenSeaOrder(orderhash, contract, owner, nftID, listingTime, expireTime)
		if err != nil {
			log.Fatalf("database.AddOpenSeaOrder error: %v", err)
		}
	})

	client.OnItemCancelled(config.EfesCollections, func(response any) {
		var itemCancelledEvent entity.ItemCancelledEvent
		err := mapstructure.Decode(response, &itemCancelledEvent)
		if err != nil {
			fmt.Println("mapstructure.Decode err:", err)
		}
		err = database.CancelOpenSeaOrder(itemCancelledEvent.Payload.OrderHash)
		if err != nil {
			log.Fatalf("database.CancelOpenSeaOrder error: %v", err)
		}
		//fmt.Printf("%+v\n", itemCancelledEvent.Payload.OrderHash)
	})

	client.OnItemListed(config.AgCollections, func(response any) {
		var itemListedEvent entity.ItemListedEvent
		err := mapstructure.Decode(response, &itemListedEvent)
		if err != nil {
			fmt.Println("mapstructure.Decode listing err:", err)
		}
		orderhash := itemListedEvent.Payload.OrderHash

		//chainName := strings.Split(itemListedEvent.Payload.Item.NftId, "/")[0]
		contract := strings.Split(itemListedEvent.Payload.Item.NftId, "/")[1]
		nftID := strings.Split(itemListedEvent.Payload.Item.NftId, "/")[2]
		owner := itemListedEvent.Payload.Maker.Address
		listingTime, err := parseTime(itemListedEvent.Payload.ListingDate)
		if err != nil {
			fmt.Println("parseTime ListingDate err:", err)
		}
		expireTime, err := parseTime(itemListedEvent.Payload.ExpirationDate)
		if err != nil {
			fmt.Println("parseTime ExpirationDate err:", err)
		}
		err = database.AddOpenSeaOrder(orderhash, contract, owner, nftID, listingTime, expireTime)
		if err != nil {
			log.Fatalf("database.AddOpenSeaOrder error: %v", err)
		}
	})

	client.OnItemCancelled(config.AgCollections, func(response any) {
		var itemCancelledEvent entity.ItemCancelledEvent
		err := mapstructure.Decode(response, &itemCancelledEvent)
		if err != nil {
			fmt.Println("mapstructure.Decode err:", err)
		}
		err = database.CancelOpenSeaOrder(itemCancelledEvent.Payload.OrderHash)
		if err != nil {
			log.Fatalf("database.CancelOpenSeaOrder error: %v", err)
		}
		//	fmt.Printf("%+v\n", itemCancelledEvent.Payload.OrderHash)
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

func parseTime(timestr string) (int, error) {
	t, err := time.Parse(time.RFC3339Nano, timestr)
	if err != nil {
		return 0, err
	}
	return int(t.Unix()), nil
}

func HoldTime(contract, owner string) ([]int, []uint64, error) {
	NftIds, IdTime, err := holdNFTs(contract, owner)
	holdTime := make([]uint64, 0)
	if err != nil || len(NftIds) == 0 {
		return NftIds, holdTime, err
	}
	timenow := uint64(time.Now().Unix())
	for _, id := range NftIds {
		filter := bson.M{"address": strings.ToLower(contract), "owner": strings.ToLower(owner), "tokenid": id, "status": tabletypes.StatusListing}
		opts := options.Find().SetSort(bson.D{{Key: "expirationtime", Value: 1}}).SetLimit(1)
		err, idres := database.GetDocuments(config.DbcollectionOpensea, filter, &tabletypes.OpenseaOrder{}, opts)
		if err != nil {
			fmt.Println(err)
		}
		if len(idres) > 0 {
			v := idres[0].(*tabletypes.OpenseaOrder)
			if int(timenow) < v.Expirationtime {
				holdTime = append(holdTime, 0)
			}
		} else {
			holdTime = append(holdTime, timenow-IdTime[id])
		}
	}
	return NftIds, holdTime, nil
}

func holdNFTs(contract, owner string) ([]int, map[int]uint64, error) {
	NFTIds := make([]int, 0)
	IdTime := make(map[int]uint64)
	filter := bson.M{"address": strings.ToLower(contract), "owner": strings.ToLower(owner)}
	err, idres := database.GetDocuments(config.DbcollectionOwner, filter, &tabletypes.Owner{})
	if err != nil {
		return NFTIds, IdTime, fmt.Errorf("Read ownerData err: %v", err)
	}
	if len(idres) > 0 {
		for _, res := range idres {
			v := res.(*tabletypes.Owner)
			blocktimestamp, err := blockTimestamp(v.Blocknumber)
			if err != nil {
				blocktimestamp, _ = ethclientevent.ChainBlockTime(v.Blocknumber)
			}
			NFTIds = append(NFTIds, v.Tokenid)
			IdTime[v.Tokenid] = blocktimestamp
		}
	}
	return NFTIds, IdTime, nil
}

func blockTimestamp(blocknumber uint64) (uint64, error) {
	filter := bson.M{"blocknumber": blocknumber}
	opts := options.Find().SetLimit(1)
	err, idres := database.GetDocuments(config.DbcollectionTrans, filter, &tabletypes.Transfer{}, opts)
	if err != nil {
		return 0, fmt.Errorf("Read transData err: %v", err)
	}
	if len(idres) > 0 {
		res := idres[0].(*tabletypes.Transfer)
		return res.Timestamp, nil
	}
	return 0, fmt.Errorf("No Block Data: %v", err)
}
