package database

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/gofiber/fiber/v2/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"main/common/config"
	"main/common/dbconn"
	"main/common/tabletypes"
	"main/core/blocktime"
	"reflect"
	"strconv"
	"strings"
)

func Insert(logdata types.Log) {
	topic := logdata.Topics[0].Hex()
	topicTable := config.Topic[topic]
	timestamp := config.BlockWithTimestamp[logdata.BlockNumber]
	if topicTable == 1 {
		err := InsertApprovalDB(topic, timestamp, logdata) //Approval
		if err != nil {
			log.Fatalf("err in InsertTransDB: %v", err)
		}
	} else if topicTable == 2 {
		err := InsertTransDB(topic, timestamp, logdata) //Transfer
		if err != nil {
			log.Fatalf("err in InsertTransDB: %v", err)
		}
	} else if topicTable == 3 {
		err := InsertApprovalAllDB(topic, timestamp, logdata) //ApprovalForAll
		if err != nil {
			log.Fatalf("err in InsertTransDB: %v", err)
		}
	}
}

func ModifyOwner(address string, id int, owner string, blockNumber, timestamp uint64, logIndex uint) error {
	filter := bson.M{"tokenid": id, "address": strings.ToLower(address)}
	err, idres := GetDocuments(config.DbcollectionOwner, filter, &tabletypes.Owner{})
	if err != nil {
		return fmt.Errorf("InsertTransDB:err in getting Trans data: %v", err)
	}
	if len(idres) == 0 {
		var res = tabletypes.Owner{
			Id:          utils.UUIDv4(),
			Blocknumber: blockNumber,
			Timestamp:   timestamp,
			Address:     strings.ToLower(address),
			Tokenid:     id,
			Owner:       strings.ToLower(owner),
			Logindex:    logIndex,
		}
		err := InsertDocument(config.DbcollectionOwner, res)
		if err != nil {
			return fmt.Errorf("InsertNFTdataDB:err in inserting NFTData")
		}
		return nil
	} else {
		res := idres[0].(*tabletypes.Owner)
		if int(res.Blocknumber) < int(blockNumber) || int(res.Blocknumber) == int(blockNumber) && int(res.Logindex) < int(logIndex) {
			filter := bson.M{"address": strings.ToLower(address), "tokenid": id}
			update := bson.M{"$set": bson.M{"owner": strings.ToLower(owner), "blocknumber": blockNumber, "timestamp": timestamp}}
			err := UpdateDocument(config.DbcollectionOwner, filter, update)
			if err != nil {
				return fmt.Errorf("InsertNFTdataDB:err in inserting NFTData")
			}
		}
	}
	return nil
}

func UpdateOwnerTimestamp() bool {
	err, idres := GetDocuments(config.DbcollectionOwner, bson.M{}, &tabletypes.Owner{})
	if err != nil {
		return false
	}
	totalSupply := 0
	for _, contract := range config.Contracts {
		totalSupply += config.ContractSupply[strings.ToLower(contract)]
	}
	if len(idres) == totalSupply {
		for _, owner := range idres {
			res := owner.(*tabletypes.Owner)
			if res.Timestamp == 0 {
				timestamp := config.BlockWithTimestamp[res.Blocknumber]
				if timestamp == 0 {
					timestamp = blocktime.GetBlockTime(res.Blocknumber)
					config.BlockWithTimestamp[res.Blocknumber] = timestamp
				}
				filter := bson.M{"address": strings.ToLower(res.Address), "owner": res.Owner, "tokenid": res.Tokenid}
				update := bson.M{"$set": bson.M{"timestamp": timestamp}}
				err := UpdateDocument(config.DbcollectionOwner, filter, update)
				if err != nil {
					return false
				}
			}
		}
		return true
	}
	return false
}

