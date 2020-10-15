package config

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/mitchellh/go-homedir"
	logger "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"proxypool-go/middleware/database"
	"proxypool-go/middleware/logutil"
	"proxypool-go/models/configModel"
	"proxypool-go/util/fileutil"
)

var (
	Vip           = viper.New()
	configFile    = ""
	SystemSetting = new(configModel.System)
	DbSetting     = new(configModel.Database)
	LogSetting    = new(configModel.Log)

	StartCmd = &cobra.Command{
		Use:          "config",
		SilenceUsage: true,
		Short:        "Get Application config info",
		Example:      "proxy-pool config -f conf/config.yml",
		Run: func(cmd *cobra.Command, args []string) {
		},
	}
)

func init() {
	cobra.OnInitialize(initConfig)

	StartCmd.PersistentFlags().StringVarP(&configFile, "configFile", "f", "conf/config.yaml", "config file")
	StartCmd.PersistentFlags().StringVarP(&SystemSetting.HttpAddr, "httpAddr", "a", "0.0.0.0", "http addr")
	StartCmd.PersistentFlags().StringVar(&DbSetting.DbType, "dbType", "mysql", "database type")
	StartCmd.PersistentFlags().StringVar(&DbSetting.Host, "host", "127.0.0.1", "database host")
	StartCmd.PersistentFlags().IntVarP(&DbSetting.Port, "port", "p", 3306, "database port")
	StartCmd.PersistentFlags().StringVar(&DbSetting.DbName, "dbName", "", "database name")
	StartCmd.PersistentFlags().StringVar(&DbSetting.Username, "username", "", "database username")
	StartCmd.PersistentFlags().StringVar(&DbSetting.Password, "password", "", "database password")
	//
	_ = StartCmd.MarkFlagRequired("configFile")

	// 使用viper可以绑定flag
	_ = Vip.BindPFlag("system.httpAddr", StartCmd.PersistentFlags().Lookup("httpAddr"))
	_ = Vip.BindPFlag("database.dbType", StartCmd.PersistentFlags().Lookup("dbType"))
	_ = Vip.BindPFlag("database.host", StartCmd.PersistentFlags().Lookup("host"))
	_ = Vip.BindPFlag("database.port", StartCmd.PersistentFlags().Lookup("port"))
	_ = Vip.BindPFlag("database.dbName", StartCmd.PersistentFlags().Lookup("dbName"))
	_ = Vip.BindPFlag("database.username", StartCmd.PersistentFlags().Lookup("username"))
	_ = Vip.BindPFlag("database.password", StartCmd.PersistentFlags().Lookup("password"))

	// 设置默认值
	Vip.SetDefault("system.appName", "")
	Vip.SetDefault("system.httpAddr", "0.0.0.0")
	Vip.SetDefault("system.httpPort", "3000")
	Vip.SetDefault("system.sessionExpires", "168h0m0s")

	Vip.SetDefault("database.dbType", "mysql")
	Vip.SetDefault("database.host", "127.0.0.1")
	Vip.SetDefault("database.port", 3306)
	Vip.SetDefault("database.dbName", "")
	Vip.SetDefault("database.username", "")
	Vip.SetDefault("database.password", "")
	Vip.SetDefault("database.prefix", "proxy_")
	Vip.SetDefault("database.charset", "utf8mb4")
	Vip.SetDefault("database.maxIdleConns", 5)
	Vip.SetDefault("database.maxOpenConns", 100)
	// For "postgres" only, either "disable", "require" or "verify-full"
	Vip.SetDefault("database.sslMode", "disable")
	// For "sqlite3" and "tidb", use absolute path when you start as service
	Vip.SetDefault("database.path", "data/ProxyPool.db")

	Vip.SetDefault("log.filePath", "logs")
	Vip.SetDefault("log.fileName", "run.log")
	Vip.SetDefault("log.level", "info")
	Vip.SetDefault("log.mode", "console")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if configFile != "" {
		if !fileutil.PathExists(configFile) {
			logger.Errorf("no such file or directory: %s", configFile)
			os.Exit(-1)
		} else {
			// Use config file from the flag.
			Vip.SetConfigFile(configFile)
			Vip.SetConfigType("yaml")
		}
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			logger.Errorf("no such file or directory: %s", configFile)
			os.Exit(-1)
		}

		// Search config in home directory with name ".newApp" (without extension).
		Vip.AddConfigPath(home)
		Vip.SetConfigType("yaml")
		Vip.SetConfigName(".proxypool.yaml")
	}
	// If a config file is found, read it in.
	err := Vip.ReadInConfig()
	if err != nil {
		logger.Errorf("no such file or directory: %s", configFile)
		logger.Errorf("Failed to get config file: %s", configFile)
	}
	Vip.WatchConfig()
	Vip.OnConfigChange(func(e fsnotify.Event) {
		logger.Infof("Config file changed: %s\n", e.Name)
		fmt.Printf("Config file changed: %s\n", e.Name)
		GetInitConfig(Vip)
	})
	GetInitConfig(Vip)
}

func GetParams(vip *viper.Viper) {
	// system
	SystemSetting.AppName = vip.GetString("system.appName")
	SystemSetting.HttpAddr = vip.GetString("system.httpAddr")
	SystemSetting.HttpPort = vip.GetString("system.httpPort")
	SystemSetting.SessionExpires = vip.GetString("system.sessionExpires")

	// database
	DbSetting.DbType = vip.GetString("database.dbType")
	DbSetting.Host = vip.GetString("database.host")
	DbSetting.Port = vip.GetInt("database.port")
	DbSetting.DbName = vip.GetString("database.dbName")
	DbSetting.Username = vip.GetString("database.username")
	DbSetting.Password = vip.GetString("database.password")
	DbSetting.Prefix = vip.GetString("database.prefix")
	DbSetting.Charset = vip.GetString("database.charset")
	DbSetting.MaxIdleConns = vip.GetInt("database.maxIdleConns")
	DbSetting.MaxOpenConns = vip.GetInt("database.maxOpenConns")
	DbSetting.SslMode = vip.GetString("database.sslMode")
	DbSetting.Path = vip.GetString("database.path")

	// log
	LogSetting.FilePath = vip.GetString("log.filePath")
	LogSetting.FileName = vip.GetString("log.fileName")
	LogSetting.Level = vip.GetString("log.level")
	LogSetting.Mode = vip.GetString("log.mode")
}

func GetInitConfig(vip *viper.Viper) {
	GetParams(vip)
	// 将日志写入文件或打印到控制台
	logutil.InitLog(LogSetting)
	// 初始化数据库连接
	database.InitDB(DbSetting)
}
