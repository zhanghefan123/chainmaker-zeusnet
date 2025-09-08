package network

import (
	"fmt"
	"github.com/zhanghefan123/chainmaker-zeusnet/modules/config"
	"github.com/zhanghefan123/chainmaker-zeusnet/tools/execute"
	"strings"
)

// GetConnectedTcpConnectionCount 获取已建立连接的数量
func GetConnectedTcpConnectionCount() (int, error) {
	result, err := execute.CommandWithResult("netstat", []string{"-antp"})
	listenAddrWithPrefix := config.EnvLoaderInstance.ListenIpAddr
	listenAddr := listenAddrWithPrefix[:len(listenAddrWithPrefix)-3]
	// fmt.Printf("listen addr = %s\n", listenAddr)
	if err != nil {
		return 0, fmt.Errorf("execute netstat -antp | grep ESTABLISHED error: %v", err)
	} else {
		count := 0
		lines := strings.Split(result, "\n")
		// fmt.Println("----------------------------------")
		for _, line := range lines {
			if strings.Contains(line, listenAddr) && strings.Contains(line, "ESTABLISHED") {
				count += 1
			}
		}
		// fmt.Println("----------------------------------")
		return count, nil
	}
}

// GetHalfConnectedTcpConnectionCount 获取半开连接队列中的半开连接数量
func GetHalfConnectedTcpConnectionCount() (int, error) {
	result, err := execute.CommandWithResult("netstat", []string{"-antp"})
	listenAddrWithPrefix := config.EnvLoaderInstance.ListenIpAddr
	listenAddr := listenAddrWithPrefix[:len(listenAddrWithPrefix)-3]
	// [:len(fabricOrderNode.Interfaces[0].SourceIpv4Addr)-3]
	if err != nil {
		return 0, fmt.Errorf("execute netstat -antp | grep SYN_RECV error: %v", err)
	} else {
		count := 0
		lines := strings.Split(result, "\n")
		// fmt.Println("----------------------------------")
		for _, line := range lines {
			if strings.Contains(line, listenAddr) && strings.Contains(line, "SYN_RECV") {
				count += 1
			}
		}
		// fmt.Println("----------------------------------")
		return count, nil
	}
}
