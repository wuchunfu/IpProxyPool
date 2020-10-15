package database

import (
	"database/sql"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"github.com/sirupsen/logrus"
	//_ "github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
	"os"
	"proxypool-go/models/configModel"
	"strings"
	"time"
)

var dbPingInterval = 90 * time.Second
var DB *gorm.DB

func InitDB(setting *configModel.Database) *gorm.DB {
	//setting := new(config.Database)
	dsn := getDbEngineDSN(setting)
	//db, err := gorm.Open(setting.DbType, dsn)

	//newLogger := logger.New(
	//	log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
	//	logger.Config{
	//		SlowThreshold: time.Second, // 慢 SQL 阈值
	//		LogLevel:      logger.Info, // Log level
	//		Colorful:      false,       // 禁用彩色打印
	//	},
	//)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   setting.Prefix, // 表名前缀，`User` 的表名应该是 `t_users`
			SingularTable: true,           // 使用单数表名，启用该选项，此时，`User` 的表名应该是 `t_user`
		},
		PrepareStmt:            true, // 执行任何 SQL 时都创建并缓存预编译语句，可以提高后续的调用速度
		DisableAutomaticPing:   false,
		SkipDefaultTransaction: true, // 对于写操作（创建、更新、删除），为了确保数据的完整性，GORM 会将它们封装在事务内运行。但这会降低性能，你可以在初始化时禁用这种方式
		Logger:                 logger.Default.LogMode(logger.Info),
		//Logger:            newLogger,
		AllowGlobalUpdate: false,
	})
	if err != nil {
		logrus.Errorf("fail to connect database: %v\n", err)
		os.Exit(-1)
	}
	sqlDb, dbErr := db.DB()
	if dbErr != nil {
		logrus.Errorf("fail to connect database: %v\n", dbErr)
		os.Exit(-1)
	}
	// 设置连接池
	// 用于设置连接池中空闲连接的最大数量。
	sqlDb.SetMaxIdleConns(10)
	// 设置打开数据库连接的最大数量
	sqlDb.SetMaxOpenConns(100)
	//db.LogMode(true)
	//sqlDb.LogMode(false)

	// * 解决中文字符问题：Error 1366
	db.Set("gorm:table_options", "ENGINE=InnoDB CHARSET=utf8mb4")
	go keepDbAlived(sqlDb)
	DB = db
	return db
}

func GetDB() *gorm.DB {
	return DB
}

// 获取数据库引擎DSN  mysql,postgres
func getDbEngineDSN(setting *configModel.Database) string {
	//setting := new(config.Database)
	engine := strings.ToLower(setting.DbType)
	dsn := ""
	switch engine {
	case "mysql":
		// parseTime: 想要能正确的处理 time.Time，你需要添加 parseTime 参数。
		// loc: 设置时间的位置
		dsn = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&allowNativePasswords=true&parseTime=True&loc=Local",
			// 连接数据库的用户名
			setting.Username,
			// 连接数据库的密码
			setting.Password,
			// 连接数据库的地址
			setting.Host,
			// 连接数据库的端口号
			setting.Port,
			// 连接数据库的具体数据库名称
			setting.DbName,
			// 连接数据库的编码格式
			setting.Charset)
	case "postgres":
		dsn = fmt.Sprintf("host=%s port=%d dbname=%s user=%s password=%s sslmode=disable TimeZone=Asia/Shanghai",
			setting.Host,
			setting.Port,
			setting.DbName,
			setting.Username,
			setting.Password)
	//dsn = fmt.Sprintf("postgres://%s:%s@%s:%s/%ssslmode=%s",
	//	url.QueryEscape(setting.Username),
	//	url.QueryEscape(setting.Password),
	//	setting.Host,
	//	setting.Port,
	//	setting.DbName,
	//	setting.SslMode)
	default:
		return fmt.Sprintf("Unknown database type: %s", setting.DbType)
	}
	return dsn
}

// parsePostgreSQLHostPort parses given input in various forms defined in
// https://www.postgresql.org/docs/current/static/libpq-connect.html#LIBPQ-CONNSTRING
// and returns proper host and port number.
func parsePostgreSQLHostPort(info string) (string, string) {
	host, port := "127.0.0.1", "5432"
	if strings.Contains(info, ":") && !strings.HasSuffix(info, "]") {
		idx := strings.LastIndex(info, ":")
		host = info[:idx]
		port = info[idx+1:]
	} else if len(info) > 0 {
		host = info
	}
	return host, port
}

func parseMSSQLHostPort(info string) (string, string) {
	host, port := "127.0.0.1", "1433"
	if strings.Contains(info, ":") {
		host = strings.Split(info, ":")[0]
		port = strings.Split(info, ":")[1]
	} else if strings.Contains(info, ",") {
		host = strings.Split(info, ",")[0]
		port = strings.TrimSpace(strings.Split(info, ",")[1])
	} else if len(info) > 0 {
		host = info
	}
	return host, port
}

func keepDbAlived(engine *sql.DB) {
	t := time.Tick(dbPingInterval)
	var err error
	for {
		<-t
		err = engine.Ping()
		if err != nil {
			logrus.Infof("database ping: %s\n", err)
		}
	}
}
