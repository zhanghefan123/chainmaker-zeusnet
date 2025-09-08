package broadcast

import (
	"zeusnet.com/zhf/zeusnet/modules/config"
	"zeusnet.com/zhf/zeusnet/modules/network_interface"
)

// CheckIfUnderDdos 检测是否遭受到 Ddos 攻击
func (d *DefenceDdosModule) CheckIfUnderDdos(interfaceName string) bool {
	var currentRecvBytes int64 // 本次收到的数据量
	var delta int64            // 数据量的差值
	var dataRate float64       // 数据速率
	currentRecvBytes = network_interface.GetRxRecv(interfaceName)
	delta = currentRecvBytes - d.LastRecvBytes  // 更新差值
	dataRate = float64(delta) / 1024.0 / 1024.0 // 更新速率
	d.LastRecvBytes = currentRecvBytes          // 更新上次收到的数据量
	//fmt.Printf("data rate: %f, ddos warning rate: %f\n", dataRate, config.EnvLoaderInstance.DDoSWarningRate)
	if dataRate > config.EnvLoaderInstance.DDoSWarningRate {
		return true
	} else {
		return false
	}
}