func CreOrUpdateStartBlock(contract string, blocknumber uint64) error {
	filter := bson.M{"address": strings.ToLower(contract)}
	err, idres := GetDocuments(config.DbcollectionSB, filter, &tabletypes.Startblocks{})
	if err != nil {
		return fmt.Errorf("InsertTransDB:err in getting Trans data: %v", err)
	}
	if len(idres) == 0 {
		var res = tabletypes.Startblocks{
			Id:                 utils.UUIDv4(),
			Historyblocknumber: blocknumber,
			//Newblocknumber:     blockNumber,
			Address: strings.ToLower(contract),
		}
		err := InsertDocument(config.DbcollectionSB, res)
		if err != nil {
			return fmt.Errorf("CreOrUpdateStartBlock:err in inserting")
		}

	} else {
		res := idres[0].(*tabletypes.Startblocks)
		if res.Historyblocknumber < blocknumber {
			update := bson.M{"$set": bson.M{"historyblocknumber": blocknumber}}
			err := UpdateDocument(config.DbcollectionSB, filter, update)
			if err != nil {
				return fmt.Errorf("CreOrUpdateStartBlock:err in updating")
			}
		}
	}
	return nil
}

func GetStartBlockNumber(contract string) uint64 {
	filter := bson.M{"address": strings.ToLower(contract)}
	_, idres := GetDocuments(config.DbcollectionSB, filter, &tabletypes.Startblocks{})
	if len(idres) != 0 {
		res := idres[0].(*tabletypes.Startblocks)
		return res.Historyblocknumber
	}
	return config.ContractDeployHeight[strings.ToLower(contract)]
}

func InsertTransDB(topic string, timestamp uint64, logdata types.Log) error {
	address := strings.ToLower(logdata.Address.Hex())
	txhash := strings.ToLower(logdata.TxHash.Hex())
	from := strings.ToLower("0x" + logdata.Topics[1].Hex()[26:])
	to := strings.ToLower("0x" + logdata.Topics[2].Hex()[26:])
	tokenIDInt, _ := strconv.ParseInt(logdata.Topics[3].Hex()[2:], 16, 64)
	filter := bson.M{"txhash": txhash, "logindex": logdata.Index, "address": address}
	err, idres := GetDocuments(config.DbcollectionTrans, filter, &tabletypes.Transfer{})
	if err != nil {
		return fmt.Errorf("InsertTransDB:err in getting Trans data: %v", err)
	}
	if len(idres) == 0 {
		var res = tabletypes.Transfer{
			Id:          utils.UUIDv4(),
			Blocknumber: logdata.BlockNumber,
			Timestamp:   timestamp,
			Address:     address,
			Func:        topic,
			From:        from,
			Operator:    to,
			Tokenid:     tokenIDInt,
			Txhash:      txhash,
			Logindex:    logdata.Index,
		}
		err := InsertDocument(config.DbcollectionTrans, res)
		if err != nil {
			return fmt.Errorf("InsertTransDB:err in getting Trans data: %v", err)
		}
	} else {
		res := idres[0].(*tabletypes.Transfer)
		timestamp = res.Timestamp
	}
	err = ModifyOwner(address, int(tokenIDInt), to, logdata.BlockNumber, timestamp, logdata.Index)
	if err != nil {
		return fmt.Errorf("ModifyOwner:err %v", err)
	}

	//err = UpdateStartBlock(address, logdata.BlockNumber)
	//if err != nil {
	//	return fmt.Errorf("CreOrUpdateStartBlock:err %v", err)
	//}
	return nil
}
func InsertApprovalDB(topic string, timestamp uint64, logdata types.Log) error {
	address := strings.ToLower(logdata.Address.Hex())
	txhash := strings.ToLower(logdata.TxHash.Hex())
	from := strings.ToLower("0x" + logdata.Topics[1].Hex()[26:])
	to := strings.ToLower("0x" + logdata.Topics[2].Hex()[26:])

	tokenIDInt, _ := strconv.ParseInt(logdata.Topics[3].Hex()[2:], 16, 64)
	filter := bson.M{"txhash": txhash, "logindex": logdata.Index, "address": address}
	err, idres := GetDocuments(config.DbcollectionApproval, filter, &tabletypes.Approval{})
	if err != nil {
		return fmt.Errorf("InsertApprovalDB:err in getting Trans data: %v", err)
	}
	if len(idres) == 0 {
		var res = tabletypes.Approval{
			Id:          utils.UUIDv4(),
			Blocknumber: logdata.BlockNumber,
			Timestamp:   timestamp,
			Address:     address,
			Func:        topic,
			From:        from,
			Operator:    to,
			Tokenid:     tokenIDInt,
			Txhash:      txhash,
			Logindex:    logdata.Index,
		}
		err := InsertDocument(config.DbcollectionApproval, res)
		if err != nil {
			return fmt.Errorf("InsertApprovalDB:err in inserting NFTApproval")
		}
	}
	//err = UpdateStartBlock(address, logdata.BlockNumber)
	//if err != nil {
	//	return fmt.Errorf("CreOrUpdateStartBlock:err %v", err)
	//}
	return nil
}
func InsertApprovalAllDB(topic string, timestamp uint64, logdata types.Log) error {
	address := strings.ToLower(logdata.Address.Hex())
	txhash := strings.ToLower(logdata.TxHash.Hex())
	from := strings.ToLower("0x" + logdata.Topics[1].Hex()[26:])
	to := strings.ToLower("0x" + logdata.Topics[2].Hex()[26:])
	filter := bson.M{"txhash": txhash, "logindex": logdata.Index, "address": address}
	err, idres := GetDocuments(config.DbcollectionApproForAll, filter, &tabletypes.Transfer{})
	if err != nil {
		return fmt.Errorf("InsertTransDB:err in getting Trans data: %v", err)
	}
	if len(idres) == 0 {
		var res = tabletypes.ApprovalForAll{
			Id:          utils.UUIDv4(),
			Blocknumber: logdata.BlockNumber,
			Timestamp:   timestamp,
			Address:     address,
			Func:        topic,
			From:        from,
			Operator:    to,
			Txhash:      txhash,
			Logindex:    logdata.Index,
		}
		err := InsertDocument(config.DbcollectionApproForAll, res)
		if err != nil {
			return fmt.Errorf("InsertNFTdataDB:err in inserting NFTData")
		}
	}
	//err = UpdateStartBlock(address, logdata.BlockNumber)
	//if err != nil {
	//	return fmt.Errorf("CreOrUpdateStartBlock:err %v", err)
	//}
	return nil
}

