module zeusnet.com/zhf/zeusnet

go 1.18

// go.etcd.io/etcd/client/v3 v3.5.0

require (
	chainmaker.org/chainmaker/common/v2 v2.3.2
	chainmaker.org/chainmaker/consensus-tbft/v2 v2.3.3
	chainmaker.org/chainmaker/pb-go/v2 v2.3.3
	chainmaker.org/chainmaker/protocol/v2 v2.3.3
	github.com/gin-gonic/gin v1.5.0
	github.com/vishvananda/netlink v1.3.0
	go.etcd.io/etcd/client/v3 v3.5.0
)

// 注意: 由于 zeusnet 之中引用了 consensus-tbft-v2.3.3 所以不能在 consensus-tbft-v2.3.3 之中再进行引入 zeusnet

replace (
	chainmaker.org/chainmaker/common/v2 v2.3.2 => ../chainmaker-go/common-v2.3.2
	chainmaker.org/chainmaker/consensus-tbft/v2 v2.3.3 => ../chainmaker-go/consensus-tbft-v2.3.3
	chainmaker.org/chainmaker/pb-go/v2 v2.3.3 => ../chainmaker-go/pb-go-v2.3.3
)
