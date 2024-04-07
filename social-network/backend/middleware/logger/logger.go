package logger

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"runtime/debug"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

const (
	Logfilepath = "./middleware/runtime/log"
)

func init() {
	logrus.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
	logrus.SetReportCaller(false)
}
func Write(msg string, filename string) {
	setOutPutFile(logrus.InfoLevel, filename)
	logrus.Info(msg)
}

func Debug(fields logrus.Fields, args ...interface{}) {
	setOutPutFile(logrus.DebugLevel, "debug")
	logrus.WithFields(fields).Debug(args...)
}

func Info(fields logrus.Fields, args ...interface{}) {
	setOutPutFile(logrus.InfoLevel, "info")
	logrus.WithFields(fields).Debug(args...)
}
func Warn(fields logrus.Fields, args ...interface{}) {
	setOutPutFile(logrus.WarnLevel, "warn")
	logrus.WithFields(fields).Debug(args...)
}

func Fatal(fields logrus.Fields, args ...interface{}) {
	setOutPutFile(logrus.FatalLevel, "fatal")
	logrus.WithFields(fields).Debug(args...)
}
func Error(fields logrus.Fields, args ...interface{}) {
	setOutPutFile(logrus.ErrorLevel, "error")
	logrus.WithFields(fields).Debug(args...)
}

func Panic(fields logrus.Fields, args ...interface{}) {
	setOutPutFile(logrus.PanicLevel, "panic")
	logrus.WithFields(fields).Debug(args...)
}
func Trace(fields logrus.Fields, args ...interface{}) {
	setOutPutFile(logrus.TraceLevel, "trace")
	logrus.WithFields(fields).Debug(args...)
}
func setOutPutFile(level logrus.Level, logName string) {

	if _, err := os.Stat(Logfilepath); os.IsNotExist(err) {
		err := os.MkdirAll(Logfilepath, 0777)
		if err != nil {
			panic(fmt.Errorf("创建日志 %s 出错：%s /n", Logfilepath, err))
		}
	}

	timestr := time.Now().Format("2006-01-02")
	fileName := path.Join(Logfilepath, logName+"_"+timestr+".log")

	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)

	if err != nil {
		fmt.Println("打开日志文件出错", err)
	}
	logrus.SetOutput(file)
	logrus.SetLevel(level)
}
func LoggerToFile() gin.LoggerConfig {
	if _, err := os.Stat(Logfilepath); os.IsNotExist(err) {
		err := os.MkdirAll(Logfilepath, 0777)
		if err != nil {
			panic(fmt.Errorf("创建日志 %s 出错：%s /n", Logfilepath, err))
		}
	}

	timeStr := time.Now().Format("2006-01-02")
	fileName := path.Join(Logfilepath, "success_"+timeStr+".log")

	file, _ := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)

	var conf = gin.LoggerConfig{
		Formatter: func(params gin.LogFormatterParams) string {
			return fmt.Sprintf("%s - %s \"%s %s %s %d %s \"%s\" %s\"\n",
				params.TimeStamp.Format("2006-01-02 15:04:05"),
				params.ClientIP,
				params.Method,
				params.Path,
				params.Request.Proto,
				params.StatusCode,
				params.Latency,
				params.Request.UserAgent(),
				params.ErrorMessage,
			)

		},
		Output: io.MultiWriter(os.Stdout, file),
	}

	return conf
}

func Recover(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			if _, errDir := os.Stat("./runtime/log"); os.IsNotExist(errDir) {
				errDir = os.MkdirAll("./runtime/log", 0777)
				if errDir != nil {
					panic(fmt.Errorf("创建日志路径 '%s' 出错: '%s", "./runtime/log", errDir))
				}
			}
			timeStr := time.Now().Format("2006-01-02")
			fileName := path.Join("./runtime/log", "error_"+timeStr+".log")
			f, errFile := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)

			if errFile != nil {
				fmt.Println(errFile)
			}
			timeFileStr := time.Now().Format("2006-01-02 15:04:05")
			f.WriteString("painc error time:" + timeFileStr + "\n")
			f.WriteString(fmt.Sprintf("%v", err) + "\n")
			f.WriteString("stacktrace from panic:" + string(debug.Stack()) + "\n")
			f.Close()
			c.JSON(http.StatusOK, gin.H{
				"code": 500,
				"msg":  fmt.Sprintf("%v", err),
			})
			c.Abort()

		}

	}()
	c.Next()

}
