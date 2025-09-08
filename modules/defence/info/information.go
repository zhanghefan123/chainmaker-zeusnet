package info

import (
	tbft "chainmaker.org/chainmaker/consensus-tbft/v2"
	"chainmaker.org/chainmaker/logger/v2"
	consensusPb "chainmaker.org/chainmaker/pb-go/v2/consensus"
	"fmt"
	"io/ioutil"
	"os"
	"time"
	"zeusnet.com/zhf/zeusnet/modules/config"
	"zeusnet.com/zhf/zeusnet/modules/network"
)

var (
	InformationInstance *DefenceModuleInformation
)

// DefenceModuleInformation 所需要的信息
type DefenceModuleInformation struct {
	Logger            *logger.CMLogger          // 日志
	LocalPeerId       string                    // 本地节点的 peerId
	PeersHeight       *map[string]uint64        // 对端节点的高度
	ConsensusType     consensusPb.ConsensusType // 共识类型
	TbftConsensusImpl *tbft.ConsensusTBFTImpl   // tbft 实现
	StopChannel       chan struct{}             // 用来进行停止信号的发送
}

type RecordInformation struct {
	BlockHeight          int
	ConnectedTcpCount    int
	HalfConnetedTcpCount int
	TimeoutCount         int
	BusMessageCount      int
}

// ShowPeersHeight 展示对端区块高度
// -------------------------------------------------------------------------------
/*
func (info *DefenceModuleInformation) ShowPeersHeight() {
	for {
		info.Logger.Errorf("--------------------peer block height------------------")
		for nodeId, height := range *info.PeersHeight {
			info.Logger.Errorf("node id %s, height %d", nodeId, height)
		}
		info.Logger.Errorf("--------------------peer block height------------------")
		time.Sleep(time.Second)
	}
}
*/
// -------------------------------------------------------------------------------

// GetMaxBlockHeight 获取最大区块高度
func (info *DefenceModuleInformation) GetMaxBlockHeight() uint64 {
	var maxBlockHeight uint64
	for _, peerHeight := range *info.PeersHeight {
		if peerHeight > maxBlockHeight {
			maxBlockHeight = peerHeight
		}
	}
	return maxBlockHeight
}

// WritePeersHeightPeriodically 周期性进行 peers 高度的写入
func (info *DefenceModuleInformation) WritePeersHeightPeriodically() {
Loop:
	for {
		select {
		case <-info.StopChannel:
			break Loop
		default:
			// 1. 获取高度
			// -------------------------------------------------------------------------------
			blockHeight := int(info.TbftConsensusImpl.Height)
			// -------------------------------------------------------------------------------
			// 2. 获取超时次数/消息总线中的消息数量
			timeoutCount := info.TbftConsensusImpl.TimeoutCount
			busMessageCount := info.TbftConsensusImpl.GetMessageCount()
			// 3. 获取 tcp 连接数和半连接数量
			// -------------------------------------------------------------------------------
			connectedConnectionCount, err := network.GetConnectedTcpConnectionCount()
			if err != nil {
				fmt.Printf("get connected connection count err: %v", err)
			}
			halfConnectedConnectionCount, err := network.GetHalfConnectedTcpConnectionCount()
			if err != nil {
				fmt.Printf("get half connected connection count err: %v", err)
			}
			// -------------------------------------------------------------------------------
			// 4. 记录信息
			// -------------------------------------------------------------------------------
			information := &RecordInformation{
				BlockHeight:          blockHeight,
				ConnectedTcpCount:    connectedConnectionCount,
				HalfConnetedTcpCount: halfConnectedConnectionCount,
				TimeoutCount:         timeoutCount,
				BusMessageCount:      busMessageCount,
			}
			err = WriteInformation(information)
			if err != nil {
				fmt.Printf("write information error: %v", err)
			}
			// -------------------------------------------------------------------------------

			time.Sleep(time.Second)
		}
	}
}

func WriteInformation(information *RecordInformation) error {
	// 路径
	filePath := fmt.Sprintf("/configuration/%s/information.stat", config.EnvLoaderInstance.ContainerName)
	// 进行结果的写入
	result := fmt.Sprintf("%d,%d,%d,%d,%d",
		information.BlockHeight,
		information.ConnectedTcpCount,
		information.HalfConnetedTcpCount,
		information.TimeoutCount,
		information.BusMessageCount)
	// 构建临时文件
	tmpFilename := filePath + ".tmp"
	err := ioutil.WriteFile(tmpFilename, []byte(result), 0644)
	if err != nil {
		return fmt.Errorf("write tmp file error: %v", err)
	}
	// 进行文件重命名
	err = os.Rename(tmpFilename, filePath)
	if err != nil {
		return fmt.Errorf("rename tmp file error: %v", err)
	}
	return nil
}