func InsertDocument(collectionname string, data interface{}) error {
	collection := dbconn.GetCollection(collectionname)
	_, err := collection.InsertOne(context.Background(), data)
	if err != nil {
		return fmt.Errorf("failed to insert document: %v", err)
	}
	return nil
}

func GetDocuments(collectionname string, filter bson.M, result interface{}, opts ...*options.FindOptions) (error, []interface{}) {
	collection := dbconn.GetCollection(collectionname)
	cur, err := collection.Find(context.Background(), filter, opts...)
	if err != nil {
		return fmt.Errorf(fmt.Sprintf("failed to get documents: %v", err)), nil
	}
	defer cur.Close(context.Background())

	var results []interface{}
	for cur.Next(context.Background()) {
		// Create a new instance of the result type for each document
		res := reflect.New(reflect.TypeOf(result).Elem()).Interface()
		err := cur.Decode(res)
		if err != nil {
			return fmt.Errorf(fmt.Sprintf("failed to decode document: %v", err)), nil
		}
		results = append(results, res)
	}

	if err := cur.Err(); err != nil {
		return fmt.Errorf(fmt.Sprintf("cursor error: %v", err)), nil
	}

	return nil, results
}
func UpdateDocument(collectionname string, filter bson.M, update bson.M) error {
	//	fmt.Println("up")
	collection := dbconn.GetCollection(collectionname)
	_, err := collection.UpdateMany(context.TODO(), filter, update)
	if err != nil {
		return fmt.Errorf("failed to update document: %v", err)
	}
	return nil
}

