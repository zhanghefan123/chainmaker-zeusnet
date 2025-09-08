package zeusnet

import (
	"fmt"
	"github.com/zhanghefan123/chainmaker-zeusnet/modules/config"
	"github.com/zhanghefan123/chainmaker-zeusnet/modules/connections"
	"github.com/zhanghefan123/chainmaker-zeusnet/modules/defence/interface_down"
	"github.com/zhanghefan123/chainmaker-zeusnet/modules/etcd"
	"github.com/zhanghefan123/chainmaker-zeusnet/modules/frr"
	"github.com/zhanghefan123/chainmaker-zeusnet/modules/network_interface"
	"github.com/zhanghefan123/chainmaker-zeusnet/services"
	"net/http"
)

// Start 自定义启动流程
func Start() error {
	// step1: 加载环境变量
	err := config.EnvLoaderInstance.LoadEnv()
	if err != nil {
		return fmt.Errorf("error loading environment variables: %w", err)
	}

	// step2: 测试能否进行 etcd client 的建立
	etcd.Client, err = etcd.NewEtcdClient(config.EnvLoaderInstance.EtcdListenAddr, config.EnvLoaderInstance.EtcdListenPort)
	if err != nil {
		return fmt.Errorf("error creating etcd client: %w", err)
	}

	// step3: 进行监听
	go etcd.HandleStartDefenceChange()

	// step4: 进行tcp建立连接的监听
	go connections.ListenTcpEstablishedConnections(config.EnvLoaderInstance.ListenIpAddr)

	// step3: 直到接口启动为止都别启动
	err = network_interface.ReturnUntilInterfaceUp(config.EnvLoaderInstance.InterfaceName,
		config.EnvLoaderInstance.ListenIpAddr)
	if err != nil {
		return fmt.Errorf("error while starting network interface: %w", err)
	}

	// step4: 启动 frr
	if config.EnvLoaderInstance.EnableFrr {
		frr.Start(config.EnvLoaderInstance.ContainerName)
		fmt.Println("frr started")
	} else {
		fmt.Println("frr is disabled")
	}

	// step5: 判断是否需要监听接口, 并直接进行接口的断开操作
	if config.EnvLoaderInstance.SpeedCheck {
		go interface_down.SpeedCheckAndInterfaceDown(config.EnvLoaderInstance.InterfaceName)
		fmt.Println("speed check started")
	} else {
		fmt.Println("speed check is disabled")
	}

	// step6: 进行启动
	router := services.InitRouter()
	fabricServer := &http.Server{
		Addr:    fmt.Sprintf(":%d", config.EnvLoaderInstance.WebServerListenPort),
		Handler: router,
	}
	go func() {
		// 这个也是一个阻塞的进程
		err = fabricServer.ListenAndServe()
		if err != nil {
			fmt.Printf("failed to start self gin server: %s", err)
		}
		fmt.Printf("start fabric server on port %d successfully\n", config.EnvLoaderInstance.WebServerListenPort)
	}()
	return nil
}
