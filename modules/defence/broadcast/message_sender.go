package broadcast

import (
	"chainmaker.org/chainmaker/common/v2/msgbus"
	netpb "chainmaker.org/chainmaker/pb-go/v2/net"
	"github.com/gogo/protobuf/proto"
	"github.com/zhanghefan123/chainmaker-zeusnet/modules/defence/info"
	defenceMessage "github.com/zhanghefan123/chainmaker-zeusnet/modules/defence/message"
	"github.com/zhanghefan123/chainmaker-zeusnet/tools/serialization"
)

// GenerateDdosAnnouncementMessage 生成DDOS攻击通告消息
func (d *DefenceDdosModule) GenerateDdosAnnouncementMessage(underDdosAttack bool) *defenceMessage.DdosAnnouncementMsg {
	// 创建 DdosAnnouncementMsg
	ddosAnnouncementMsg := defenceMessage.DdosAnnouncementMsg{
		Id:              info.InformationInstance.LocalPeerId, // 当前的节点的 Id
		UnderDdosAttack: underDdosAttack,                      // 是否正在遭受攻击
	}
	return &ddosAnnouncementMsg
}

// ConvertToDefenceMsg 转换成防御消息
func (d *DefenceDdosModule) ConvertToDefenceMsg(msg proto.Message) *defenceMessage.DefenceMsg {
	return &defenceMessage.DefenceMsg{
		Type: defenceMessage.DefenceMsgType_MSG_DDOS_ANNOUNCEMENT,
		Msg:  serialization.MustMarshal(msg),
	}
}

// BroadCastInNetMsg 以 NetMsg 的形式向所有的 validators 进行广播
func (d *DefenceDdosModule) BroadCastInNetMsg(msg proto.Message) {
	// 所有的验证者
	var validators []string
	// 错误
	var err error
	// 当前 id
	var currentId string
	// 获取验证者集合
	validators, err = info.InformationInstance.TbftConsensusImpl.GetValidators()
	if err != nil {
		return
	}
	// 获取当前 validator id
	currentId = info.InformationInstance.LocalPeerId
	// 遍历所有的验证者除了自己
	for _, v := range validators {
		// 不向自己进行消息的发送
		if v == currentId {
			continue
		} else {
			// 开启线程发送消息
			go func(validator string) {
				netMsg := &netpb.NetMsg{
					Payload: serialization.MustMarshal(msg),
					Type:    netpb.NetMsg_DDOS_STATUS_MSG,
					To:      validator,
				}
				d.MsgBus.Publish(msgbus.SendDdosStatusMsg, netMsg)
			}(v)
		}
	}
}

// BroadcastDefenceMsg 广播防御性消息
func (d *DefenceDdosModule) BroadcastDefenceMsg(underAttack bool) {
	// 创建 ddos 通告消息
	ddosAnnouncementMsg := d.GenerateDdosAnnouncementMessage(underAttack)
	// 转换为防御消息
	defenceMsg := d.ConvertToDefenceMsg(ddosAnnouncementMsg)
	// 将消息广播除去
	d.BroadCastInNetMsg(defenceMsg)
}
