package config
//
//import (
//	"fmt"
//	"github.com/fsnotify/fsnotify"
//	"github.com/mitchellh/go-homedir"
//	logger "github.com/sirupsen/logrus"
//	"github.com/spf13/cobra"
//	"github.com/spf13/viper"
//	"os"
//	"proxypool-go/api"
//	"proxypool-go/middleware/database"
//	"proxypool-go/middleware/logutil"
//	"proxypool-go/models/configModel"
//	"proxypool-go/util/fileutil"
//)
//
//var vip = viper.New()
//
//var configFile = ""
//var systemSetting = new(configModel.System)
//var dbSetting = new(configModel.Database)
//var logSetting = new(configModel.Log)
//
//// rootCmd represents the base command when called without any subcommands
//var rootCmd = &cobra.Command{
//	Use:   "proxypool-go",
//	Short: "A brief description of your application",
//	Long: `A longer description that spans multiple lines and likely contains
//examples and usage of using your application. For example:
//
//Cobra is a CLI library for Go that empowers applications.
//This application is a tool to generate the needed files
//to quickly create a Cobra application.`,
//	// Uncomment the following line if your bare application
//	// has an action associated with it:
//	Run: func(cmd *cobra.Command, args []string) {
//		fmt.Println("Run Done.")
//	},
//}
//
//// Execute adds all child commands to the root command and sets flags appropriately.
//// This is called by main.main(). It only needs to happen once to the rootCmd.
//func Execute() {
//	if err := rootCmd.Execute(); err != nil {
//		fmt.Println(err)
//		os.Exit(-1)
//	}
//}
//
//func init() {
//	cobra.OnInitialize(InitConfig)
//
//	rootCmd.PersistentFlags().StringVarP(&configFile, "configFile", "f", "conf/config.yaml", "config file")
//	rootCmd.PersistentFlags().StringVarP(&systemSetting.HttpAddr, "httpAddr", "a", "0.0.0.0", "http addr")
//	rootCmd.PersistentFlags().StringVarP(&systemSetting.HttpPort, "httpPort", "P", "3000", "http port")
//	rootCmd.PersistentFlags().StringVar(&dbSetting.DbType, "dbType", "mysql", "database type")
//	rootCmd.PersistentFlags().StringVar(&dbSetting.Host, "host", "127.0.0.1", "database host")
//	rootCmd.PersistentFlags().IntVarP(&dbSetting.Port, "port", "p", 3306, "database port")
//	rootCmd.PersistentFlags().StringVar(&dbSetting.DbName, "dbName", "", "database name")
//	rootCmd.PersistentFlags().StringVar(&dbSetting.Username, "username", "", "database username")
//	rootCmd.PersistentFlags().StringVar(&dbSetting.Password, "password", "", "database password")
//	//
//	_ = rootCmd.MarkFlagRequired("configFile")
//
//	// 使用viper可以绑定flag
//	_ = vip.BindPFlag("system.httpAddr", rootCmd.PersistentFlags().Lookup("httpAddr"))
//	_ = vip.BindPFlag("system.httpPort", rootCmd.PersistentFlags().Lookup("httpPort"))
//	_ = vip.BindPFlag("database.dbType", rootCmd.PersistentFlags().Lookup("dbType"))
//	_ = vip.BindPFlag("database.host", rootCmd.PersistentFlags().Lookup("host"))
//	_ = vip.BindPFlag("database.port", rootCmd.PersistentFlags().Lookup("port"))
//	_ = vip.BindPFlag("database.dbName", rootCmd.PersistentFlags().Lookup("dbName"))
//	_ = vip.BindPFlag("database.username", rootCmd.PersistentFlags().Lookup("username"))
//	_ = vip.BindPFlag("database.password", rootCmd.PersistentFlags().Lookup("password"))
//
//	// 设置默认值
//	vip.SetDefault("system.appName", "")
//	vip.SetDefault("system.httpAddr", "0.0.0.0")
//	vip.SetDefault("system.httpPort", "3000")
//	vip.SetDefault("system.sessionExpires", "168h0m0s")
//
//	vip.SetDefault("database.dbType", "mysql")
//	vip.SetDefault("database.host", "127.0.0.1")
//	vip.SetDefault("database.port", 3306)
//	vip.SetDefault("database.dbName", "")
//	vip.SetDefault("database.username", "")
//	vip.SetDefault("database.password", "")
//	vip.SetDefault("database.prefix", "proxy_")
//	vip.SetDefault("database.charset", "utf8mb4")
//	vip.SetDefault("database.maxIdleConns", 5)
//	vip.SetDefault("database.maxOpenConns", 100)
//	// For "postgres" only, either "disable", "require" or "verify-full"
//	vip.SetDefault("database.sslMode", "disable")
//	// For "sqlite3" and "tidb", use absolute path when you start as service
//	vip.SetDefault("database.path", "data/ProxyPool.db")
//
//	vip.SetDefault("log.filePath", "logs")
//	vip.SetDefault("log.fileName", "run.log")
//	vip.SetDefault("log.level", "info")
//	vip.SetDefault("log.mode", "console")
//}
//
//// initConfig reads in config file and ENV variables if set.
//func InitConfig() {
//	if configFile != "" {
//		if !fileutil.PathExists(configFile) {
//			logger.Errorf("no such file or directory: %s", configFile)
//			os.Exit(-1)
//		} else {
//			// Use config file from the flag.
//			vip.SetConfigFile(configFile)
//			vip.SetConfigType("yaml")
//		}
//	} else {
//		// Find home directory.
//		home, err := homedir.Dir()
//		fmt.Println(home)
//		if err != nil {
//			fmt.Println(err)
//			os.Exit(-1)
//		}
//
//		// Search config in home directory with name ".newApp" (without extension).
//		vip.AddConfigPath(home)
//		vip.SetConfigType("yaml")
//		vip.SetConfigName(".proxypool.yaml")
//	}
//	// If a config file is found, read it in.
//	err := vip.ReadInConfig()
//	if err != nil {
//		logger.Errorf("no such file or directory: %s", configFile)
//		logger.Errorf("Failed to get config file: %s", configFile)
//	}
//	vip.WatchConfig()
//	vip.OnConfigChange(func(e fsnotify.Event) {
//		logger.Infof("Config file changed: %s\n", e.Name)
//		fmt.Printf("Config file changed: %s\n", e.Name)
//		GetParams(vip)
//		logutil.InitLog(logSetting)
//		database.InitDB(dbSetting)
//	})
//	GetParams(vip)
//	// 将日志写入文件或打印到控制台
//	logutil.InitLog(logSetting)
//	// 初始化数据库连接
//	database.InitDB(dbSetting)
//	// Start HTTP
//	go func() {
//		api.Run(systemSetting)
//	}()
//}
//
//func GetParams(vip *viper.Viper) {
//	// system
//	systemSetting.AppName = vip.GetString("system.appName")
//	systemSetting.HttpAddr = vip.GetString("system.httpAddr")
//	systemSetting.HttpPort = vip.GetString("system.httpPort")
//	systemSetting.SessionExpires = vip.GetString("system.sessionExpires")
//
//	// database
//	dbSetting.DbType = vip.GetString("database.dbType")
//	dbSetting.Host = vip.GetString("database.host")
//	dbSetting.Port = vip.GetInt("database.port")
//	dbSetting.DbName = vip.GetString("database.dbName")
//	dbSetting.Username = vip.GetString("database.username")
//	dbSetting.Password = vip.GetString("database.password")
//	dbSetting.Prefix = vip.GetString("database.prefix")
//	dbSetting.Charset = vip.GetString("database.charset")
//	dbSetting.MaxIdleConns = vip.GetInt("database.maxIdleConns")
//	dbSetting.MaxOpenConns = vip.GetInt("database.maxOpenConns")
//	dbSetting.SslMode = vip.GetString("database.sslMode")
//	dbSetting.Path = vip.GetString("database.path")
//
//	// log
//	logSetting.FilePath = vip.GetString("log.filePath")
//	logSetting.FileName = vip.GetString("log.fileName")
//	logSetting.Level = vip.GetString("log.level")
//	logSetting.Mode = vip.GetString("log.mode")
//}
