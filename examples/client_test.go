package examples

import (
	"context"
	"fmt"
	"main/common/config"
	"main/common/ethconn"
	"testing"

	"github.com/ethereum/go-ethereum/common"
)

func Test_Scanblockdata(t *testing.T) {
	client := ethconn.ConnBlockchain(config.EthServer)
	type args struct {
		address string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test with tx",
			args: args{
				address: "0xff2b4721f997c242ff406a626f17df083bd2c568",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bytecode, _ := client.CodeAt(context.Background(), common.HexToAddress(tt.args.address), nil)
			if len(bytecode) > 0 {
				fmt.Println(1)
			}
		})
		// t.Run(tt.name, func(t *testing.T) {
		// 	txinfo, _, err := client.TransactionByHash(context.Background(), common.HexToHash(tt.args.address))
		// 	if err != nil {
		// 		t.Errorf("Client.TransactionByHash() error = %v", err)
		// 		return
		// 	}
		// 	by, _ := txinfo.MarshalJSON()
		// 	fmt.Println(string(by))
		// })
	}
}
