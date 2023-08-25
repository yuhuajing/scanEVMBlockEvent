# scanEVMBlockData

1. main 分支作为扫链用的分支
      
扫链流程：

获取当前最新区块高度 H1
启动两个协程 A 和 B
A 协程 监听网络区块，用区块高度监听Event数据
B 协程 解析处理 历史Event数据

所有Event数据写入通道，等候处理
数据库设计：

监听的Event topic为：

	"0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925": 1, //Approval
	"0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef": 2, // Transfer
	"0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31": 3 // ApprovalForAll
数据表结构为：

package tabletypes

type Transfer struct {
	ID          uint   `gorm:"primary_key"`
	Blocknumber uint   `json:"blocknumber"`
	Func        string `json:"func"`
	From        string `json:"from" gencodec:"required"`
	To          string `json:"to" gencodec:"required"`
	Tokenid     int64  `json:"tokenid"`
	Txhash      string `json:"txhash" gencodec:"required"`
	Txindex     uint   `json:"txindex"`
}

type Approval struct {
	ID          uint   `gorm:"primary_key"`
	Blocknumber uint   `json:"blocknumber"`
	Func        string `json:"func"`
	Owner       string `json:"owner"`
	Approved    string `json:"approved"`
	Tokenid     int64  `json:"tokenid"`
	Txhash      string `json:"txhash" gencodec:"required"`
	Txindex     uint   `json:"txindex"`
}

type ApprovalForAll struct {
	ID          uint   `gorm:"primary_key"`
	Blocknumber uint   `json:"blocknumber"`
	Func        string `json:"func"`
	Owner       string `json:"owner"`
	Operator    string `json:"operator"`
	Approved    string `json:"approved"`
	Txhash      string `json:"txhash" gencodec:"required"`
	Txindex     uint   `json:"txindex"`
}

type Owner struct {
	ID      uint   `gorm:"primary_key"`
	Address string `json:"address"`
	Owner   string `json:"owner"`
	Tokenid int64  `json:"tokenid"`
}


calculate keccak256 hash of the event topic

https://emn178.github.io/online-tools/keccak_256.html

编译 
> go build ./scanblockdata.go

1. 开始扫链

>  nohup ./scanblockdata >> scandata.log 2>&1 &


