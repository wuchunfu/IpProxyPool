package config

//const defaultConfigFile = "conf/config.yaml"
//const defaultConfigDir = "../conf"

//var _ *viper.Viper
//var InitConfig Server

//func init() {
//	if !fileutil.PathExists(defaultConfigFile) {
//		logger.Errorf("no such config directory: %s", defaultConfigFile)
//	}
//	vip := viper.New()
//	//vip.AddConfigPath(defaultConfigDir)
//	vip.SetConfigFile(defaultConfigFile)
//	//vip.SetConfigName("config")
//	vip.SetConfigType("yaml")
//	err := vip.ReadInConfig()
//	if err != nil {
//		logger.Errorf("Failed to get config file!")
//		panic(err.Error())
//	}
//	vip.WatchConfig()
//	vip.OnConfigChange(func(e fsnotify.Event) {
//		logger.Infof("Config file changed: %s", e.Name)
//		if err := vip.Unmarshal(&InitConfig); err != nil {
//			logger.Errorf("Failed to resolve config file!")
//			panic(err.Error())
//		}
//	})
//	if err := vip.Unmarshal(&InitConfig); err != nil {
//		logger.Errorf("Failed to resolve config file!")
//		panic(err.Error())
//	}
//	_ = vip
//}

//type Server struct {
//	System   System   `mapstructure:"system" json:"system" yaml:"system"`
//	Database Database `mapstructure:"database" json:"database" yaml:"database"`
//	Log      Log      `mapstructure:"log" json:"log" yaml:"log"`
//}
//
//type System struct {
//	AppName        string `mapstructure:"appName" json:"appName" yaml:"appName"`
//	HttpAddr       string `mapstructure:"httpAddr" json:"httpAddr" yaml:"httpAddr"`
//	HttpPort       string `mapstructure:"httpPort" json:"httpPort" yaml:"httpPort"`
//	SessionExpires string `mapstructure:"sessionExpires" json:"sessionExpires" yaml:"sessionExpires"`
//}
//
//type Database struct {
//	DbType       string `mapstructure:"dbType" json:"dbType" yaml:"dbType"`
//	Host         string `mapstructure:"host" json:"host" yaml:"host"`
//	Port         int    `mapstructure:"port" json:"port" yaml:"port"`
//	DbName       string `mapstructure:"dbName" json:"dbName" yaml:"dbName"`
//	Username     string `mapstructure:"username" json:"username" yaml:"username"`
//	Password     string `mapstructure:"password" json:"password" yaml:"password"`
//	Prefix       string `mapstructure:"prefix" json:"prefix" yaml:"prefix"`
//	Charset      string `mapstructure:"charset" json:"charset" yaml:"charset"`
//	MaxIdleConns int    `mapstructure:"maxIdleConns" json:"maxIdleConns" yaml:"maxIdleConns"`
//	MaxOpenConns int    `mapstructure:"maxOpenConns" json:"maxOpenConns" yaml:"maxOpenConns"`
//	SslMode      string `mapstructure:"sslMode" json:"sslMode" yaml:"sslMode"`
//	Path         string `mapstructure:"path" json:"path" yaml:"path"`
//}
//
//type Log struct {
//	LogDirPath  string `mapstructure:"logDirPath" json:"logDirPath" yaml:"logDirPath"`
//	LogFileName string `mapstructure:"logFileName" json:"logFileName" yaml:"logFileName"`
//	LogLevel    string `mapstructure:"logLevel" json:"logLevel" yaml:"logLevel"`
//	Mode        string `mapstructure:"mode" json:"mode" yaml:"mode"`
//	Path        string `mapstructure:"path" json:"path" yaml:"path"`
//	Name        string `mapstructure:"name" json:"name" yaml:"name"`
//	Level       string `mapstructure:"level" json:"level" yaml:"level"`
//}
