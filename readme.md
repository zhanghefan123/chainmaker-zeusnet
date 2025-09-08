# 1. 启动步骤

- [1]. chainmaker-go/main/cmd/main.go 之中调用 StartCMD 
- [2]. StartCMD 内部调用 zeusnet.Start(), 而 zeusnet.Start 函数定义在 chainmaker-go/zeusnet/starter.go 之中