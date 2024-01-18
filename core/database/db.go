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
	owner := "0x" + logdata.Topics[1].Hex()[26:]
	operator := "0x" + logdata.Topics[2].Hex()[26:]
	txhash := logdata.TxHash.Hex()
	logindex := int(logdata.Index)
	address := logdata.Address.Hex()
	timestamp := config.BlockWithTimestamp[logdata.BlockNumber]
	if timestamp == 0 {
		timestamp = blocktime.GetBlockTime(logdata.BlockNumber)
	}
	if topicTable == 1 {
		err := InsertApprovalDB(topic, txhash, address, owner, operator, logindex, timestamp, logdata) //Approval
		if err != nil {
			log.Fatalf("err in InsertTransDB: %v", err)
		}
	} else if topicTable == 2 {
		err := InsertTransDB(topic, txhash, address, owner, operator, logindex, timestamp, logdata) //Transfer
		if err != nil {
			log.Fatalf("err in InsertTransDB: %v", err)
		}
	} else if topicTable == 3 {
		err := InsertApprovalAllDB(topic, txhash, address, owner, operator, logindex, timestamp, logdata) //ApprovalForAll
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
			update := bson.M{"$set": bson.M{"owner": strings.ToLower(owner)}}
			err := UpdateDocument(config.DbcollectionOwner, filter, update)
			if err != nil {
				return fmt.Errorf("InsertNFTdataDB:err in inserting NFTData")
			}
		}
	}
	return nil
}

func InsertTransDB(topic, txhash, address, from, to string, logindex int, timestamp uint64, logdata types.Log) error {
	tokenIDInt, _ := strconv.ParseInt(logdata.Topics[3].Hex()[2:], 16, 64)
	filter := bson.M{"txhash": strings.ToLower(txhash), "logindex": logindex, "address": strings.ToLower(address)}
	err, idres := GetDocuments(config.DbcollectionTrans, filter, &tabletypes.Transfer{})
	if err != nil {
		return fmt.Errorf("InsertTransDB:err in getting Trans data: %v", err)
	}
	if len(idres) == 0 {
		var res = tabletypes.Transfer{
			Id:          utils.UUIDv4(),
			Blocknumber: logdata.BlockNumber,
			Timestamp:   timestamp,
			Address:     strings.ToLower(logdata.Address.Hex()),
			Func:        topic,
			From:        strings.ToLower(from),
			Operator:    strings.ToLower(to),
			Tokenid:     tokenIDInt,
			Txhash:      strings.ToLower(logdata.TxHash.Hex()),
			Logindex:    logdata.Index,
		}
		err := InsertDocument(config.DbcollectionTrans, res)
		if err != nil {
			return fmt.Errorf("InsertTransDB:err in getting Trans data: %v", err)
		}
	}

	err = ModifyOwner(logdata.Address.Hex(), int(tokenIDInt), to, logdata.BlockNumber, logdata.Index)
	if err != nil {
		return fmt.Errorf("ModifyOwner:err %v", err)
	}
	return nil
}
func InsertApprovalDB(topic, txhash, address, from, to string, logindex int, timestamp uint64, logdata types.Log) error {
	tokenIDInt, _ := strconv.ParseInt(logdata.Topics[3].Hex()[2:], 16, 64)
	filter := bson.M{"txhash": strings.ToLower(txhash), "logindex": logindex, "address": strings.ToLower(address)}
	err, idres := GetDocuments(config.DbcollectionApproval, filter, &tabletypes.Approval{})
	if err != nil {
		return fmt.Errorf("InsertApprovalDB:err in getting Trans data: %v", err)
	}
	if len(idres) == 0 {
		var res = tabletypes.Approval{
			Id:          utils.UUIDv4(),
			Blocknumber: logdata.BlockNumber,
			Timestamp:   timestamp,
			Address:     strings.ToLower(logdata.Address.Hex()),
			Func:        topic,
			From:        strings.ToLower(from),
			Operator:    strings.ToLower(to),
			Tokenid:     tokenIDInt,
			Txhash:      strings.ToLower(logdata.TxHash.Hex()),
			Logindex:    logdata.Index,
		}
		err := InsertDocument(config.DbcollectionApproval, res)
		if err != nil {
			return fmt.Errorf("InsertApprovalDB:err in inserting NFTApproval")
		}
		return nil
	}
	return nil
}
func InsertApprovalAllDB(topic, txhash, address, from, to string, logindex int, timestamp uint64, logdata types.Log) error {
	filter := bson.M{"txhash": strings.ToLower(txhash), "logindex": logindex, "address": strings.ToLower(address)}
	err, idres := GetDocuments(config.DbcollectionApproForAll, filter, &tabletypes.Transfer{})
	if err != nil {
		return fmt.Errorf("InsertTransDB:err in getting Trans data: %v", err)
	}
	if len(idres) == 0 {
		var res = tabletypes.ApprovalForAll{
			Id:          utils.UUIDv4(),
			Blocknumber: logdata.BlockNumber,
			Timestamp:   timestamp,
			Address:     strings.ToLower(logdata.Address.Hex()),
			Func:        topic,
			From:        strings.ToLower(from),
			Operator:    strings.ToLower(to),
			Txhash:      strings.ToLower(logdata.TxHash.Hex()),
			Logindex:    logdata.Index,
		}
		err := InsertDocument(config.DbcollectionApproForAll, res)
		if err != nil {
			return fmt.Errorf("InsertNFTdataDB:err in inserting NFTData")
		}
		return nil
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
