package interface_down

import (
	"fmt"
	"time"
	"zeusnet.com/zhf/zeusnet/global_var"
	"zeusnet.com/zhf/zeusnet/modules/config"
	"zeusnet.com/zhf/zeusnet/modules/network_interface"
)

// SpeedCheckAndInterfaceDown 进行速度检查
func SpeedCheckAndInterfaceDown(interfaceName string) {
	var currentRecvBytes int64 // 本次收到的数据量
	var lastRecvBytes int64    // 上次收到的数据量
	var delta int64            // 数据量的差值
	var dataRate float64       // 数据速率
	for {
		currentRecvBytes = network_interface.GetRxRecv(interfaceName)
		delta = currentRecvBytes - lastRecvBytes    // 更新差值
		dataRate = float64(delta) / 1024.0 / 1024.0 // 更新速率
		lastRecvBytes = currentRecvBytes            // 更新上次收到的数据量
		if dataRate > config.EnvLoaderInstance.DDoSWarningRate {
			// ------------------- 关闭 tcp 连接 ---------------------
			fmt.Println("reach ddos warning rate -> disable interface")
			err := global_var.GlobalStructure.Net.Stop()
			if err != nil {
				fmt.Println("could not disable network interface")
			} else {
				break
			}
			// ------------------- 关闭 tcp 连接 ---------------------
		} else {
			fmt.Println("still not reach limit rate, not disable interface")
		}
		time.Sleep(time.Second)
	}
}
