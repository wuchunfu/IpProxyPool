package config

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/mitchellh/go-homedir"
	logger "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"proxypool-go/middleware/database"
	"proxypool-go/middleware/logutil"
	"proxypool-go/models/configModel"
	"proxypool-go/util/fileutil"
)

var (
	Vip           = viper.New()
	ConfigFile    = ""
	SystemSetting = new(configModel.System)
	DbSetting     = new(configModel.Database)
	LogSetting    = new(configModel.Log)
)

// InitConfig reads in config file and ENV variables if set.
func InitConfig() {
	if ConfigFile != "" {
		if !fileutil.PathExists(ConfigFile) {
			logger.Errorf("no such file or directory: %s", ConfigFile)
			os.Exit(-1)
		} else {
			// Use config file from the flag.
			Vip.SetConfigFile(ConfigFile)
			Vip.SetConfigType("yaml")
		}
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			logger.Errorf("no such file or directory: %s", ConfigFile)
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
		logger.Errorf("no such file or directory: %s", ConfigFile)
		logger.Errorf("Failed to get config file: %s", ConfigFile)
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
