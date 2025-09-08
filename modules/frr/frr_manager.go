package frr

import (
	"bytes"
	"fmt"
	"os/exec"
)

// CopyFrrConfigurationFile 拷贝 frr 配置文件
func CopyFrrConfigurationFile(containerName string) {
	sourceFilePath := fmt.Sprintf("/configuration/%s/route/frr.conf", containerName)

	destFilePath := "/etc/frr/frr.conf"

	cmd := exec.Command("cp", sourceFilePath, destFilePath)
	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
	}
}

// StartFrr 开启 frr 服务
func StartFrr() {
	cmd := exec.Command("service", "frr", "start")
	output, err := cmd.CombinedOutput()
	if err != nil {
		panic(err)
	}
	fmt.Println(string(output))
}

// Start 启动 frr
func Start(containerName string) {
	go func() {
		fmt.Println("start frr")
		// 拷贝 frr 配置文件
		CopyFrrConfigurationFile(containerName)
		// 执行启动命令
		StartFrr()
	}()
}
