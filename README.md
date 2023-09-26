# scanEVMBlockData

1. flag 分支作为扫链用的分支

用户可以提供以下参数：

```text
	startblock   // 起始高度，默认是Transfer表格中最新一条数据的区块高度 (可选)
	address    // 合约地址，默认为"0xff2b4721f997c242ff406a626f17df083bd2c568"（可选）            
	blockchain        // 区块链网络，默认是以太坊 （可选）
	rpcserver         // RPC，可以自定义RPC节点（可选） 
	database      // 数据库，默认是"mongodb" （可选）
```

运行命令

编译 
> go build ./scanblockdata.go

1. 开始扫链

>  nohup ./scanblockdata --address="0xff2b4721f997c242ff406a626f17df083bd2c568" --startblock=17948501 >> scandata.log 2>&1 &


