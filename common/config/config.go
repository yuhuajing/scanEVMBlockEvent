package config

import (
	"go.mongodb.org/mongo-driver/mongo"
	"main/common/ethconn"
	"strings"
)

var (
	EthServer = "wss://eth-mainnet.g.alchemy.com/v2/cnkFUCDI1vY0c2Xg8aUmSjaK73tIveHy" //"https://eth-mainnet.g.alchemy.com/v2/cnkFUCDI1vY0c2Xg8aUmSjaK73tIveHy"
	Client    = ethconn.ConnBlockchain(EthServer)
	//Ethereum Mainnet
	Contracts      = []string{strings.ToLower("0x1aae1A668c92Eb411eAfD80DD0c60ca67ad17a1c"), strings.ToLower("0xff2B4721F997c242fF406a626f17df083Bd2C568")}
	ContractSupply = map[string]int{
		strings.ToLower("0x1aae1A668c92Eb411eAfD80DD0c60ca67ad17a1c"): 1155,
		strings.ToLower("0xff2B4721F997c242fF406a626f17df083Bd2C568"): 515,
	}
	Collections      = []string{strings.ToLower("efesspacenation"), strings.ToLower("alphagatespacenation")}
	StartBlockHeight = 19039000
)

var BlockWithTimestamp = make(map[uint64]uint64)
var NftOwners = make(map[string]map[string]bool)
var Topic = map[string]int{ // index starts from 1
	"0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925": 1, //Approval
	"0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef": 2, // Transfer
	"0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31": 3} // ApprovalForAll

var (
	Mongodburl              = "mongodb://clay:password@127.0.0.1:27017"
	Dbname                  = "holdnft"
	DbcollectionTrans       = "transfer"
	DbcollectionApproval    = "approval"
	DbcollectionApproForAll = "approvalall"
	DbcollectionOwner       = "owner"
	DbcollectionSB          = "startblock"
	DbcollectionOpensea     = "openseaorders"
	Mongoclient             *mongo.Client
)
