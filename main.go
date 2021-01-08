package main

import (
	"github.com/wuchunfu/IpProxyPool/api"
	"github.com/wuchunfu/IpProxyPool/cmd"
	"github.com/wuchunfu/IpProxyPool/middleware/config"
	"github.com/wuchunfu/IpProxyPool/run"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	// 检查或设置命令行参数
	cmd.Execute()

	// Start HTTP
	go func() {
		api.Run(config.SystemSetting)
	}()

	// Start Task
	run.Task()
}
