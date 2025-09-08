package broadcast

import (
	"chainmaker.org/chainmaker/common/v2/msgbus"
	netpb "chainmaker.org/chainmaker/pb-go/v2/net"
	"fmt"
	defenceMessage "zeusnet.com/zhf/zeusnet/modules/defence/message"
	"zeusnet.com/zhf/zeusnet/tools/serialization"
)

// ProcessMessage 处理消息
func (d *DefenceDdosModule) ProcessMessage(message *msgbus.Message) (string, bool, error) {
	netMessage, ok := message.Payload.(*netpb.NetMsg)
	if ok {
		newDefenceMessage := new(defenceMessage.DefenceMsg)
		newAnnouncementMessage := new(defenceMessage.DdosAnnouncementMsg)
		serialization.MustUnmarshal(netMessage.Payload, newDefenceMessage)
		serialization.MustUnmarshal(newDefenceMessage.Msg, newAnnouncementMessage)
		nodeId := newAnnouncementMessage.Id
		underAttack := newAnnouncementMessage.UnderDdosAttack
		return nodeId, underAttack, nil
	} else {
		return "", false, fmt.Errorf("cannot process message")
	}
}
