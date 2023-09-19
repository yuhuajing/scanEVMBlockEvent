# 代理合约

factory.sol: 通过写入字节吗的方式实现代理，将收到的所有数据通过fallback函数 delegate 到代理中去

Wallet.sol: 实现合约，实现工厂合约的实际业务逻辑

测试：
1. 部署 Wallet.sol
2. 部署 factory.sol
3. 调用 factory.sol 中的 createAccount 函数，输入 Wallet.sol 的地址和其他参数，生成新的钱包地址
4. 通过  Wallet.sol 的 abi 和 bytecode 生成 对 Wallet.sol 合约的调用 golang代码 （abigen --bin=./wallet.bin --abi=./wallet.abi --pkg=wallet --out=wallet.go）
5. 将 对合约 Wallet.sol 的调用全部转到 新的钱包地址，然后再 fallback 回 Wallet.sol 合约

测试代码：
```golang
func Test_wallet(t *testing.T) {
	//priKey, _ := crypto.HexToECDSA("d34772897e5d6bc952cff1094945d6d05bca81decd773f0cfb3575fbc4a73493")
	client := ethconn.ConnBlockchain("http://localhost:8545")
	// publicKey := priKey.Public()
	// publicKeyECDSA, _ := publicKey.(*ecdsa.PublicKey)
	// fromaddress := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	//fmt.Println(fromaddress)
	address := common.HexToAddress("0xb3CC975072D317069df4Bb6C143358DF05507Be5") // 钱包地址
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