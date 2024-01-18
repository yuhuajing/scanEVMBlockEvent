package tabletypes

type Transfer struct {
	Id          string `json:"id" gorm:"primary_key"`
	Blocknumber uint64 `json:"blocknumber"`
	Timestamp   uint64 `json:"timestamp"`
	Address     string `json:"address"`
	Func        string `json:"func"`
	From        string `json:"from" gencodec:"required"`
	Operator    string `json:"operator"`
	Tokenid     int64  `json:"tokenid"`
	Txhash      string `json:"txhash" gencodec:"required"`
	Logindex    uint   `json:"logindex"`
}

type Approval struct {
	Id          string `json:"id" gorm:"primary_key"`
	Blocknumber uint64 `json:"blocknumber"`
	Timestamp   uint64 `json:"timestamp"`
	Address     string `json:"address"`
	Func        string `json:"func"`
	From        string `json:"from" gencodec:"required"`
	Operator    string `json:"operator"`
	//Approved    string `json:"approved"`
	Tokenid  int64  `json:"tokenid"`
	Txhash   string `json:"txhash" gencodec:"required"`
	Logindex uint   `json:"logindex"`
}

type ApprovalForAll struct {
	Id          string `json:"id" gorm:"primary_key"`
	Blocknumber uint64 `json:"blocknumber"`
	Timestamp   uint64 `json:"timestamp"`
	Address     string `json:"address"`
	Func        string `json:"func"`
	From        string `json:"from"`
	Operator    string `json:"operator"`
	//Approved    bool   `json:"approved"`
	Txhash   string `json:"txhash" gencodec:"required"`
	Logindex uint   `json:"logindex"`
}

type Owner struct {
	Id          string `json:"id" gorm:"primary_key"`
	Blocknumber uint64 `json:"blocknumber"`
	Logindex    uint   `json:"logindex"`
	Address     string `json:"address"`
	Owner       string `json:"owner"`
	Tokenid     int    `json:"tokenid"`
}
