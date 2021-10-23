package logging

import (
	"gin-template/core/util"
	"io"
	"os"
	"path"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"github.com/spf13/viper"
)

var logFile = logrus.New()
var logConsole = logrus.New()
var enableFileLog = false

func Setup() {
	enableFileLog = viper.GetBool("app.log.file.enable")
	logConsole.SetFormatter(&logrus.TextFormatter{
		ForceColors:     true,
		TimestampFormat: "YYYY",
	})

	// 是否打印日志到文件中
	if enableFileLog {
		// 设置gin server的日志输出
		serverFile := CreateLogFile("server.log")
		gin.DisableConsoleColor()
		gin.DefaultWriter = io.MultiWriter(os.Stdout, serverFile)

		// 设置application的日志输出
		appFile := CreateLogFile("app.log")
		logFile.SetFormatter(&logrus.JSONFormatter{})
		level, err := logrus.ParseLevel(viper.GetString("app.log.file.level"))
		if err != nil {
			level = logrus.InfoLevel
		}
		logFile.SetLevel(level)
		logFile.SetOutput(io.MultiWriter(appFile))
	}

}

func CreateLogFile(fileName string) *os.File {
	logPath := viper.GetString("app.log.file.path")
	if err := util.MkdirIfNotExist(logPath); err != nil {
		Error(err)
	}
	f, err := os.Create(path.Join(logPath, fileName))
	if err != nil {
		Error(err)
	}
	return f
}

func Trace(args ...interface{}) {
	logConsole.Trace(args)
	if enableFileLog {
		logFile.Trace(args)
	}
}

func Info(args ...interface{}) {
	logConsole.Info(args)
	if enableFileLog {
		logFile.Info(args)
	}
}

func Warn(args ...interface{}) {
	logConsole.Warn(args)
	if enableFileLog {
		logFile.Warn(args)
	}
}

func Error(args ...interface{}) {
	logConsole.Error(args)
	if enableFileLog {
		logFile.Error(args)
	}
}

func Fatal(args ...interface{}) {
	logConsole.Fatal(args)
	if enableFileLog {
		logFile.Fatal(args)
	}
}

func Panic(args ...interface{}) {
	logConsole.Panic(args)
	if enableFileLog {
		logFile.Panic(args)
	}
}
