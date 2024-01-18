package ethconn

import (
	"log"

	"github.com/ethereum/go-ethereum/ethclient"
)

func ConnBlockchain(str string) *ethclient.Client {
	nclient, err := ethclient.Dial(str)
	if err != nil {
		log.Fatalf("Eth connect error:%v", err)
	}
	return nclient
}
