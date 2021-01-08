package server

import (
	"github.com/spf13/cobra"
	"github.com/wuchunfu/IpProxyPool/middleware/config"
)

var StartCmd = &cobra.Command{
	Use:          "config",
	SilenceUsage: true,
	Short:        "Get Application config info",
	Example:      "proxy-pool config -f conf/config.yml",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	cobra.OnInitialize(config.InitConfig)

	StartCmd.PersistentFlags().StringVarP(&config.ConfigFile, "configFile", "f", "conf/config.yaml", "config file")
	StartCmd.PersistentFlags().StringVarP(&config.SystemSetting.HttpAddr, "httpAddr", "a", "0.0.0.0", "http addr")
	StartCmd.PersistentFlags().StringVar(&config.DbSetting.DbType, "dbType", "mysql", "database type")
	StartCmd.PersistentFlags().StringVar(&config.DbSetting.Host, "host", "127.0.0.1", "database host")
	StartCmd.PersistentFlags().IntVarP(&config.DbSetting.Port, "port", "p", 3306, "database port")
	StartCmd.PersistentFlags().StringVar(&config.DbSetting.DbName, "dbName", "", "database name")
	StartCmd.PersistentFlags().StringVar(&config.DbSetting.Username, "username", "", "database username")
	StartCmd.PersistentFlags().StringVar(&config.DbSetting.Password, "password", "", "database password")
	// 必须配置项
	_ = StartCmd.MarkFlagRequired("configFile")

	// 使用viper可以绑定flag
	_ = config.Vip.BindPFlag("system.httpAddr", StartCmd.PersistentFlags().Lookup("httpAddr"))
	_ = config.Vip.BindPFlag("database.dbType", StartCmd.PersistentFlags().Lookup("dbType"))
	_ = config.Vip.BindPFlag("database.host", StartCmd.PersistentFlags().Lookup("host"))
	_ = config.Vip.BindPFlag("database.port", StartCmd.PersistentFlags().Lookup("port"))
	_ = config.Vip.BindPFlag("database.dbName", StartCmd.PersistentFlags().Lookup("dbName"))
	_ = config.Vip.BindPFlag("database.username", StartCmd.PersistentFlags().Lookup("username"))
	_ = config.Vip.BindPFlag("database.password", StartCmd.PersistentFlags().Lookup("password"))

	// 设置默认值
	config.Vip.SetDefault("system.appName", "")
	config.Vip.SetDefault("system.httpAddr", "0.0.0.0")
	config.Vip.SetDefault("system.httpPort", "3000")
	config.Vip.SetDefault("system.sessionExpires", "168h0m0s")

	config.Vip.SetDefault("database.dbType", "mysql")
	config.Vip.SetDefault("database.host", "127.0.0.1")
	config.Vip.SetDefault("database.port", 3306)
	config.Vip.SetDefault("database.dbName", "")
	config.Vip.SetDefault("database.username", "")
	config.Vip.SetDefault("database.password", "")
	config.Vip.SetDefault("database.prefix", "proxy_")
	config.Vip.SetDefault("database.charset", "utf8mb4")
	config.Vip.SetDefault("database.maxIdleConns", 5)
	config.Vip.SetDefault("database.maxOpenConns", 100)
	config.Vip.SetDefault("database.level", "silent")
	config.Vip.SetDefault("database.sslMode", "disable")
	config.Vip.SetDefault("database.timeZone", "Asia/Shanghai")

	config.Vip.SetDefault("log.filePath", "logs")
	config.Vip.SetDefault("log.fileName", "run.log")
	config.Vip.SetDefault("log.level", "info")
	config.Vip.SetDefault("log.mode", "console")
}
