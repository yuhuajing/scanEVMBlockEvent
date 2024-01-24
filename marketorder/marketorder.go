package marketorder

import (
	"context"
	"encoding/json"
	"fmt"

	"io"
	"log"
	"main/common/config"
	"main/core/database"
	"net/http"
	"nhooyr.io/websocket"
	"nhooyr.io/websocket/wsjson"
	"time"
)

type listbycoll struct {
	Orders []struct {
		Id string `json:"id"`
		//Kind               string `json:"kind"`
		//Side               string `json:"side"`
		Status string `json:"status"`
		//TokenSetId         string `json:"tokenSetId"`
		//TokenSetSchemaHash string `json:"tokenSetSchemaHash"`
		Contract string `json:"contract"`
		//	ContractKind       string `json:"contractKind"`
		Maker string `json:"maker"`
		//Taker              string `json:"taker"`
		//Price              struct {
		//	Currency struct {
		//		Contract string `json:"contract"`
		//		Name     string `json:"name"`
		//		Symbol   string `json:"symbol"`
		//		Decimals int    `json:"decimals"`
		//	} `json:"currency"`
		//	Amount struct {
		//		Raw     string  `json:"raw"`
		//		Decimal float64 `json:"decimal"`
		//		Usd     float64 `json:"usd"`
		//		Native  float64 `json:"native"`
		//	} `json:"amount"`
		//	NetAmount struct {
		//		Raw     string  `json:"raw"`
		//		Decimal float64 `json:"decimal"`
		//		Usd     float64 `json:"usd"`
		//		Native  float64 `json:"native"`
		//	} `json:"netAmount"`
		//} `json:"price"`
		ValidFrom int `json:"validFrom"`
		//ValidUntil int `json:"validUntil"`
		//QuantityFilled    int         `json:"quantityFilled"`
		//QuantityRemaining int         `json:"quantityRemaining"`
		//DynamicPricing    interface{} `json:"dynamicPricing"`
		Criteria struct {
			//Kind string `json:"kind"`
			Data struct {
				Token struct {
					TokenId string `json:"tokenId"`
				} `json:"token"`
			} `json:"data"`
		} `json:"criteria"`
		Source struct {
			//	Id     string `json:"id"`
			Domain string `json:"domain"`
			Name   string `json:"name"`
			//Icon   string `json:"icon"`
			//Url    string `json:"url"`
		} `json:"source"`
		//FeeBps       int           `json:"feeBps"`
		//FeeBreakdown []interface{} `json:"feeBreakdown"`
		Expiration int `json:"expiration"`
		//IsReservoir  interface{}   `json:"isReservoir"`
		//IsDynamic    bool          `json:"isDynamic"`
		//CreatedAt    time.Time     `json:"createdAt"`
		//UpdatedAt    time.Time     `json:"updatedAt"`
		//OriginatedAt time.Time     `json:"originatedAt"`
	} `json:"orders"`
}

func ParseOrderListing() {
	url := "https://api.reservoir.tools/orders/asks/v5"
	contracturl := ""
	for index, contractaddress := range config.Contracts {
		if index == 0 {
			contracturl = fmt.Sprintf("?contracts=%s", contractaddress)
		} else {
			contracturl += fmt.Sprintf("&contracts=%s", contractaddress)
		}
	}
	url += contracturl
	fmt.Println(url)
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("accept", "*/*")
	req.Header.Add("x-api-key", config.MarketToken)
	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	var response listbycoll
	err := json.Unmarshal(body, &response)
	if err != nil {
		fmt.Println("unmarshal orders data error: %v", err)
	}
	for _, order := range response.Orders {
		err = database.AddMarketOrder(order.Id, order.Contract, order.Maker, order.Criteria.Data.Token.TokenId, order.ValidFrom, order.Expiration, order.Status, order.Source.Domain, order.Source.Name)
		if err != nil {
			fmt.Println("database.AddMarketOrder error: %v", err)
		}
	}
}

type MarketEventTags struct {
	Contract string `json:"contract"`
	Source   string `json:"source"`
	Maker    string `json:"maker"`
	Taker    string `json:"taker"`
}

