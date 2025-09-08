package network_interface

import (
	"fmt"
	"github.com/vishvananda/netlink"
	"golang.org/x/sys/unix"
	"net"
	"time"
)

// CheckNetworkInterface 用来进行检查接口是否UP并且被设置了IP地址
// @param interfaceName 接口的名称
// @param listenIpAddress 这个接口需要是这个地址
// @return 是否可以监听这个地址了
func CheckNetworkInterface(interfaceName, listenIpAddr string) error {
	ch := make(chan netlink.LinkUpdate)
	done := make(chan struct{})
	defer close(done)
	if err := netlink.LinkSubscribe(ch, done); err != nil {
		return fmt.Errorf("cannot subscribe to link updates: %v", err)
	}
	// 进行循环的监听
	for {
		select {
		case update := <-ch:
			if update.Header.Type == unix.RTM_NEWLINK {
				currentInterfaceName := update.Link.Attrs().Name
				if currentInterfaceName == interfaceName && ((update.Link.Attrs().Flags & net.FlagUp) != 0) {
					// 打印所有的ip地址
					list, err := netlink.AddrList(update.Link, netlink.FAMILY_V4)
					if err != nil {
						return fmt.Errorf("cannot list IP addresses: %v", err)
					} else {
						ipAddrSet := false
						for _, ip := range list {
							if listenIpAddr[:(len(listenIpAddr)-3)] == ip.IP.String() {
								ipAddrSet = true
								break
							}
						}
						if ipAddrSet {
							return nil
						}
					}
				}
			}
		default:
			fmt.Println("cannot find interface wait 5 seconds")
			time.Sleep(5 * time.Second)
		}
	}
}

// ReturnUntilInterfaceUp 循环检查一个接口是否存在, 如果这个接口存在并且 up, 就进行返回
func ReturnUntilInterfaceUp(interfaceName, listenIpAddr string) error {
	err := CheckNetworkInterface(interfaceName, listenIpAddr)
	if err != nil {
		return fmt.Errorf("cannot check network interface: %w", err)
	}
	fmt.Printf("Network interface %s exists!\n", interfaceName)
	return nil
}
