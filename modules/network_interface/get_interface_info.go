package network_interface

import (
	"github.com/zhanghefan123/chainmaker-zeusnet/tools/fileop"
	"github.com/zhanghefan123/chainmaker-zeusnet/tools/intfop"
	"strings"
)

// GetRxRecv 获取接口速度
func GetRxRecv(interfaceName string) int64 {
	var currentRecvBytes int64 // 本次收到的数据量
	// 读取 /proc/net/dev 文件
	lines, _ := fileop.ReadLines("/proc/net/dev")
	for _, singleLine := range lines {
		if strings.Contains(singleLine, interfaceName) {
			interfaceData := intfop.ResolveNetworkInterfaceLine(singleLine)
			currentRecvBytes = int64(interfaceData.RxBytes) // 更新当前收到的数据量
			break                                           // 已经找到退出循环
		} else {
			continue // 否则继续循环, 找到相应的接口
		}
	}
	return currentRecvBytes
}
