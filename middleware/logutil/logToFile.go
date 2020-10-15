package logutil
//
//import (
//	"bufio"
//	"fmt"
//	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
//	"github.com/rifflock/lfshook"
//	"github.com/sirupsen/logrus"
//	"io"
//	"os"
//	"path"
//	"proxypool-go/util/fileutil"
//	"time"
//)
//
//// 日志记录到文件
//func LoggerToFile() gin.HandlerFunc {
//
//	logFilePath := ""
//	logFileName := ""
//
//	if !fileutil.PathExists(logFilePath) {
//		if err := os.MkdirAll(logFilePath, 0755); err != nil {
//			panic(err)
//		}
//	}
//
//
//
//	// 日志文件
//	fileName := path.Join(logFilePath, logFileName)
//
//	// 写入文件
//	src, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
//	if err != nil {
//		fmt.Println("err", err)
//	}
//
//	// 实例化
//	logger := logrus.New()
//
//	// 设置输出
//	logger.Out = src
//
//	// 设置日志级别
//	logger.SetLevel(logrus.DebugLevel)
//
//	// 设置 rotatelogs
//	logWriter, err := rotatelogs.New(
//		// 分割后的文件名称
//		fileName+".%Y%m%d.log",
//		// 生成软链，指向最新日志文件
//		rotatelogs.WithLinkName(fileName),
//		// 设置最大保存时间(7天)
//		rotatelogs.WithMaxAge(7*24*time.Hour),
//		// 设置日志切割时间间隔(1天)
//		rotatelogs.WithRotationTime(24*time.Hour),
//	)
//
//	writeMap := lfshook.WriterMap{
//		logrus.InfoLevel:  logWriter,
//		logrus.FatalLevel: logWriter,
//		logrus.DebugLevel: logWriter,
//		logrus.WarnLevel:  logWriter,
//		logrus.ErrorLevel: logWriter,
//		logrus.PanicLevel: logWriter,
//	}
//
//	lfHook := lfshook.NewHook(writeMap, &logrus.JSONFormatter{
//		TimestampFormat: "2006-01-02 15:04:05",
//	})
//
//	// 新增 Hook
//	logger.AddHook(lfHook)
//
//	return func(c *gin.Context) {
//		// 开始时间
//		startTime := time.Now()
//
//		// 处理请求
//		c.Next()
//
//		// 结束时间
//		endTime := time.Now()
//
//		// 执行时间
//		latencyTime := endTime.Sub(startTime)
//
//		// 请求方式
//		reqMethod := c.Request.Method
//
//		// 请求路由
//		reqUri := c.Request.RequestURI
//
//		// 状态码
//		statusCode := c.Writer.Status()
//
//		// 请求IP
//		clientIP := c.ClientIP()
//
//		// 日志格式
//		logger.WithFields(logrus.Fields{
//			"status_code":  statusCode,
//			"latency_time": latencyTime,
//			"client_ip":    clientIP,
//			"req_method":   reqMethod,
//			"req_uri":      reqUri,
//		}).Info()
//	}
//}
//
//func newLfsHook(fileName string) logrus.Hook {
//	// 设置 rotatelogs
//	logWriter, err := rotatelogs.New(
//		// 分割后的文件名称
//		fileName+".%Y%m%d.log",
//		// 生成软链，指向最新日志文件, WithLinkName为最新的日志建立软连接,以方便随着找到当前日志文件
//		rotatelogs.WithLinkName(fileName),
//		// WithMaxAge和WithRotationCount二者只能设置一个,
//		// 设置文件清理前的最长保存时间, 设置最大保存时间(7天)
//		rotatelogs.WithMaxAge(7*24*time.Hour),
//		// 设置文件清理前最多保存的个数
//		//rotatelogs.WithRotationCount(5),
//		// 设置日志分割的时间,这里设置为一天分割一次, 设置日志切割时间间隔(1天)
//		rotatelogs.WithRotationTime(24*time.Hour),
//	)
//
//	if err != nil {
//		logrus.Errorf("config local file system for logger error: %v", err)
//	}
//	multiWriter := io.MultiWriter(logWriter, os.Stdout)
//	//将函数名和行数放在日志里面
//	logrus.SetReportCaller(true)
//	logrus.SetFormatter(&logrus.JSONFormatter{})
//	logrus.SetOutput(multiWriter)
//
//	switch level := GlobalConfig.LogConf.LogLevel; level {
//	// 如果日志级别不是debug就不要打印日志到控制台了
//	case "debug":
//		logrus.SetLevel(logrus.DebugLevel)
//		logrus.SetOutput(os.Stderr)
//	case "info":
//		setNull()
//		logrus.SetLevel(logrus.InfoLevel)
//	case "warn":
//		setNull()
//		logrus.SetLevel(logrus.WarnLevel)
//	case "error":
//		setNull()
//		logrus.SetLevel(logrus.ErrorLevel)
//	default:
//		setNull()
//		logrus.SetLevel(logrus.InfoLevel)
//	}
//
//	writeMap := lfshook.WriterMap{
//		logrus.InfoLevel:  logWriter,
//		logrus.FatalLevel: logWriter,
//		logrus.DebugLevel: logWriter,
//		logrus.WarnLevel:  logWriter,
//		logrus.ErrorLevel: logWriter,
//		logrus.PanicLevel: logWriter,
//	}
//
//	lfsHook := lfshook.NewHook(writeMap, &logrus.JSONFormatter{
//		TimestampFormat: "2006-01-02 15:04:05",
//	})
//
//	return lfsHook
//}

//func setNull() {
//	src, err := os.OpenFile(os.DevNull, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
//	if err != nil {
//		fmt.Println("err", err)
//	}
//	writer := bufio.NewWriter(src)
//	logrus.SetOutput(writer)
//}