func GetOwnerByNFTId(contractaddress string, id int) ([]int, string, error) {
	owner := ""
	allIds := make([]int, 0)
	filter := bson.M{"address": strings.ToLower(contractaddress), "tokenid": id}
	err, idres := GetDocuments(config.DbcollectionOwner, filter, &tabletypes.Owner{})
	if err != nil {
		return allIds, owner, err
	}
	if len(idres) > 0 {
		res := idres[0].(*tabletypes.Owner)
		owner = res.Owner
		if !config.NftOwners[strings.ToLower(contractaddress)][strings.ToLower(owner)] {
			allIds, err = getAllIdByOwner(contractaddress, owner)
			if err != nil {
				return allIds, owner, err
			}
		}
	}
	return allIds, owner, nil
}

func getAllIdByOwner(contractaddress, owner string) ([]int, error) {
	ownerids := make([]int, 0)
	filter := bson.M{"address": strings.ToLower(contractaddress), "owner": strings.ToLower(owner)}
	err, idres := GetDocuments(config.DbcollectionOwner, filter, &tabletypes.Owner{})
	if err != nil {
		return ownerids, err
	}
	if len(idres) > 0 {
		for _, v := range idres {
			res := v.(*tabletypes.Owner)
			ownerids = append(ownerids, res.Tokenid)
		}
		config.NftOwners[strings.ToLower(contractaddress)][strings.ToLower(owner)] = true
	}
	return ownerids, nil
}

func AddMarketOrder(id, address, owner string, tokenId string, listTime int, expirationtime int, status, domain, name string) error {
	tokenIDInt, _ := strconv.ParseInt(tokenId, 10, 64)
	filter := bson.M{"id": strings.ToLower(id)}
	err, idres := GetDocuments(config.DbcollectionMarket, filter, &tabletypes.MarketOrder{})
	if err != nil {
		return fmt.Errorf("AddMarketOrder:err in getting market data: %v", err)
	}
	if len(idres) == 0 {
		var res = tabletypes.MarketOrder{
			Id:             id,
			Listingtime:    listTime,
			Expirationtime: expirationtime,
			Address:        strings.ToLower(address),
			Tokenid:        int(tokenIDInt),
			Owner:          strings.ToLower(owner),
			Status:         status,
			Domain:         domain,
			Name:           name,
		}
		err := InsertDocument(config.DbcollectionMarket, res)
		if err != nil {
			return fmt.Errorf("AddMarketOrder:err in inserting marketData")
		}
	}
	return nil
}

func AddOpenSeaOrder(orderhash, address, owner string, id string, listTime int, expirationtime int) error {
	tokenIDInt, _ := strconv.ParseInt(id, 10, 64)
	filter := bson.M{"orderhash": strings.ToLower(orderhash)}
	err, idres := GetDocuments(config.DbcollectionOpensea, filter, &tabletypes.OpenseaOrder{})
	if err != nil {
		return fmt.Errorf("AddOpenSeaOrder:err in getting opensea data: %v", err)
	}
	if len(idres) == 0 {
		var res = tabletypes.OpenseaOrder{
			Id:             utils.UUIDv4(),
			Listingtime:    listTime,
			Expirationtime: expirationtime,
			Orderhash:      strings.ToLower(orderhash),
			Address:        strings.ToLower(address),
			Tokenid:        int(tokenIDInt),
			Owner:          strings.ToLower(owner),
			Status:         tabletypes.StatusListing,
		}
		err := InsertDocument(config.DbcollectionOpensea, res)
		if err != nil {
			return fmt.Errorf("AddOpenSeaOrder:err in inserting openseaData")
		}
	}
	//else {
	//	res := idres[0].(*tabletypes.OpenseaOrder)
	//	if res.Listingtime < listTime {
	//		update := bson.M{"$set": bson.M{"listingtime": listTime, "expirationtime": expirationtime, "orderhash": orderhash}}
	//		err := UpdateDocument(config.DbcollectionOpensea, filter, update)
	//		if err != nil {
	//			return fmt.Errorf("UpdateOpenSeaOrder:err in updating openseaData")
	//		}
	//	}
	//}
	return nil
}

