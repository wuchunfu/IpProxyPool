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
//	"proxypool-go/util/fileutil"
//)
//
//var (
//	AppName    = ""
//	configFile = ""
//	// service listening address
//	HttpAddr = "0.0.0.0"
//	// service listening port
//	HttpPort = "3000"
//	// Session expires time
//	SessionExpires = "168h0m0s"
//
//	DbType       = "mysql"
//	Host         = "127.0.0.1"
//	Port         = 3306
//	DbName       = ""
//	Username     = ""
//	Password     = ""
//	Prefix       = "proxy_"
//	Charset      = "utf8mb4"
//	MaxIdleConns = 5
//	MaxOpenConns = 100
//	// For "postgres" only, either "disable", "require" or "verify-full"
//	SslMode = "disable"
//	// For "sqlite3" and "tidb", use absolute path when you start as service
//	Path = "data/ProxyPool.db"
//
//	FilePath = "logs"
//	FileName = "run.log"
//	Level    = "info"
//	Mode    = "console"
//)
//
//var vip = viper.New()
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
//	rootCmd.PersistentFlags().StringVarP(&HttpAddr, "httpAddr", "a", "0.0.0.0", "http addr")
//	rootCmd.PersistentFlags().StringVarP(&HttpPort, "httpPort", "P", "3000", "http port")
//	rootCmd.PersistentFlags().StringVar(&DbType, "dbType", "mysql", "database type")
//	rootCmd.PersistentFlags().StringVar(&Host, "host", "127.0.0.1", "database host")
//	rootCmd.PersistentFlags().IntVarP(&Port, "port", "p", 3306, "database port")
//	rootCmd.PersistentFlags().StringVar(&DbName, "dbName", "", "database name")
//	rootCmd.PersistentFlags().StringVar(&Username, "username", "", "database username")
//	rootCmd.PersistentFlags().StringVar(&Password, "password", "", "database password")
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
//	vip.SetDefault("system.appName", AppName)
//	vip.SetDefault("system.httpAddr", HttpAddr)
//	vip.SetDefault("system.httpPort", HttpPort)
//	vip.SetDefault("system.sessionExpires", SessionExpires)
//
//	vip.SetDefault("database.dbType", DbType)
//	vip.SetDefault("database.host", Host)
//	vip.SetDefault("database.port", Port)
//	vip.SetDefault("database.dbName", DbName)
//	vip.SetDefault("database.username", Username)
//	vip.SetDefault("database.password", Password)
//	vip.SetDefault("database.prefix", Prefix)
//	vip.SetDefault("database.charset", Charset)
//	vip.SetDefault("database.maxIdleConns", MaxIdleConns)
//	vip.SetDefault("database.maxOpenConns", MaxOpenConns)
//	vip.SetDefault("database.sslMode", SslMode)
//	vip.SetDefault("database.path", Path)
//
//	vip.SetDefault("log.filePath", FilePath)
//	vip.SetDefault("log.fileName", FileName)
//	vip.SetDefault("log.level", Level)
//	vip.SetDefault("log.mode", Mode)
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
//	})
//	GetParams(vip)
//}
//
//func GetParams(vip *viper.Viper) {
//	fmt.Println("==================")
//	// system
//	AppName = vip.GetString("system.appName")
//	configFile = vip.GetString("system.configFile")
//	HttpAddr = vip.GetString("system.httpAddr")
//	HttpPort = vip.GetString("system.httpPort")
//	SessionExpires = vip.GetString("system.sessionExpires")
//
//	// database
//	DbType = vip.GetString("database.dbType")
//	Host = vip.GetString("database.host")
//	Port = vip.GetInt("database.port")
//	DbName = vip.GetString("database.dbName")
//	Username = vip.GetString("database.username")
//	Password = vip.GetString("database.password")
//	fmt.Println("===Password===")
//	fmt.Println(Password)
//	Prefix = vip.GetString("database.prefix")
//	Charset = vip.GetString("database.charset")
//	MaxIdleConns = vip.GetInt("database.maxIdleConns")
//	MaxOpenConns = vip.GetInt("database.maxOpenConns")
//	SslMode = vip.GetString("database.sslMode")
//	Path = vip.GetString("database.path")
//
//	// log
//	FilePath = vip.GetString("log.filePath")
//	FileName = vip.GetString("log.fileName")
//	Level = vip.GetString("log.level")
//	Mode = vip.GetString("log.mode")
//}