type MarketEventData struct {
	Id                 string `json:"id"`
	Kind               string `json:"kind"`
	Side               string `json:"side"`
	Status             string `json:"status"`
	TokenSetId         string `json:"tokenSetId"`
	TokenSetSchemaHash string `json:"tokenSetSchemaHash"`
	Nonce              int    `json:"nonce"`
	Contract           string `json:"contract"`
	Maker              string `json:"maker"`
	Taker              string `json:"taker"`
	Price              struct {
		Currency struct {
			Contract string `json:"contract"`
			Name     string `json:"name"`
			Symbol   string `json:"symbol"`
			Decimals int    `json:"decimals"`
		} `json:"currency"`
		Amount struct {
			Raw     string  `json:"raw"`
			Decimal float64 `json:"decimal"`
			Usd     float64 `json:"usd"`
			Native  float64 `json:"native"`
		} `json:"amount"`
		NetAmount struct {
			Raw     string  `json:"raw"`
			Decimal float64 `json:"decimal"`
			Usd     float64 `json:"usd"`
			Native  float64 `json:"native"`
		} `json:"netAmount"`
	} `json:"price"`
	ValidFrom         int `json:"validFrom"`
	ValidUntil        int `json:"validUntil"`
	QuantityFilled    int `json:"quantityFilled"`
	QuantityRemaining int `json:"quantityRemaining"`
	Criteria          struct {
		Kind string `json:"kind"`
		Data struct {
			Token struct {
				TokenId string `json:"tokenId"`
				Name    string `json:"name"`
				Image   string `json:"image"`
			} `json:"token"`
			Collection struct {
				Id    string `json:"id"`
				Name  string `json:"name"`
				Image string `json:"image"`
			} `json:"collection"`
		} `json:"data"`
	} `json:"criteria"`
	Source struct {
		Id     string `json:"id"`
		Domain string `json:"domain"`
		Name   string `json:"name"`
		Icon   string `json:"icon"`
		Url    string `json:"url"`
	} `json:"source"`
	FeeBps       int           `json:"feeBps"`
	FeeBreakdown []interface{} `json:"feeBreakdown"`
	Expiration   int           `json:"expiration"`
	IsReservoir  interface{}   `json:"isReservoir"`
	IsDynamic    bool          `json:"isDynamic"`
	CreatedAt    time.Time     `json:"createdAt"`
	UpdatedAt    time.Time     `json:"updatedAt"`
	RawData      struct {
		Kind  string `json:"kind"`
		Salt  string `json:"salt"`
		Zone  string `json:"zone"`
		Offer []struct {
			Token                string `json:"token"`
			ItemType             int    `json:"itemType"`
			EndAmount            string `json:"endAmount"`
			StartAmount          string `json:"startAmount"`
			IdentifierOrCriteria string `json:"identifierOrCriteria"`
		} `json:"offer"`
		Counter       string `json:"counter"`
		EndTime       int    `json:"endTime"`
		Offerer       string `json:"offerer"`
		Partial       bool   `json:"partial"`
		ZoneHash      string `json:"zoneHash"`
		OrderType     int    `json:"orderType"`
		StartTime     int    `json:"startTime"`
		ConduitKey    string `json:"conduitKey"`
		Consideration []struct {
			Token                string `json:"token"`
			ItemType             int    `json:"itemType"`
			EndAmount            string `json:"endAmount"`
			Recipient            string `json:"recipient"`
			StartAmount          string `json:"startAmount"`
			IdentifierOrCriteria string `json:"identifierOrCriteria"`
		} `json:"consideration"`
	} `json:"rawData"`
}

func SubOrder() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
	defer cancel()

	// sepolia: wss://ws.ws-sepolia.reservoir.tools
	url := fmt.Sprintf("wss://ws.reservoir.tools?api_key=%s", config.MarketToken)
	conn, _, err := websocket.Dial(ctx, url, &websocket.DialOptions{
		HTTPHeader: http.Header{
			"Authorization": []string{fmt.Sprintf("Bearer %s", config.MarketToken)},
		},
	})
	if err != nil {
		log.Fatal("Error connecting to Reservoir: ", err)
	}
	defer conn.Close(websocket.StatusInternalError, "the sky is falling")

	fmt.Println("Connected to Reservoir")

	message := make(map[string]interface{})
	for {
		err := wsjson.Read(ctx, conn, &message)
		if err != nil {
			err := wsjson.Read(ctx, conn, &message)
			if err != nil {
				log.Println("Error reading message: ", err)
				break
			}
		}

		fmt.Println("Message received: ", message)
		if message["tags"] != nil {
			var response MarketEventTags
			err = json.Unmarshal([]byte(fmt.Sprintf("%s", message["tags"])), &response)
			if err != nil {
				fmt.Println("unmarshal orders data error: %v", err)
			}
			fmt.Println(response)
		}
		// When the connection is ready, subscribe to the top-bids event
		if message["status"] == "ready" {
			fmt.Println("Subscribing")
			err = wsjson.Write(ctx, conn, map[string]interface{}{
				"type":  "subscribe",
				"event": "top-bid.changed",
			})
			if err != nil {
				log.Println("Error subscribing: ", err)
				break
			}

			// To unsubscribe, send the following message
			// err = wsjson.Write(ctx, conn, map[string]interface{}{
			//     "type":    "unsubscribe",
			//     "event": "top-bid.changed",
			// })
			// if err != nil {
			//     log.Println("Error unsubscribing: ", err)
			//     break
			// }
		}
	}
}
