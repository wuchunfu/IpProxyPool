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

	setting := config.ServerSetting

	StartCmd.PersistentFlags().StringVarP(&config.ConfigFile, "configFile", "f", "conf/config.yaml", "config file")
	StartCmd.PersistentFlags().StringVarP(&setting.System.HttpAddr, "httpAddr", "a", "0.0.0.0", "http addr")
	StartCmd.PersistentFlags().StringVar(&setting.Database.DbType, "dbType", "mysql", "database type")
	StartCmd.PersistentFlags().StringVar(&setting.Database.Host, "host", "127.0.0.1", "database host")
	StartCmd.PersistentFlags().IntVarP(&setting.Database.Port, "port", "p", 3306, "database port")
	StartCmd.PersistentFlags().StringVar(&setting.Database.DbName, "dbName", "", "database name")
	StartCmd.PersistentFlags().StringVar(&setting.Database.Username, "username", "", "database username")
	StartCmd.PersistentFlags().StringVar(&setting.Database.Password, "password", "", "database password")
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
	config.Vip.SetDefault("system.httpAddr", "0.0.0.0")
	config.Vip.SetDefault("system.httpPort", "3000")
}
