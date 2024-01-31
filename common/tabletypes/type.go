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

type Startblocks struct {
	Id                 string `json:"id" gorm:"primary_key"`
	Historyblocknumber uint64 `json:"historyblocknumber"`
	//Newblocknumber     uint64 `json:"newblocknumber"`
	Address string `json:"address"`
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
	Timestamp   uint64 `json:"timestamp"`
	Logindex    uint   `json:"logindex"`
	Address     string `json:"address"`
	Owner       string `json:"owner"`
	Tokenid     int    `json:"tokenid"`
	Level       string `json:"level"`
}

const (
	//StatusUnknown is a Status of type unknown.
	StatusUnknown string = ""
	//StatusListing is a Status of type listing.
	StatusListing string = "listing"
	//StatusCancel is a Status of type cancel.
	StatusCancel  string = "cancel"
	StatusSold    string = "sold"
	StatusInvalid string = "invalid"
)

type OpenseaOrder struct {
	Id             string `json:"id" gorm:"primary_key"`
	Listingtime    int    `json:"listingtime"`
	Expirationtime int    `json:"expirationtime"`
	Orderhash      string `json:"orderhash"`
	Owner          string `json:"owner"`
	Address        string `json:"address"`
	Tokenid        int    `json:"tokenid"`
	Status         string `json:"status"`
	Collection     string `json:"collection"`
}

type MarketOrder struct {
	Id             string `json:"id" gorm:"primary_key"`
	Listingtime    int    `json:"listingtime"`
	Expirationtime int    `json:"expirationtime"`
	Owner          string `json:"owner"`
	Address        string `json:"address"`
	Tokenid        int    `json:"tokenid"`
	Status         string `json:"status"`
	Domain         string `json:"domain"`
	Name           string `json:"name"`
}
