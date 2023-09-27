package config

import (
	"strings"
)

var (
	EthServer = "wss://eth.getblock.io/ab0b1aa0-b490-4dc0-9bda-817c897a4580/mainnet" //"https://cloudflare-eth.com" //wss://cool-muddy-butterfly.discover.quiknode.pro/0e41f42d5a7c9611f30ef800444bfcb93d3ae9a6
	//Ethereum Mainnet
	Address = strings.ToLower("0xff2b4721f997c242ff406a626f17df083bd2c568") //0xdAC17F958D2ee523a2206206994597C13D831ec7
)

var BlockWithTimestamp map[string]string = make(map[string]string)

var TransferTopic = "0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"
var ApprovalTopic = "0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925"
var ApprovalForAllTopic = "0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31"

var Topic map[string]int = map[string]int{ // index starts from 1
	"0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925": 1, //Approval
	"0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef": 2, // Transfer
	"0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31": 3} // ApprovalForAll

type MysqlConFig struct {
	Addr            string
	Port            int
	Db              string
	Username        string
	Password        string
	MaxIdealConn    int
	MaxOpenConn     int
	ConnMaxLifetime int
}

var MysqlCon = MysqlConFig{
	"127.0.0.1",
	3306,
	"eventLog",
	"root",
	"123456",
	10,
	256,
	600,
}
