package etcd

import (
	"context"
	"fmt"
	"github.com/zhanghefan123/chainmaker-zeusnet/modules/config"
	"github.com/zhanghefan123/chainmaker-zeusnet/modules/defence/broadcast"
)

// HandleStartDefenceChange 处理 startDefence 的变更
func HandleStartDefenceChange() {
	watchChan := Client.Watch(context.Background(),
		config.EnvLoaderInstance.StartDefenceKey)
	for response := range watchChan {
		for _, event := range response.Events {
			if string(event.Kv.Value) == "true" {
				fmt.Println(string(event.Kv.Value))
				broadcast.StartDefenceQueue <- true
			} else {
				broadcast.StartDefenceQueue <- false
			}
		}
	}
}
