package checkaddress

import (
	"context"
	"fmt"
	"regexp"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func IsValidAddress(addr string) bool {
	re := regexp.MustCompile("0x[0-9a-fA-F]{40}$")
	return re.MatchString(addr)
}

func IsContractAddress(address string, client *ethclient.Client) bool {
	if !IsValidAddress(address) {
		fmt.Println("INCORRECT_ADDRESS")
		return false
	}
	bytecode, err := client.CodeAt(context.Background(), common.HexToAddress(address), nil)
	if err != nil {
		fmt.Printf("GET_CONTRACT_RUNTIME_CODE_ERROR: %s", err)
		return false
	}
	return len(bytecode) > 0
}
