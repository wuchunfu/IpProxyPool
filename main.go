package main

import (
	"github.com/wuchunfu/IpProxyPool/api"
	"github.com/wuchunfu/IpProxyPool/cmd"
	"github.com/wuchunfu/IpProxyPool/middleware/config"
	"github.com/wuchunfu/IpProxyPool/middleware/database"
	"github.com/wuchunfu/IpProxyPool/middleware/logutil"
	"github.com/wuchunfu/IpProxyPool/run"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	// 检查或设置命令行参数
	cmd.Execute()

	setting := config.ServerSetting

	// 将日志写入文件或打印到控制台
	logutil.InitLog(&setting.Log)
	// 初始化数据库连接
	database.InitDB(&setting.Database)

	// Start HTTP
	go func() {
		api.Run(&setting.System)
	}()

	// Start Task
	run.Task()
}