func CancelOpenSeaOrder(orderhash string) error {
	filter := bson.M{"orderhash": strings.ToLower(orderhash)}
	err, idres := GetDocuments(config.DbcollectionOpensea, filter, &tabletypes.OpenseaOrder{})
	if err != nil {
		return fmt.Errorf("CancelOpenSeaOrder:err in getting opensea data: %v", err)
	}
	if len(idres) != 0 {
		update := bson.M{"$set": bson.M{"status": tabletypes.StatusCancel}}
		err := UpdateDocument(config.DbcollectionOpensea, filter, update)
		if err != nil {
			return fmt.Errorf("UpdateOpenSeaOrder:err in updating openseaData")
		}
	}
	return nil
}

//func UpdateLevel(collectionname string, tokenID int, level string) error {
//	update := bson.M{"$set": bson.M{"level": level}}
//	fmt.Println(level)
//	itopkenID := strconv.FormatInt(int64(tokenID), 10)
//	fmt.Println(itopkenID)
//	filter := bson.M{"tokenID": itopkenID}
//	err := UpdateDocument(collectionname, filter, update)
//	if err != nil {
//		return fmt.Errorf("err in updating order's tx status: %v", err)
//	}
//	return nil
//}

//func GetLatestBlockNumber() uint64 {
//	_, idres := GetDocuments(config.DbcollectionSB, bson.M{}, &tabletypes.Startblocks{})
//	if len(idres) != 0 {
//		res := idres[0].(*tabletypes.Startblocks)
//		return res.Newblocknumber
//	}
//	return 0
//}

//func UpdateStartBlock(address string, blockNumber uint64) error {
//
//	if len(idres) != 0 {
//		res := idres[0].(*tabletypes.Startblocks)
//		if int(res.Historyblocknumber) < int(blockNumber) {
//			log.Printf("UpdateStartBlock from: %d to: %d\n", res.Historyblocknumber, blockNumber)
//			filter := bson.M{"address": strings.ToLower(address)}
//			update := bson.M{"$set": bson.M{"blocknumber": blockNumber, "historyblocknumber": blockNumber}}
//			err := UpdateDocument(config.DbcollectionSB, filter, update)
//			if err != nil {
//				return fmt.Errorf("CreOrUpdateStartBlock:err in inserting")
//			}
//		}
//	}
//	return nil
//}

//func CreOrUpdateLatestBlock(blockNumber uint64) error {
//	err, idres := GetDocuments(config.DbcollectionSB, bson.M{}, &tabletypes.Startblocks{})
//	if err != nil {
//		return fmt.Errorf("InsertTransDB:err in getting Trans data: %v", err)
//	}
//
//	if len(idres) == 0 {
//		for _, contractaddress := range config.Contracts {
//			var res = tabletypes.Startblocks{
//				Id:                 utils.UUIDv4(),
//				Historyblocknumber: config.ContractDeployHeight[strings.ToLower(contractaddress)],
//			//	Newblocknumber:     blockNumber,
//				Address:            strings.ToLower(contractaddress),
//			}
//			err := InsertDocument(config.DbcollectionSB, res)
//			if err != nil {
//				return fmt.Errorf("CreOrUpdateLatestBlock:err in inserting")
//			}
//		}
//	} else {
//		res := idres[0].(*tabletypes.Startblocks)
//		if int(res.Newblocknumber) < int(blockNumber) {
//			update := bson.M{"$set": bson.M{"newblocknumber": blockNumber}}
//			err := UpdateDocument(config.DbcollectionSB, bson.M{}, update)
//			if err != nil {
//				return fmt.Errorf("CreOrUpdateLatestBlock:err in updating")
//			}
//		}
//	}
//	return nil
//}
