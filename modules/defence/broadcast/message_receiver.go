package broadcast

import (
	"chainmaker.org/chainmaker/common/v2/msgbus"
	tbftpb "chainmaker.org/chainmaker/pb-go/v2/consensus/tbft"
	"fmt"
	"zeusnet.com/zhf/zeusnet/modules/defence/info"
)

// ResolveDdosAnnouncementMessage 解析DdosAnnouncement消息
func (d *DefenceDdosModule) ResolveDdosAnnouncementMessage(message *msgbus.Message) error {
	// 进行消息的解析
	nodeId, underAttack, err := d.ProcessMessage(message)
	if err != nil {
		return fmt.Errorf("cannot process message: %w", err)
	}

	// 如果消息之中是某个节点正在遭受攻击
	if underAttack {
		fmt.Println("receive message under attack")
		// 1. 如果这个通告的被攻击节点还不在受攻击节点 map 之中
		if _, ok := (*d.AttackedPointMap)[nodeId]; !ok {
			// 获取当前的 proposer
			currentProposer := info.InformationInstance.TbftConsensusImpl.GetProposer()
			// 如果当前的 proposer 就是那个被攻击的节点, 那么立马进入下一轮 (所有的其他节点立即进入下一轮的共识)
			if (info.InformationInstance.TbftConsensusImpl.Step == tbftpb.Step_PROPOSE) && (currentProposer == nodeId) {
				// 旧的高度
				currentHeight := info.InformationInstance.TbftConsensusImpl.Height
				// 新的轮次
				newRound := info.InformationInstance.TbftConsensusImpl.Round + 1
				// 进入新的一轮
				info.InformationInstance.TbftConsensusImpl.EnterNewRound(currentHeight, newRound)
			}
			// 将被攻击节点放到被攻击节点 map 之中
			(*d.AttackedPointMap)[nodeId] = struct{}{}
		}
	} else { // 2.如果消息之中是某个节点已经不再遭受攻击
		fmt.Println("receive message not under attack")
		// 2.1 判断是否是直接删除
		if d.EnableDirectRemove {
			if _, ok := (*d.AttackedPointMap)[nodeId]; ok {
				delete(*d.AttackedPointMap, nodeId)
			}
		} else {
			if _, ok := (*d.AttackedPointMap)[nodeId]; ok {
				// 等到同步之后再进行删除
				d.RemoveFromAttackedPointMapUntilSync(nodeId)
			}
		}
	}
	return nil
}
