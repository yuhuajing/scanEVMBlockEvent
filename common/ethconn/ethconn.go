package ethconn

import (
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/ethclient"
)

func ConnBlockchain(str string) *ethclient.Client {
	nclient, err := ethclient.Dial(str)
	if err != nil {
		fmt.Printf("Eth connect error:%s\n", err)
		log.Fatal(err)
	}
	return nclient
}
