package tabletypes

type Transfer struct {
	ID          uint   `gorm:"primary_key"`
	Blocknumber uint64 `json:"blocknumber"`
	Timestamp   string `json:"timestamp"`
	Address     string `json:"address"`
	Func        string `json:"func"`
	From        string `json:"from" gencodec:"required"`
	To          string `json:"to" gencodec:"required"`
	Tokenid     int64  `json:"tokenid"`
	Txhash      string `json:"txhash" gencodec:"required"`
	Logindex    uint   `json:"logindex"`
}

type Approval struct {
	ID          uint   `gorm:"primary_key"`
	Blocknumber uint64 `json:"blocknumber"`
	Timestamp   string `json:"timestamp"`
	Address     string `json:"address"`
	Func        string `json:"func"`
	Owner       string `json:"owner"`
	Approved    string `json:"approved"`
	Tokenid     int64  `json:"tokenid"`
	Txhash      string `json:"txhash" gencodec:"required"`
	Logindex    uint   `json:"logindex"`
}

type ApprovalForAll struct {
	ID          uint   `gorm:"primary_key"`
	Blocknumber uint64 `json:"blocknumber"`
	Timestamp   string `json:"timestamp"`
	Address     string `json:"address"`
	Func        string `json:"func"`
	Owner       string `json:"owner"`
	Operator    string `json:"operator"`
	Approved    bool   `json:"approved"`
	Txhash      string `json:"txhash" gencodec:"required"`
	Logindex    uint   `json:"logindex"`
}

type Owner struct {
	ID      uint   `gorm:"primary_key"`
	Address string `json:"address"`
	Owner   string `json:"owner"`
	Tokenid int64  `json:"tokenid"`
}
