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
	// client, err := ethclient.Dial("wss://eth.getblock.io/ab0b1aa0-b490-4dc0-9bda-817c897a4580/mainnet")
	// if err != nil {
	// 	fmt.Printf("Eth connect error:%s\n", err)
	// 	log.Fatal(err)
	// }
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
				address: "0x7e12871085579eca01f4513825f46a4c3123816ccc00203c35906079d531a319",
			},
		},
		{
			name: "test with tx2", //53
			args: args{
				address: "0x1fb3c04883ba857bbb7074f347c75f17724d6b8167804aca2c3136b117c4c6ef",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			txinfo, _, err := client.TransactionByHash(context.Background(), common.HexToHash(tt.args.address))
			if err != nil {
				t.Errorf("Client.TransactionByHash() error = %v", err)
				return
			}
			by, _ := txinfo.MarshalJSON()
			fmt.Println(string(by))
		})
	}
}
