package logging

import (
	"gin-template/core/util"
	"io"
	"os"
	"path"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"

	"github.com/spf13/viper"
)

// TODO: 使用我们自己的文件输出类，而不是直接调用logrus
func Setup() {
	enableFileLog := viper.GetBool("app.filelog.enable")
	// 是否打印日志到文件中
	// TODO: 应该设置两个log实例，一个写文件的格式，一个写标准输出的格式。标准输出用text，写文件需要json
	if enableFileLog {
		// 设置gin的日志输出
		// TODO: 需要重新设定gin的logger中间件，使用我们自己的log输出，并且每次启动都写入到不同文件中
		serverFile := CreateLogFile("server.log")
		gin.DisableConsoleColor()
		gin.DefaultWriter = io.MultiWriter(os.Stdout, serverFile)

		appFile := CreateLogFile("app.log")
		log.SetFormatter(&log.JSONFormatter{})
		log.SetOutput(io.MultiWriter(os.Stdout, appFile))
	} else {

		log.SetFormatter(&log.TextFormatter{
			ForceColors: true,
		})
	}

}

func CreateLogFile(fileName string) *os.File {
	logPath := viper.GetString("app.filelog.path")
	if err := util.MkdirIfNotExist(logPath); err != nil {
		log.Error(err)
	}
	f, err := os.Create(path.Join(logPath, fileName))
	if err != nil {
		log.Error(err)
	}
	return f
}
