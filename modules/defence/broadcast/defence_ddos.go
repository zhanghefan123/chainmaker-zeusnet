package broadcast

import (
	"chainmaker.org/chainmaker/common/v2/msgbus"
	"chainmaker.org/chainmaker/protocol/v2"
	"fmt"
	"math"
	"time"
	"zeusnet.com/zhf/zeusnet/modules/config"
	"zeusnet.com/zhf/zeusnet/modules/defence/info"
)

var (
	MsgBusTopics      = []msgbus.Topic{msgbus.RecvDdosStatusMsg}
	StartDefenceQueue = make(chan bool)
)

type DefenceDdosModule struct {
	Name       string            // 模块的名称
	Logger     protocol.Logger   // 可能需要日志记录器
	MsgBus     msgbus.MessageBus // 消息总线
	StartOrNot bool              // 是否调用启动函数

	LastRecvBytes int64 // 上一次接收到的字节数量

	EnabledBroadcastDefence bool    // 启用 ddos 攻击防御机制
	EnableDirectRemove      bool    // 直接删除
	DdosWarningRate         float64 // 到达什么攻击速率才进行通告

	AttackedPointMap *map[string]struct{} // 被攻击的节点的列表
	SyncDeleteMap    map[string]struct{}  // sync delete map

	StartDefenceOrNot bool // 是否启动周期性广播防御
}

// Start DefenceDdosModule 的启动方法
func (d *DefenceDdosModule) Start() error {
	if d.EnabledBroadcastDefence {
		d.Logger.Errorf("ddos defence module start")
		d.SubscribeTopics()  // 进行消息的订阅
		d.BroadcastDefence() // 周期性的进行广播
		return nil
	} else {
		d.Logger.Errorf("ddos defence module start")
		return nil
	}
}

// SubscribeTopics 订阅消息
func (d *DefenceDdosModule) SubscribeTopics() {
	for _, topic := range MsgBusTopics {
		d.MsgBus.Register(topic, d) // 订阅主题
	}
}

// Stop DefenceDdosModule 的停止方法
func (d *DefenceDdosModule) Stop() error {
	d.Logger.Errorf("ddos defence module stop")
	return nil
}

// OnMessage 要成为 Subscriber 需要实现的第一个方法
func (d *DefenceDdosModule) OnMessage(message *msgbus.Message) {
	switch message.Topic {
	case msgbus.RecvDdosStatusMsg:
		// 如果收到了 ddos 通告消息则进行处理
		err := d.ResolveDdosAnnouncementMessage(message)
		if err != nil {
			fmt.Printf("resolve ddos announcement message error %v", err)
		}
	default:
		panic("unhandled default case")
	}
}

// OnQuit 要成为 Subscriber 需要实现的第二个方法
func (d *DefenceDdosModule) OnQuit() {
	d.Logger.Errorf("quit")
}

// BroadcastDefence 通过广播来进行防御
func (d *DefenceDdosModule) BroadcastDefence() {
	go func() {
		for {
			// fmt.Println("---------------------------------------")
			for item, _ := range *d.AttackedPointMap {
				fmt.Printf("attacked node=%s\n", item)
			}
			// fmt.Println("---------------------------------------")
			var underDdosAttack bool
			// 检查是否发生了 ddos 攻击
			// --------------------------------------------------------------------------------
			underDdosAttack = d.CheckIfUnderDdos(config.EnvLoaderInstance.InterfaceName)
			// --------------------------------------------------------------------------------
			select {
			case d.StartDefenceOrNot = <-StartDefenceQueue:
			default:
				if d.StartDefenceOrNot {
					// 对 ddos 攻击进行处理
					// --------------------------------------------------------------------------------
					if underDdosAttack {
						// 进行消息的打印
						d.Logger.Errorf("%s under ddos attack", info.InformationInstance.LocalPeerId)
						// 如果遭受攻击, 将自己加入
						(*d.AttackedPointMap)[info.InformationInstance.LocalPeerId] = struct{}{}
						// 并且向其他节点进行消息的广播
						d.BroadcastDefenceMsg(underDdosAttack)
					} else {
						// 如果没有遭受攻击
						if d.EnableDirectRemove { // 处理方式1: 直接删除自己
							if _, ok := (*d.AttackedPointMap)[info.InformationInstance.LocalPeerId]; ok {
								delete(*d.AttackedPointMap, info.InformationInstance.LocalPeerId)
							}
						} else { // 处理方式2: 判断自己是否存在 -> 如果存在的话等到同步了之后再删除
							if _, ok := (*d.AttackedPointMap)[info.InformationInstance.LocalPeerId]; ok {
								// 等到区块高度同步之后再将自己移除
								_ = d.SelfRemoveUntilSync()
							}
							// 如果区块高度已经同步, 并且已经移除
							d.BroadcastDefenceMsg(false)
						}
					}
				}
			}
			// --------------------------------------------------------------------------------

			// 周期性进行检测
			time.Sleep(time.Second)
		}
	}()
}

// SelfRemoveUntilSync 自己等到区块同步之后，再让这个节点成为主节点, 本节点的操作
func (d *DefenceDdosModule) SelfRemoveUntilSync() bool {
	// 获取最高的区块高度
	var maxHeight = info.InformationInstance.GetMaxBlockHeight()
	// 获取本节点区块高度
	var currentHeight = info.InformationInstance.TbftConsensusImpl.Height
	// 判断是否高度相等
	if math.Abs(float64(currentHeight-maxHeight)) < 10 {
		// 从字典之中移除自己, 让自己重新能成为主节点
		delete(*d.AttackedPointMap, info.InformationInstance.TbftConsensusImpl.Id)
		// 进行日志的输出 -> 说明已经移除
		d.Logger.Errorf("SelfRemoveUntilSync success currentHeight: %d rightHeight: %d",
			info.InformationInstance.TbftConsensusImpl.Height, maxHeight)
		return true
	} else {
		// 进行日志的输出 -> 说明还不进行移除
		d.Logger.Errorf("SelfRemoveUntilSync failed currentHeight: %d rightHeight: %d wating for sync",
			info.InformationInstance.TbftConsensusImpl.Height, maxHeight)
		return false
	}
}

// RemoveFromAttackedPointMapUntilSync 等到区块同步之后，再让这个节点成为主节点, 其他节点收到的操作
func (d *DefenceDdosModule) RemoveFromAttackedPointMapUntilSync(removedNodeId string) {
	// 获取最高的区块高度
	var maxHeight = info.InformationInstance.GetMaxBlockHeight()
	// 获取要被移除的节点的区块高度
	removedNodeHeight := (*info.InformationInstance.PeersHeight)[removedNodeId]
	// 一定需要等到完全同步之后才能重新恢复其成为主节点的权利
	if math.Abs(float64(removedNodeHeight-maxHeight)) < 10 {
		delete(*d.AttackedPointMap, removedNodeId)
		d.Logger.Errorf("RemoveFromAttackedPointMapUntilSync success removedNodeHeight: %d rightHeight: %d",
			removedNodeHeight, maxHeight)
	} else {
		d.Logger.Errorf("RemoveFromAttackedPointMapUntilSync failed removedNodeHeight: %d rightHeight: %d wating for sync",
			removedNodeHeight, maxHeight)
	}
}
