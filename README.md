# scanEVMBlockData

1. 测试

在 example 目录下，执行```go test``` 测试 RPC 功能 

根据项目函数增加测试函数

2. main 分支作为扫链用的分支
      

编译 
> go build ./scanblockdata.go

1. 开始扫链

>  nohup ./scanblockdata >> scandata.log 2>&1 &


