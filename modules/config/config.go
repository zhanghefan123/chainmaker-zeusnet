package config

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
	"zeusnet.com/zhf/zeusnet/tools/fileop"
)

var (
	EnvLoaderInstance = &EnvLoader{
		PeerIdToContainerNameMapping: make(map[string]string),
	}
)

type EnvLoader struct {
	EnableBroadcastDefence bool
	EnableDirectRemove     bool
	EnableFrr              bool
	SpeedCheck             bool
	InterfaceName          string
	ListenIpAddr           string
	ContainerName          string
	DDoSWarningRate        float64
	BlocksPerProposer      int
	TimeoutProposeOptimal  time.Duration
	ProposeOptimal         bool
	WebServerListenPort    int

	StartDefenceKey string
	EtcdListenAddr  string
	EtcdListenPort  int

	PeerIdToContainerNameMapping map[string]string
}

// LoadEnv 进行环境变量的加载
func (el *EnvLoader) LoadEnv() error {
	// 1. 加载是否启用 ddos 通告方式防御
	enableDDoSDefence := os.Getenv("ENABLE_BROADCAST_DEFENCE")
	if enableDDoSDefence == "true" {
		el.EnableBroadcastDefence = true
		fmt.Println("enable broadcast defence")
	} else {
		el.EnableBroadcastDefence = false
		fmt.Println("disable broadcast defence")
	}

	// 2. 是否直接进行从 peer 中删除
	enableDirectRemove := os.Getenv("DIRECT_REMOVE")
	if enableDirectRemove == "true" {
		el.EnableDirectRemove = true
	} else {
		el.EnableDirectRemove = false
	}

	// 3. 获取 ddos 警告速率v
	ddosWarningRate := os.Getenv("DDOS_WARNING_RATE")
	ddosWarningRateFloat, err := strconv.ParseFloat(ddosWarningRate, 64)
	if err != nil {
		return fmt.Errorf("unable to parse DDOS_WARNING_RATE: %w", err)
	}
	el.DDoSWarningRate = ddosWarningRateFloat

	// 4. 获取是否进行 speed check
	speedCheck := os.Getenv("SPEED_CHECK")
	if speedCheck == "true" {
		el.SpeedCheck = true
	} else {
		el.SpeedCheck = false
	}

	// 5. 获取第一个接口的名称
	el.InterfaceName = os.Getenv("INTERFACE_NAME")

	// 6. 获取监听的地址
	el.ListenIpAddr = os.Getenv("LISTEN_ADDR")

	// 7. 判断是否启动 frr
	enableFrr := os.Getenv("ENABLE_FRR")
	if enableFrr == "true" {
		el.EnableFrr = true
	} else {
		el.EnableFrr = false
	}

	// 8. 获取容器名
	el.ContainerName = os.Getenv("CONTAINER_NAME")

	// 9. 获取 webServerListenPort 端口
	webServerListenPort, err := strconv.ParseInt(os.Getenv("WEB_SERVER_LISTEN_PORT"), 10, 64)
	if err != nil {
		return fmt.Errorf("parse WEB_SERVER_LISTEN_PORT error: %v", err)
	}
	el.WebServerListenPort = int(webServerListenPort)

	// 10. 获取 blocksPerProposer
	blocksPerProposerString := os.Getenv("BLOCKS_PER_PROPOSER")
	el.BlocksPerProposer, err = strconv.Atoi(blocksPerProposerString)
	if err != nil {
		return fmt.Errorf("unable to parse BLOCKS_PER_PROPOSER: %w", err)
	}

	// 11. TimeoutProposeOptimal
	timeoutProposeOptimalStr := os.Getenv("TIMEOUT_PROPOSE_OPTIMAL")
	timeoutProposeOptimalInt, err := strconv.Atoi(timeoutProposeOptimalStr)
	if err != nil {
		return fmt.Errorf("unable to parse TIMEOUT_PROPOSE_OPTIMAL")
	}
	el.TimeoutProposeOptimal = time.Millisecond * time.Duration(timeoutProposeOptimalInt)

	// 12. ProposeOptimal
	proposeOptimal := os.Getenv("PROPOSE_OPTIMAL")
	fmt.Printf("config loader load proposeOptimal == %v\n", proposeOptimal)
	if proposeOptimal == "true" {
		el.ProposeOptimal = true
	} else {
		el.ProposeOptimal = false
	}

	// 11. 获取 startDefenceKey
	el.StartDefenceKey = os.Getenv("START_DEFENCE_KEY")

	// 12. 获取 etcdListenAddr
	el.EtcdListenAddr = os.Getenv("ETCD_LISTEN_ADDR")

	// 13. 获取 etcdListenPort
	el.EtcdListenPort, err = strconv.Atoi(os.Getenv("ETCD_LISTEN_PORT"))
	if err != nil {
		return fmt.Errorf("unable to parse ETCD_LISTEN_PORT: %w", err)
	}

	// 14. 进行 peerId 到 containerName mapping 的加载
	err = el.LoadPeerIdToContainerNameMapping()
	if err != nil {
		return fmt.Errorf("unable to load peerIdToContainerNameMapping: %v", err)
	}

	return nil
}

func (el *EnvLoader) LoadPeerIdToContainerNameMapping() error {
	peerIdToContainerNamePath := fmt.Sprintf("%s/%s/%s",
		"/configuration", el.ContainerName, "peerIdToContainerName.txt")
	lines, err := fileop.ReadLines(peerIdToContainerNamePath)
	if err != nil {
		return fmt.Errorf("read lines failed")
	}
	for _, line := range lines {
		results := strings.Split(line, ",")
		el.PeerIdToContainerNameMapping[results[0]] = results[1]
	}
	return nil
}
