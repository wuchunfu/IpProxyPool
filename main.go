package main

import (
	"proxypool-go/api"
	"proxypool-go/cmd"
	"proxypool-go/middleware/config"
	"proxypool-go/run"
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
