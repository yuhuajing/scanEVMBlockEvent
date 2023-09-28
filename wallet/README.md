# 代理合约

factory.sol: 通过写入字节吗的方式实现代理，将收到的所有数据通过fallback函数 delegate 到代理中去

Wallet.sol: 实现合约，实现工厂合约的实际业务逻辑

测试：
1. 部署 Wallet.sol
2. 部署 factory.sol
3. 调用 factory.sol 中的 createAccount 函数，输入 Wallet.sol 的地址和其他参数，生成新的钱包地址
4. 通过  Wallet.sol 的 abi 和 bytecode 生成 对 Wallet.sol 合约的调用 golang代码 （abigen --bin=./wallet.bin --abi=./wallet.abi --pkg=wallet --out=wallet.go）
5. 将 对合约 Wallet.sol 的调用全部转到 新的钱包地址，然后再 fallback 回 Wallet.sol 合约

## 测试签名
```javascript
    //Signer 0xd514Ca657E536bB30962A31b22B4F39183328E0F
    const signer_prikey = "7174cd9f4f8cd8bd8c91898744ba231b5db50d5191f14819408d51ccf8c6a8c9"
	const signerstr = "signerMessage"
	// Signer签名:0x0dafb1883ddad52462db50271df078c380b7737d5a98dbc404585d7a629187556caf6982e56969b3c342bc42664b0e1c4b23452ee6343e0995ee954f05a08c081c

    //Manager 0xf5fBB766074124A574fc9aFaF9c9f139e7efB981
    const manager_prikey = "f07a77cb019764a524dce24cb47ac62bb231b4f0d7bab5f864f603f8cb0e344c"
	const managerstr = "managerMessage"
	// Manager签名:0xb70c777d43a46f6f10907d0ce5cbb4b5f41d9946d1ddf2e9dbb9a65116d4cacd01666d10b1d20b38e6e41ccd8cb82a5da9c6245b9a35c7cdf02517a407b6e2991b

    //Owner 0x156b6c24e78fede687950ba52a0b6b15a2c0ae11
    const owner_prikey = "867601ac4dc7028894d3ec525199a5289eeaa9ae38deba3a02511b31ce274901"
    //New_owner 0x9ab95fbf671a3b40f977eb116f948f69b26e663d
    const new_owner_prikey = "5567ceafd8404b4d3578f454cd7b78e82c1bdd7711cabdc0b10da72c4d0a24f8"

```

测试读数据代码：
```golang
func Test_wallet(t *testing.T) {
	//priKey, _ := crypto.HexToECDSA("d34772897e5d6bc952cff1094945d6d05bca81decd773f0cfb3575fbc4a73493")
	client := ethconn.ConnBlockchain("http://localhost:8545")
	// publicKey := priKey.Public()
	// publicKeyECDSA, _ := publicKey.(*ecdsa.PublicKey)
	// fromaddress := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	//fmt.Println(fromaddress)
	address := common.HexToAddress("0xbB56d2FAd712F29D74241e9C45936Ac9fd049640") // 钱包地址
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
```

测试写数据：
```golang
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
```