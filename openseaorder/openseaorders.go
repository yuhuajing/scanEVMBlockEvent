package openseaorder

import (
	"encoding/json"
	"fmt"
	"github.com/foundVanting/opensea-stream-go/entity"
	"github.com/foundVanting/opensea-stream-go/opensea"
	openseatypes "github.com/foundVanting/opensea-stream-go/types"
	"github.com/mitchellh/mapstructure"
	"github.com/xiaowang7777/phx"
	"io"
	"main/core/database"
	"math"
	"net/http"
)

type Response struct {
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
	for idLen > 0 && index < idLen {
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
		req, _ := http.NewRequest("GET", url, nil)
		req.Header.Add("accept", "application/json")
		req.Header.Add("x-api-key", "9602c2e9de24426196b5c317099155c7")
		res, _ := http.DefaultClient.Do(req)
		defer res.Body.Close()
		body, _ := io.ReadAll(res.Body)
		var response Response
		err := json.Unmarshal(body, &response)
		if err != nil {
			fmt.Println("Error:", err)
		}
		for _, order := range response.Orders {
			err = database.AddOpenSeaOrder(order.OrderHash, order.ProtocolData.Parameters.Offer[0].Token, order.ProtocolData.Parameters.Offerer, order.ProtocolData.Parameters.Offer[0].IdentifierOrCriteria, order.ListingTime, order.ExpirationTime)
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
