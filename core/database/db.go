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
	if timestamp == 0 {
		timestamp = blocktime.GetBlockTime(logdata.BlockNumber)
	}
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

func ModifyOwner(address string, id int, owner string, blockNumber uint64, logIndex uint) error {
	filter := bson.M{"tokenid": id, "address": strings.ToLower(address)}
	err, idres := GetDocuments(config.DbcollectionOwner, filter, &tabletypes.Owner{})
	if err != nil {
		return fmt.Errorf("InsertTransDB:err in getting Trans data: %v", err)
	}
	if len(idres) == 0 {
		var res = tabletypes.Owner{
			Id:          utils.UUIDv4(),
			Blocknumber: blockNumber,
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
			update := bson.M{"$set": bson.M{"owner": strings.ToLower(owner), "blocknumber": blockNumber}}
			err := UpdateDocument(config.DbcollectionOwner, filter, update)
			if err != nil {
				return fmt.Errorf("InsertNFTdataDB:err in inserting NFTData")
			}
		}
	}
	return nil
}

func CreOrUpdateStartBlock(address string, blockNumber uint64) error {
	filter := bson.M{"address": strings.ToLower(address)}
	err, idres := GetDocuments(config.DbcollectionSB, filter, &tabletypes.Startblocks{})
	if err != nil {
		return fmt.Errorf("InsertTransDB:err in getting Trans data: %v", err)
	}
	if len(idres) == 0 {
		var res = tabletypes.Startblocks{
			Id:          utils.UUIDv4(),
			Blocknumber: blockNumber,
			Address:     strings.ToLower(address),
		}
		err := InsertDocument(config.DbcollectionSB, res)
		if err != nil {
			return fmt.Errorf("CreOrUpdateStartBlock:err in inserting")
		}
		return nil
	} else {
		res := idres[0].(*tabletypes.Startblocks)
		if int(res.Blocknumber) < int(blockNumber) {
			filter := bson.M{"address": strings.ToLower(address)}
			update := bson.M{"$set": bson.M{"blocknumber": blockNumber}}
			err := UpdateDocument(config.DbcollectionSB, filter, update)
			if err != nil {
				return fmt.Errorf("CreOrUpdateStartBlock:err in inserting")
			}
		}
	}
	return nil
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
	}

	err = ModifyOwner(address, int(tokenIDInt), to, logdata.BlockNumber, logdata.Index)
	if err != nil {
		return fmt.Errorf("ModifyOwner:err %v", err)
	}

	err = CreOrUpdateStartBlock(address, logdata.BlockNumber)
	if err != nil {
		return fmt.Errorf("CreOrUpdateStartBlock:err %v", err)
	}
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
	err = CreOrUpdateStartBlock(address, logdata.BlockNumber)
	if err != nil {
		return fmt.Errorf("CreOrUpdateStartBlock:err %v", err)
	}
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
	err = CreOrUpdateStartBlock(address, logdata.BlockNumber)
	if err != nil {
		return fmt.Errorf("CreOrUpdateStartBlock:err %v", err)
	}
	return nil
}

func UpdateLevel(collectionname string, tokenID int, level string) error {
	update := bson.M{"$set": bson.M{"level": level}}
	fmt.Println(level)
	itopkenID := strconv.FormatInt(int64(tokenID), 10)
	fmt.Println(itopkenID)
	filter := bson.M{"tokenID": itopkenID}
	err := UpdateDocument(collectionname, filter, update)
	if err != nil {
		return fmt.Errorf("err in updating order's tx status: %v", err)
	}
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
