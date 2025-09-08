package connections

import (
	"fmt"
	"github.com/shirou/gopsutil/net"
	"time"
)

// ListenTcpEstablishedConnections 处理 TCP 建立的连接
func ListenTcpEstablishedConnections(localAddr string) {
	for {
		count := 0
		connections, err := net.Connections("tcp")
		if err != nil {
			fmt.Printf("canot listen tcp established connections %v", err)
		}
		for _, conn := range connections {
			if conn.Status == "ESTABLISHED" {
				listenAddr := conn.Laddr.IP
				if listenAddr == localAddr[:len(localAddr)-3] {
					count += 1
				}
				//fmt.Printf("Local: %s:%d -> Remote: %s:%d\n",
				//	conn.Laddr.IP, conn.Laddr.Port,
				//	conn.Raddr.IP, conn.Raddr.Port)
			}
		}
		// fmt.Printf("established connections = %d\n", count)
		time.Sleep(1 * time.Second)
	}
}
