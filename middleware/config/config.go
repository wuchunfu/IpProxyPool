package config

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	logger "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/wuchunfu/IpProxyPool/middleware/database"
	"github.com/wuchunfu/IpProxyPool/middleware/logutil"
	"github.com/wuchunfu/IpProxyPool/models/configModel"
	"github.com/wuchunfu/IpProxyPool/util/fileutil"
	"os"
)

var (
	Vip         = viper.New()
	ConfigFile  = ""
	YamlSetting = new(configModel.YamlSetting)
)

// InitConfig reads in config file and ENV variables if set.
func InitConfig() {
	if ConfigFile != "" {
		if !fileutil.PathExists(ConfigFile) {
			logger.Errorf("No such file or directory: %s", ConfigFile)
			os.Exit(-1)
		} else {
			// Use config file from the flag.
			Vip.SetConfigFile(ConfigFile)
			Vip.SetConfigType("yaml")
		}
	} else {
		logger.Errorf("Could not find config file: %s", ConfigFile)
		os.Exit(-1)
	}
	// If a config file is found, read it in.
	err := Vip.ReadInConfig()
	if err != nil {
		logger.Errorf("Failed to get config file: %s", ConfigFile)
	}
	Vip.WatchConfig()
	Vip.OnConfigChange(func(e fsnotify.Event) {
		logger.Infof("Config file changed: %s\n", e.Name)
		fmt.Printf("Config file changed: %s\n", e.Name)
		GetInitConfig(Vip)
	})
	Vip.AllSettings()
	GetInitConfig(Vip)
}

// 解析配置文件，反序列化
func parseYaml(vip *viper.Viper) {
	setting := new(configModel.YamlSetting)
	if err := vip.Unmarshal(setting); err != nil {
		logger.Errorf("Unmarshal yaml faild: %s", err)
		os.Exit(-1)
	}
	YamlSetting = setting
}

func GetInitConfig(vip *viper.Viper) {
	// 解析配置文件，反序列化
	parseYaml(vip)
	// 将日志写入文件或打印到控制台
	logutil.InitLog(&YamlSetting.Log)
	// 初始化数据库连接
	database.InitDB(&YamlSetting.Database)
}
