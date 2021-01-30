package config

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	logger "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/wuchunfu/IpProxyPool/util/fileutil"
	"os"
)

type System struct {
	AppName  string `yaml:"appName"`
	HttpAddr string `yaml:"httpAddr"`
	HttpPort string `yaml:"httpPort"`
}

type Database struct {
	DbType       string `yaml:"dbType"`
	Host         string `yaml:"host"`
	Port         int    `yaml:"port"`
	DbName       string `yaml:"dbName"`
	Username     string `yaml:"username"`
	Password     string `yaml:"password"`
	Prefix       string `yaml:"prefix"`
	Charset      string `yaml:"charset"`
	MaxIdleConns int    `yaml:"maxIdleConns"`
	MaxOpenConns int    `yaml:"maxOpenConns"`
	Level        string `yaml:"level"`
	SslMode      string `yaml:"sslMode"`
	TimeZone     string `yaml:"timeZone"`
}

type Log struct {
	FilePath string `yaml:"filePath"`
	FileName string `yaml:"fileName"`
	Level    string `yaml:"level"`
	Mode     string `yaml:"mode"`
}

type YamlSetting struct {
	System   System   `yaml:"system"`
	Database Database `yaml:"database"`
	Log      Log      `yaml:"log"`
}

var (
	Vip           = viper.New()
	ConfigFile    = ""
	ServerSetting = new(YamlSetting)
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
		ServerSetting = GetConfig(Vip)
	})
	Vip.AllSettings()
	ServerSetting = GetConfig(Vip)
}

// 解析配置文件，反序列化
func GetConfig(vip *viper.Viper) *YamlSetting {
	setting := new(YamlSetting)
	// 解析配置文件，反序列化
	if err := vip.Unmarshal(setting); err != nil {
		logger.Errorf("Unmarshal yaml faild: %s", err)
		os.Exit(-1)
	}
	return setting
}
