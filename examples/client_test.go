package examples

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"main/common/config"
	"main/common/ethconn"
	"main/wallet"
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

func Test_generateAddr(t *testing.T) {
	priKey, _ := crypto.HexToECDSA("94498129232b8fcbd91e53c1c3cc86e9c415ac0eec5acedd7dba4aaba1226ac8")
	publicKey := priKey.Public()
	publicKeyECDSA, _ := publicKey.(*ecdsa.PublicKey)
	fromaddress := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	fmt.Println(fromaddress)
}

func Test_wallet_read(t *testing.T) {
	//priKey, _ := crypto.HexToECDSA("d34772897e5d6bc952cff1094945d6d05bca81decd773f0cfb3575fbc4a73493")
	client := ethconn.ConnBlockchain("http://localhost:8545")
	// publicKey := priKey.Public()
	// publicKeyECDSA, _ := publicKey.(*ecdsa.PublicKey)
	// fromaddress := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	//fmt.Println(fromaddress)
	address := common.HexToAddress("0xbB56d2FAd712F29D74241e9C45936Ac9fd049640")
	instance, err := wallet.NewWallet(address, client)
	if err != nil {
		fmt.Println("error creating instance")
		log.Fatal(err)
	}

	number, _ := instance.Gettestnum(nil)
	fmt.Println(number)

	// bal, _ := client.BalanceAt(context.Background(), common.HexToAddress(fromaddress), nil)
	// fmt.Println(bal)
}
func Test_wallet_write(t *testing.T) {
	client := ethconn.ConnBlockchain("http://localhost:8545")
	privateKey, _ := crypto.HexToECDSA("d34772897e5d6bc952cff1094945d6d05bca81decd773f0cfb3575fbc4a73493")
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice

	//fmt.Println(fromaddress)
	address := common.HexToAddress("0xbB56d2FAd712F29D74241e9C45936Ac9fd049640")
	instance, err := wallet.NewWallet(address, client)
	if err != nil {
		fmt.Println("error creating instance")
		log.Fatal(err)
	}
	newnumber := big.NewInt(99)
	tx, err := instance.Settestnum(auth, newnumber)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("tx sent: %s\n", tx.Hash().Hex()) // tx sent: 0x8d490e535678e9a24360e955d75b27ad307bdfb97a1dca51d0f3035dcee3e870
	time.Sleep(10 * time.Second)
	number, _ := instance.Gettestnum(nil)
	fmt.Println("new Number: ", number)
}

//94498129232b8fcbd91e53c1c3cc86e9c415ac0eec5acedd7dba4aaba1226ac8
//fad9c8855b740a0b7ed4c221dbad0f33a83a49cad6b3fe8d5817ac83d38b6b1a

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
