package plot

import (
	"fmt"
	"time"
)

func ProposeTimeout(round int, timeoutPropose, timeoutProposeDelta time.Duration) time.Duration {
	return time.Duration(
		timeoutPropose.Nanoseconds()+timeoutProposeDelta.Nanoseconds()*int64(round),
	) * time.Nanosecond
}

func CalculateCurrentTimeout(round int, height, validatorHeight uint64, timeoutPropose, timeoutProposeDelta time.Duration) time.Duration {
	fmt.Printf("enter another way")
	timeout := ProposeTimeout(round, timeoutPropose, timeoutProposeDelta)
	// 用当前高度减去主节点的高度，计算 t
	t := height - validatorHeight
	// 如果当前高度小于等于主节点高度，t 设为 1，避免除零
	if height <= validatorHeight {
		t = 1
	}
	// 超时时间按比例缩短
	timeout = time.Duration(uint64(timeout) / t)
	// 超时时间最小为 2 秒
	if timeout < (2 * time.Second) {
		timeout = 2 * time.Second
	}
	// 最终将结果进行返回
	return timeout
}

func PlotProposeTimeout() {
	timeoutPropose := 30 * time.Second
	timeoutProposeDelta := 1 * time.Second
	for round := 1; round < 100; round++ {
		proposeTimeoutValue := ProposeTimeout(round, timeoutPropose, timeoutProposeDelta)
		fmt.Printf("round: %d, proposeTimeout: %v\n", round, proposeTimeoutValue)
	}
}

func PlotTimeoutTrend() {
	round := 5
	var height uint64 = 1000
	var validatorHeight uint64 = 100
	timeoutPropose := 30 * time.Second
	timeoutProposeDelta := 1 * time.Second
	// 演示高度正在逐渐逼近的时候的时间是怎么进行变化的
	for i := validatorHeight; i < height; i++ {
		timeout := CalculateCurrentTimeout(round, height, validatorHeight, timeoutPropose, timeoutProposeDelta)
		fmt.Printf("validatorHeight: %d, height:%d, timeout: %v\n", i, height, timeout)
	}
}
