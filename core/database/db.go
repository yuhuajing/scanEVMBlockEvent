package database

import (
	"fmt"
	"main/common/config"
	"main/common/tabletypes"

	"strconv"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/jinzhu/gorm"
)

// func Insert(dba *gorm.DB, logdata *types.Log) {
func Insert(dba *gorm.DB, logdata types.Log) {
	topic := logdata.Topics[0].Hex()
	topicTable := config.Topic[topic]
	tokenIDInt := int64(0)
	approved := false
	firstAddr := ""
	secondAddr := ""
	if topicTable == 1 || topicTable == 2 || topicTable == 3 {
		firstAddr = "0x" + logdata.Topics[1].Hex()[26:]
		secondAddr = "0x" + logdata.Topics[2].Hex()[26:]
		if len(logdata.Topics) >= 4 {
			tokenIDInt, _ = strconv.ParseInt(logdata.Topics[3].Hex()[2:], 16, 64)
		} else {
			if logdata.Data[31:][0] == 1 {
				approved = true
			}
		}
	}

	if topicTable == 1 { //Approval
		res := dba.Model(&tabletypes.Approval{}).Where("txhash = ? AND logindex = ?", logdata.TxHash.Hex(), int(logdata.Index)).First(&tabletypes.Approval{})
		if res.RowsAffected == 0 {
			dba.Create(&tabletypes.Approval{
				Blocknumber: logdata.BlockNumber,
				Timestamp:   config.BlockWithTimestamp[fmt.Sprint(logdata.BlockNumber)],
				Address:     logdata.Address.Hex(),
				Func:        topic,
				Owner:       firstAddr,
				Approved:    secondAddr,
				Tokenid:     tokenIDInt,
				Txhash:      logdata.TxHash.Hex(),
				Logindex:    logdata.Index,
			})
		}

	} else if topicTable == 2 { //Transfer
		res := dba.Model(&tabletypes.Transfer{}).Where("txhash = ? AND logindex = ?", logdata.TxHash.Hex(), int(logdata.Index)).First(&tabletypes.Transfer{})
		if res.RowsAffected == 0 {
			dba.Create(&tabletypes.Transfer{
				Blocknumber: logdata.BlockNumber,
				Timestamp:   config.BlockWithTimestamp[fmt.Sprint(logdata.BlockNumber)],
				Address:     logdata.Address.Hex(),
				Func:        topic,
				From:        firstAddr,
				To:          secondAddr,
				Tokenid:     tokenIDInt,
				Txhash:      logdata.TxHash.Hex(),
				Logindex:    logdata.Index,
			})
			ModifyOwner(dba, logdata.Address.Hex(), int(tokenIDInt), secondAddr)
		}
	} else if topicTable == 3 { //ApprovalForAll
		res := dba.Model(&tabletypes.ApprovalForAll{}).Where("txhash = ? AND logindex = ?", logdata.TxHash.Hex(), int(logdata.Index)).First(&tabletypes.ApprovalForAll{})
		if res.RowsAffected == 0 {
			dba.Create(&tabletypes.ApprovalForAll{
				Blocknumber: logdata.BlockNumber,
				Timestamp:   config.BlockWithTimestamp[fmt.Sprint(logdata.BlockNumber)],
				Address:     logdata.Address.Hex(),
				Func:        topic,
				Owner:       firstAddr,
				Operator:    secondAddr,
				Approved:    approved,
				Txhash:      logdata.TxHash.Hex(),
				Logindex:    logdata.Index,
			})
		}
	}
}

func MakeOwner(db *gorm.DB, address string) {
	if !HasOwner(db, address) {
		for i := 0; i < 515; i++ {
			db.Create(&tabletypes.Owner{
				Address: address,
				Tokenid: int64(i),
			})
		}
	}
	for i := 0; i < 515; i++ {
		parseTransfer(db, address, i)
	}
}

func HasOwner(db *gorm.DB, address string) (hasOwner bool) {
	res := db.Model(&tabletypes.Owner{}).Where("address = ? AND tokenid = ?", address, 514).First(&tabletypes.Owner{})
	if res.RowsAffected == 0 {
		hasOwner = false
	} else {
		hasOwner = true
	}
	return
}

func parseTransfer(db *gorm.DB, address string, tokenid int) {
	res := []tabletypes.Transfer{}
	db.Model(&tabletypes.Transfer{}).Where("address = ? AND tokenid = ?", address, tokenid).Order("blocknumber desc").Order("txindex desc").Limit(1).Find(&res)
	if len(res) == 0 {
		return
	}
	ModifyOwner(db, res[0].Address, tokenid, res[0].To)
}

func ModifyOwner(db *gorm.DB, address string, id int, owner string) {
	res := db.Model(&tabletypes.Owner{}).Where("address = ? AND tokenid = ?", address, id).First(&tabletypes.Owner{})
	if res.RowsAffected == 0 {
		db.Create(&tabletypes.Owner{
			Address: address,
			Owner:   owner,
			Tokenid: int64(id),
		})
	} else {
		db.Model(&tabletypes.Owner{}).Where("address = ? AND tokenid = ?", address, id).Update("owner", owner)
	}
}
